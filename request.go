package s3client

import (
	"strings"
	"fmt"
	"io"
	"sort"
	"time"
	"net/http"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

type method string

const (
	get = method("GET")
	post = method("POST")
	put = method("PUT")
	delete = method("DELETE")
	empty = ""
	newLine = "\n"
)

type request struct {
	method method
	bucket string
	object string
	headers http.Header
	playload io.Reader 
	suffix string
}

func (req *request) getCanonicalizedResource() (string) {
	bl := len(req.bucket)
	ol := len(req.object)
	sl := len(req.suffix)
	if bl == 0 {
		return "/"
	}
	if ol == 0 {
		if sl == 0 {
			return fmt.Sprintf("/%s/", req.bucket)
		}
		return fmt.Sprintf("/%s?%s", req.bucket, req.suffix)
	}
	if sl == 0 {
		return fmt.Sprintf("/%s/%s", req.bucket, req.object)
	}
	return fmt.Sprintf("/%s/%s?%s", req.bucket,req.object, req.suffix) 
} 

func (req *request) getContentMD5() string { 
	for k, v := range req.headers {
		k = strings.ToLower(k)
		if k == "content-md5" {
			return v[0]
		}
	}
	return empty
}

func (req *request) getContentType() string {
	for k, v := range req.headers {
		k = strings.ToLower(k)
		if k == "content-type" {
			return v[0]
		}
	}
	return empty
}  

func (req *request) getCanonicalizedAmzHeaders() string {
	var items KeyValuePairList
	for k, v := range req.headers {
		k = strings.ToLower(k)
		if strings.HasPrefix(k, "x-amz-") {
			values := strings.Join(v, ",")
			items = append(items, KeyValuePair { k, fmt.Sprintf("%s:%s", k, values) })
		}
	}
	var xamz string
	if len(items) > 0 {
		sort.Sort(items)
		xamz = strings.Join(items.ToArray(), newLine) + newLine
	} 
	return xamz
}

func (req *request) stringToSign() string  { 
	date := time.Now().In(time.UTC).Format(time.RFC1123)
	if (req.headers == nil) {
		req.headers = http.Header{ }
	}
	req.headers["Date"] = []string {date} 
	if req.method == empty {
		req.method = get
	}
	md5 := req.getContentMD5()
	ct := req.getContentType()
	headers := req.getCanonicalizedAmzHeaders()
	resource := req.getCanonicalizedResource()
	verb := string(req.method) 
	signString := fmt.Sprintf("%s\n%s\n%s\n%s\n%s%s", verb, md5, ct, date, headers, resource) 
	return signString 
}

func (req *request) sign(accessKey, secretAccessKey string) string {
	hash := hmac.New(sha1.New, []byte(secretAccessKey))
	signString := req.stringToSign()
	hash.Write([]byte(signString))
	signature := make([]byte, base64.StdEncoding.EncodedLen(hash.Size()))
	base64.StdEncoding.Encode(signature, hash.Sum(nil))
	return fmt.Sprintf("AWS %s:%s", accessKey, string(signature)) 
}