package util

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestBinaryEncode(t *testing.T) {
	var a [24]byte
	head := []byte("eyJhbGciOiJIUzI1NiIsImtpZCI6InNpbTIifQ")
	payload := []byte("eyJhdWQiOiJodHRwczovL2FwaS54azUuY29tL2dhdGV3YXkvdjIuMCIsImV4cCI6MTY4MTIyNjE3NCwiaWF0IjoxNjgxMjA0NTc0ODQ4LCJpc3MiOiJodHRwczovL2FwaS54azUuY29tIiwic3ViIjoiYSJ9")
	signature := []byte("WxQ4xc2gnuY1cmmjoYuXkxM3Y8QYEfvxIgaAsmVFpwI")
	binary.BigEndian.PutUint64(a[:8], binary.BigEndian.Uint64(head))
	binary.BigEndian.PutUint64(a[8:16], binary.BigEndian.Uint64(payload))
	binary.BigEndian.PutUint64(a[16:], binary.BigEndian.Uint64(signature))
	t.Logf("%q\n", string(a[:]))
	var buf [48]byte
	hex.Encode(buf[:], a[:])
	t.Logf("%q\n", string(buf[:]))
	t.Logf("%x\n", buf)
	t.Logf("%x\n", a)
	t.Logf("%x\n", md5.Sum(buf[:]))

	t.Log(len(head), reflect.TypeOf(head))
}

func TestBinaryDecode(t *testing.T) {
	var a [24]byte
	b := a[:]
	hex.Decode(b, []byte("65794a686247636965794a68645751695778513478633267"))
	t.Logf("%q\n", string(b))

	var s []byte
	buf := bytes.NewReader(b)
	if err := binary.Read(buf, binary.BigEndian, &s); err != nil {
		t.Error(err)
	}
	t.Logf("%s", string(s))
}

func TestHexEncode(t *testing.T) {
	t.Log(base64.URLEncoding.EncodeToString([]byte("sgxq_msdk_android")))
	t.Log(base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("sgxq_msdk_android%v", time.Now().Unix()))))
	t.Logf("%x", md5.Sum([]byte(fmt.Sprintf("sgxq_msdk_android%v", time.Now().Unix()))))
	t.Log(hex.EncodeToString([]byte("SGXQ")))
	t.Log(hex.EncodeToString([]byte(fmt.Sprintf("sgxq%v", time.Now().Unix()))))
	t.Logf("%x", md5.Sum([]byte(fmt.Sprintf("sgxq%v", time.Now().Unix()))))

	id, _ := base64.URLEncoding.DecodeString("Z2F0ZXdheQ==")
	secret, _ := base64.URLEncoding.DecodeString("MDAxZ2F0ZXdheQ==")
	t.Log(string(id))
	t.Log(string(secret))

}

