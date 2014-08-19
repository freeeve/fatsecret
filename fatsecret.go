package fatsecret

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type FatSecretConn struct {
	apikey string
	secret string
}

var fsurl = "http://platform.fatsecret.com/rest/server.api"

func Connect(apikey, secret string) (FatSecretConn, error) {
	return FatSecretConn{
		apikey,
		secret,
	}, nil
}

func (fs FatSecretConn) get(method string, params map[string]string) (io.ReadCloser, error) {
	reqTime := fmt.Sprintf("%d", time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := map[string]string{
		"method":                 method,
		"oauth_consumer_key":     fs.apikey,
		"oauth_nonce":            fmt.Sprintf("%d", r.Int63()),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        reqTime,
		"oauth_version":          "1.0",
		"format":                 "json",
	}
	for k, v := range params {
		m[k] = v
	}
	mk := make([]string, len(m))
	i := 0
	for k, _ := range m {
		mk[i] = k
		i++
	}
	// sort keys
	sort.Strings(mk)

	// build sorted k/v string for sig
	sigQueryStr := ""
	for _, k := range mk {
		sigQueryStr += fmt.Sprintf("&%s=%s", k, escape(m[k]))
	}
	// drop initial &
	sigQueryStr = sigQueryStr[1:]
	sigBaseStr := fmt.Sprintf("GET&%s&%s", url.QueryEscape(fsurl), escape(sigQueryStr))
	//fmt.Println("sigstr:", sigBaseStr)

	mac := hmac.New(sha1.New, []byte(fs.secret+"&"))
	mac.Write([]byte(sigBaseStr))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	// add sig to map
	m["oauth_signature"] = sig
	mk = append(mk, "oauth_signature")

	// re-sort keys after adding sig
	sort.Strings(mk)
	requrl := fmt.Sprintf("%s?", fsurl)
	reqQuery := ""
	for _, k := range mk {
		reqQuery += fmt.Sprintf("&%s=%s", k, escape(m[k]))
	}
	// drop initial &
	reqQuery = reqQuery[1:]

	requrl += reqQuery
	//fmt.Println("url :", requrl)
	resp, err := http.Get(requrl)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func escape(s string) string {
	return strings.Replace(strings.Replace(url.QueryEscape(s), "+", "%20", -1), "%7E", "~", -1)
}
