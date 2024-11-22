package rtc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/util"
)

// 获取owt公用auth头
func (r *RTC) getOWTAuthorizationHeader() string {
	strBuilder := strings.Builder{}

	nonce := rand.Intn(99999)
	authTimestamp := time.Now().UnixNano() / 1e6

	strBuilder.WriteString("MAuth ")
	strBuilder.WriteString("realm=rtc.tgo.io,")
	strBuilder.WriteString("mauth_signature_method=HMAC_SHA256,")
	strBuilder.WriteString(fmt.Sprintf("mauth_serviceid=%s,", r.ctx.GetConfig().OWT.ServiceID))
	strBuilder.WriteString(fmt.Sprintf("mauth_cnonce=%d,", nonce))
	strBuilder.WriteString(fmt.Sprintf("mauth_timestamp=%d,", authTimestamp))
	sign := util.HmacSha256(fmt.Sprintf("%d,%d", authTimestamp, nonce), r.ctx.GetConfig().OWT.ServiceKey)
	strBuilder.WriteString(fmt.Sprintf("mauth_signature=%s", sign))

	return strBuilder.String()
}

func (r *RTC) post(path string, bodyMap map[string]interface{}) ([]byte, error) {
	reqURL := fmt.Sprintf("%s%s", r.ctx.GetConfig().OWT.URL, path)
	fmt.Println("reqURL-->POST ", reqURL, bodyMap)
	req, err := http.NewRequest("POST", reqURL, bytes.NewBufferString(util.ToJson(bodyMap)))
	if err != nil {
		return nil, err
	}
	auth := r.getOWTAuthorizationHeader()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)
	resp, err := r.owtClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		resultBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return resultBytes, nil
	}
	resultBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body-->", string(resultBytes))
	return nil, fmt.Errorf("请求状态错误！->%d", resp.StatusCode)
}

func (r *RTC) getRoomPresenterToken(roomID, uid string) (string, error) {
	return r.getRoomToken(roomID, uid, RolePresenter.String())
}

// 获取指定用户在房间的token
func (r *RTC) getRoomToken(roomID string, uid string, role string) (string, error) {
	resp, err := r.post(fmt.Sprintf("/rooms/%s/tokens", roomID), map[string]interface{}{
		"room": roomID,
		"user": uid,
		"role": role,
		"preference": map[string]string{
			"isp":    "isp",
			"region": "region",
		},
	})
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