func TestDecrypt(t *testing.T) {
	type args struct {
		ctx      context.Context
		tokenStr string
		secret   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{ctx: context.Background(), tokenStr: "25f3c76eaaa9a7815cbc7a1b843188de2d9ec8a7d60b23848e89829c8a1c91f95bda68ec0430dbf5bf34c522064f7f63c4cea25e94212f3cb137858e3f8ea4576393e94df248e88867b486a47e46271692ee398b45b3b1a038fb9f32a341b4ec8362f63f031dc0baf349824815074ad22294a2898cd087448b755ee3f913dbc861b2e1107ba662e5a364075f523310968bb04a1c74d74d3a1f95de0db848764cf75c87e5055393535a53ff66850c08236e96a4d017a99f6ac938de383bb2471bf6a4629fcd46219472295da6c34157477829aecfe8c5a75fc24a646cbcc7904709d11bcbb97b0b15be82470079224f8d9e3804d0c2be3f1743ceca8981041ede1a34512d0ad832cafeb82517b8e71fe64e3b3364873862869a91fc6d80f3f193fe316900dcd4d18b50d4815403bc5c59cea4", secret: "5150258a434f1a065056c54018286386"},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2VudCI6IlhLNV9TRVJWRVIiLCJhcHBsaWNhdGlvbl9ubyI6Mjk5OTksImVudiI6InNkZXYwIiwiZXhwaXJlX3RpbWUiOjE2NzYwMzY1\nNDE0NTQsIm1hcmtldF9nYW1lX2lkIjoiIiwic3ZyX2lkIjoiSU5GUkFTRVJWRVIiLCJ0aW1lc3RhbXAiOjE2NzYwMTQ5NDE0NTR9.Is-jBfOe5x6DvClqMlZ04dBJJ_n4k6OmK6yH3oliTUc",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				ctx:      context.Background(),
				tokenStr: "d6609ed949e81dd1622437749a76716c07859187914f",
				secret:   "5150258a434f1a065056c54018286386",
			},
			want:    "123456",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				ctx:      context.Background(),
				tokenStr: "MAkbeGb3VK+avw0r3010oQ==",
				secret:   "f1b5fe56585bc5a04e71df7f791deb78",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				ctx:      context.Background(),
				tokenStr: "c50e64e52144c39412d5a2bbc04a8093194a22b409448ab5f19d4dc605d8adc9666ac63d5b60d6868267f7c6993c93fa0df5566ef5680dcb2ad7dae79a043da334a3e21a8d4d0b246dbaece80c35c50e73c645da7a6bd3",
				secret:   "252bc099eaeb8eeb5435720447d506c4",
			},
			want:    "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CFBDecrypt(tt.args.tokenStr, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecryptToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	type args struct {
		ctx      context.Context
		tokenStr string
		secret   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				ctx:      context.Background(),
				tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2VudCI6IlhLNV9TRVJWRVIiLCJhcHBsaWNhdGlvbl9ubyI6Mjk5OTksImVudiI6InNkZXYwIiwiZXhwaXJlX3RpbWUiOjE2NzYwMzY1\nNDE0NTQsIm1hcmtldF9nYW1lX2lkIjoiIiwic3ZyX2lkIjoiSU5GUkFTRVJWRVIiLCJ0aW1lc3RhbXAiOjE2NzYwMTQ5NDE0NTR9.Is-jBfOe5x6DvClqMlZ04dBJJ_n4k6OmK6yH3oliTUc",
				secret:   "5150258a434f1a065056c54018286386",
			},
			want:    "0026422b7483ced598eaac65b6e01a9c",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				ctx:      context.Background(),
				tokenStr: "123456",
				secret:   "5150258a434f1a065056c54018286386",
			},
			want:    "0026422b7483ced598eaac65b6e01a9c",
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				ctx:      context.Background(),
				tokenStr: "93534750d24f4d7e",
				secret:   "f1b5fe56585bc5a04e71df7f791deb78",
			},
			want:    "0026422b7483ced598eaac65b6e01a9c",
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				ctx:      context.Background(),
				tokenStr: "sdev0:auth-v4:refresh_token:user:1000000000771703:appid:10001:devicetype:PC",
				secret:   "252bc099eaeb8eeb5435720447d506c4",
			},
			want:    "0026422b7483ced598eaac65b6e01a9c",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CFBEncrypt(tt.args.tokenStr, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncryptToken() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestBase64(t *testing.T) {
	t.Log(url.QueryEscape("https://github.com/cstcen/gore"))
	t.Log(hex.EncodeToString([]byte("5150258a434f1a065056c54018286386")))
}

func TestECBDecrypt(t *testing.T) {
	type args struct {
		data string
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				data: "1StBMP2kqImfxsIyAllDbw==",
				key:  "5150258a434f1a065056c54018286386",
			},
			want: "abcdefg123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ECBDecrypt(tt.args.data, tt.args.key); got != tt.want {
				t.Errorf("ECBDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECBEncrypt(t *testing.T) {
	type args struct {
		data string
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				data: "abcdefg123",
				key:  "5150258a434f1a065056c54018286386",
			},
			want: "1StBMP2kqImfxsIyAllDbw==",
		}, {
			name: "test",
			args: args{
				data: "25413DBCA5217507D964556F5A03814D",
				key:  "f1b5fe56585bc5a04e71df7f791deb78",
			},
			want: "1StBMP2kqImfxsIyAllDbw==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ECBEncrypt(tt.args.data, tt.args.key); got != tt.want {
				t.Errorf("ECBEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
