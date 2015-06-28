package cavirtex

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func (ex *CaVirtex) AuthenticatedPost(action string, params map[string]string) (*http.Response, error) {
	mac := hmac.New(sha256.New, []byte(ex.APIKey))
	path := fmt.Sprintf("/api2/user/%s.json", action)
	nonce := fmt.Sprintf("%d", time.Now().UnixNano())

	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	macableMessage := []string{
		nonce,
		ex.APIToken,
		path,
	}

	for _, k := range keys {
		macableMessage = append(macableMessage, params[k])
	}
	mac.Write([]byte(strings.Join(macableMessage, "")))
	signature := mac.Sum(nil)

	params["signature"] = fmt.Sprintf("%x", signature)
	params["nonce"] = nonce
	params["token"] = ex.APIToken

	vals := url.Values{}
	for k, v := range params {
		vals.Set(k, v)
	}

	return http.PostForm(fmt.Sprintf("https://www.cavirtex.com%s", path), vals)
}
