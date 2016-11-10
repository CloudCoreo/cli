package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Auth struct for API and secret key
type Auth struct {
	APIKey, SecretKey string
}

func computeHmac1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getMD5Hash(body string) string {
	hasher := md5.New()
	hasher.Write([]byte(body))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SignRequest method to sign all requests
func (a *Auth) SignRequest(req *http.Request) error {
	method := req.Method

	body := ""
	if req.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		body = buf.String()
		req.Body = ioutil.NopCloser(strings.NewReader(body))
		req.ContentLength = int64(len(body))
	}

	mediaType := req.Header.Get("Content-Type")
	date := time.Now().String()
	apiKey := a.APIKey
	secretKey := a.SecretKey

	message := fmt.Sprintf("%s\n%s\n%s\n%s", method, getMD5Hash(body), mediaType, date)

	req.Header.Add("Authorization", "Hmac "+apiKey+":"+computeHmac1(message, secretKey))
	req.Header.Add("date", date)

	return nil
}
