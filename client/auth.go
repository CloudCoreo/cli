package client

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"fmt"
	"io"
	"bytes"
	"time"
)

// Auth struct for API and secret key
type Auth struct {
	APIKey, SecretKey string
}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func computeHmac1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}


func getMD5Hash(body []byte) string {
	hasher := md5.New()
	hasher.Write(body)
	return hex.EncodeToString(hasher.Sum(nil))
}

// SignRequest method to sign all requests
func (a *Auth) SignRequest(req *http.Request) error {
	method := req.Method

	body := []byte{}
	if req.Body != nil {
		body = streamToByte(req.Body)
	}

	mediaType := req.Header.Get("Content-Type")
	date := time.Now().String()
	apiKey := a.APIKey
	secretKey := a.SecretKey


	message := fmt.Sprintf("%s\n%s\n%s\n%s", method, getMD5Hash(body), mediaType, date)

	req.Header.Add("Authorization", "Hmac "+ apiKey + ":" + computeHmac1(message, secretKey))
	req.Header.Add("date", date)

	return nil
}