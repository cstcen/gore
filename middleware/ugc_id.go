package middleware

import (
	"fmt"
	"github.com/cstcen/gore/common"
	"github.com/cstcen/gore/gonfig"
	goreHttp "github.com/cstcen/gore/http"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

func ConvertGuidToUgcId(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	ugcId, err := NewUgcIdPathGetting().GetUgcId(ctx)
	if err != nil {
		return
	}
	urlString := fmt.Sprintf("%s/ugc/v2.2/internal/ugc/validate?guid=%s", gonfig.Instance().GetString("xk5.host.external"), ugcId)
	var result common.DataResult[map[string]any]
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", token)
	resp, err := goreHttp.Instance().Do(req)
	if err != nil {
		return
	}
	if err := goreHttp.RespHandler(resp, &result); err != nil {
		return
	}
	if result.Code != 0 || result.Data == nil {
		return
	}
	val, ok := result.Data["ugcId"].(string)
	if !ok {
		return
	}
	ugcId = val

	NewUgcIdContextSetting().SetUgcId(ctx, ugcId)
}

type UgcIdGetter interface {
	GetUgcId(ctx *gin.Context) (string, error)
}
type UgcIdSetter interface {
	SetUgcId(ctx *gin.Context, ugcId string) error
}

type UgcIdPathGetting struct {
}

func NewUgcIdPathGetting() *UgcIdPathGetting {
	return &UgcIdPathGetting{}
}

func (u *UgcIdPathGetting) GetUgcId(ctx *gin.Context) (string, error) {
	req := ctx.Request
	pattern := `/ugc/([^/]+)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	curPath := req.URL.Path
	matches := regex.FindStringSubmatch(curPath)
	if len(matches) < 2 {
		return "", fmt.Errorf("ugcId not found")
	}
	return matches[1], nil
}

type UgcIdContextSetting struct {
}

func NewUgcIdContextSetting() *UgcIdContextSetting {
	return &UgcIdContextSetting{}
}

func (u *UgcIdContextSetting) SetUgcId(ctx *gin.Context, ugcId string) error {
	for i, param := range ctx.Params {
		if !strings.EqualFold(param.Key, "ugcId") {
			continue
		}
		param.Value = ugcId
		ctx.Params[i] = param
	}
	return nil
}

type UgcIdPathSetting struct {
}

func (u *UgcIdPathSetting) SetUgcId(ctx *gin.Context, ugcId string) error {
	req := ctx.Request
	pattern := `/ugc/([^/]+)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	curPath := req.URL.Path
	matches := regex.FindStringSubmatch(curPath)
	if len(matches) < 2 {
		return nil
	}
	updatedUgcId := regex.ReplaceAllString(curPath, fmt.Sprintf("/ugc/%s", ugcId))
	req.URL.Path = updatedUgcId

	return nil
}
