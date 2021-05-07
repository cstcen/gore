package gore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	GrantType      = "client_credentials"
	AuthID         = "infra_billing_server"
	AuthSecret     = "alkjsdf8jsf9n3onf78s9dhfjlk398f9hlksdfuihaoisdhf"
	AuthHostFormat = "https://m-apis-%s.xk5.com/auth/v2/infra_server/init"
)

var (
	req *InfraTokenRequest
)

type InfraTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	MacAddress   string `json:"mac_address"`
}

type InfraTokenResponse struct {
	ReturnCode    uint   `json:"return_code"`
	ReturnMessage string `json:"return_message,omitempty"`
	ExpiresIn     uint   `json:"expires_in,omitempty"`
	AccessToken   string `json:"access_token,omitempty"`
}

func init() {
	req = &InfraTokenRequest{
		GrantType:    GrantType,
		ClientId:     AuthID,
		ClientSecret: AuthSecret,
		MacAddress:   GetMACAddr(),
	}
}

func GetInfraToken(env string) (*InfraTokenResponse, error) {

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	authURL := fmt.Sprintf(AuthHostFormat, env)

	resp, err := http.Post(authURL, ContentTypeApplicationJSONCharset, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	result := new(InfraTokenResponse)
	if err = json.Unmarshal(body, result); err != nil {
		return nil, err
	}

	return result, nil
}
