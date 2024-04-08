// Package report 腾讯举报系统
// 请查看文档：举报反馈系统http接入说明.doc
package report

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/cstcen/gore/errorcode/v2"
	goreHttp "github.com/cstcen/gore/http"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	HostExp    = "service-light-exp.gamesafe.qq.com"
	HostFormal = "service-light.gamesafe.qq.com"
	Path       = "/report/report_log_data"
)

var (
	AppIdMap = map[string]int{"224": 2665}
)

type Informant struct {
	InformantAccountType int    `json:"informant_account_type"`
	InformantAccountId   string `json:"informant_account_id"`
	InformantAppId       int    `json:"informant_app_id"`
	InformantPlatId      int    `json:"informant_plat_id"`
	InformantWorldId     int    `json:"informant_world_id"`
	InformantName        string `json:"informant_name"`
	InformantRoleid      string `json:"informant_roleid"`
	InformantArea        int    `json:"informant_area,omitempty"`
}

type Reported struct {
	ReportedAccountType int    `json:"reported_account_type"`
	ReportedAccountId   string `json:"reported_account_id"`
	ReportedAppId       int    `json:"reported_app_id"`
	ReportedPlatId      int    `json:"reported_plat_id"`
	ReportedWorldId     int    `json:"reported_world_id"`
	ReportedName        string `json:"reported_name"`
	ReportedRoleid      string `json:"reported_roleid"`
	ReportedArea        int    `json:"reported_area,omitempty"`
}

type ContentId struct {
	IdType string   `json:"id_type"`
	IdList []string `json:"id_list"`
}

type BusinessData struct {
	ReportCategory     int         `json:"report_category"`
	ReportReason       []int       `json:"report_reason"`
	ReportScene        int         `json:"report_scene"`
	ReportedProfileUrl string      `json:"reported_profile_url"`
	ReportBattleId     string      `json:"report_battle_id,omitempty"`
	ReportBattleTime   int         `json:"report_battle_time,omitempty"`
	ReportDesc         string      `json:"report_desc,omitempty"`
	ReportContent      string      `json:"report_content,omitempty"`
	PicUrlArray        []string    `json:"pic_url_array,omitempty"`
	VideoUrlArray      []string    `json:"video_url_array,omitempty"`
	ReportGroupId      string      `json:"report_group_id,omitempty"`
	ReportGroupName    string      `json:"report_group_name,omitempty"`
	ContentId          []ContentId `json:"content_id,omitempty"`
	Callback           string      `json:"callback,omitempty"`
	ReportEntrance     int         `json:"report_entrance,omitempty"`
}

type RequestBody struct {
	Informant
	Reported
	ReportTime   int          `json:"report_time,omitempty"`
	BusinessData BusinessData `json:"business_data"`
}

type ResponseBody struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func host(env string) string {
	if env == "xk5" {
		return HostFormal
	}
	return HostExp
}

func token() string {
	return "HwkPvdUqZvZcDDHONfhpKWTClCuDndby"
}

func buildQuery(body []byte) string {
	params := buildParams()
	buf := bytes.NewBufferString(params.Encode())
	signature := sign(params, body)
	buf.WriteString("&sign=" + signature)
	return buf.String()
}

func sign(params *url.Values, body []byte) string {
	tk := token()
	bodyMd5 := strings.ToUpper(fmt.Sprintf("%x", md5.Sum(body)))

	buf := bytes.NewBufferString(params.Encode())
	buf.WriteString("&body_md5=" + bodyMd5)
	buf.WriteString("&token=" + tk)
	signature := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
	return strings.ToUpper(signature)
}

func AppId(gameNo string) int {
	a, _ := AppIdMap[gameNo]
	return a
}

func Request(ctx context.Context, env string, requestBody *RequestBody) (*ResponseBody, error) {

	reqBodyRaw, err := json.Marshal(requestBody)
	if err != nil {
		return nil, errorcode.BadRequestInvalidParameter.WithMsg("marshal request body error: " + err.Error())
	}
	u := url.URL{
		Scheme:   "https",
		Host:     host(env),
		Path:     Path,
		RawQuery: buildQuery(reqBodyRaw),
	}
	body := bytes.NewReader(reqBodyRaw)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), body)
	if err != nil {
		return nil, errorcode.InternalServerError.WithMsg("new request error: " + err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(body.Len()))
	resp, err := goreHttp.Instance().Do(req)
	if err != nil {
		return nil, errorcode.InternalServerError.WithMsg("do request error: " + err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.New(int32(resp.StatusCode), "invalid status code")
	}
	respBodyRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorcode.InternalServerError.WithMsg("read response error: " + err.Error())
	}
	defer resp.Body.Close()

	var respBody ResponseBody
	if err := json.Unmarshal(respBodyRaw, &respBody); err != nil {
		return nil, errorcode.InternalServerError.WithMsg("unmarshal response error: " + err.Error())
	}

	if respBody.ErrCode != 0 {
		return nil, errorcode.New(int32(respBody.ErrCode), respBody.ErrMsg)
	}

	return &respBody, nil
}

func buildParams() *url.Values {
	now := time.Now()
	// rand.Seed(now.UnixNano())
	user := "2665_report"
	timestamp := strconv.FormatInt(now.Unix(), 10)
	// serial := rand.Intn(10000)
	params := url.Values{}
	params.Set("user", user)
	params.Set("timestamp", timestamp)
	params.Set("serial", timestamp)
	return &params
}
