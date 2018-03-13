package s3client

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/webrtcn/s3client/models"
)

//Client s3 Client
type Client struct {
	endPoint        string
	accessKey       string
	secretAccessKey string
}

//NewClient init s3 client
func NewClient(endPoint, accessKey, secretAccessKey string) *Client {
	return &Client{
		endPoint:        endPoint,
		secretAccessKey: secretAccessKey,
		accessKey:       accessKey,
	}
}

//NewBucket new bucket service
func (c *Client) NewBucket() *Bucket {
	return &Bucket{
		client: c,
	}
}

func (c *Client) do(req *request, entity interface{}, bodyFunc func(resp *http.Response)) error {
	httpClient := &http.Client{}
	reqURL, err := url.Parse(c.endPoint + req.getQueryString())
	if err != nil {
		return err
	}
	sign := req.sign(c.accessKey, c.secretAccessKey)
	reqHeaders := http.Header{
		"Authorization": {sign},
		"Host":          {reqURL.Host},
	}
	for k, v := range req.headers {
		reqHeaders.Add(k, v[0])
	}
	fmt.Println(reqHeaders)
	httpRequest := &http.Request{
		URL:    reqURL,
		Method: string(req.method),
		Header: reqHeaders,
	}
	if v, ok := req.headers["Content-Length"]; ok {
		httpRequest.ContentLength, _ = strconv.ParseInt(v[0], 10, 64)
		if httpRequest.ContentLength == 0 {
			req.playload = nil
		} else {
			httpRequest.Body = req.playload
		}
	}
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return err
	}
	if bodyFunc != nil {
		bodyFunc(resp)
	} else {
		defer resp.Body.Close()
		if entity == nil {
			return nil
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		switch resp.StatusCode {
		case 200:
			if len(body) > 0 {
				err := xml.Unmarshal(body, entity)
				if err != nil {
					return err
				}
			}
		case 400, 403, 404, 405, 408, 409, 411, 412, 416, 422, 500:
			if len(body) > 0 {
				er := &models.ErrorResult{}
				err = xml.Unmarshal(body, er)
				if err != nil {
					return err
				}
				return errors.New(er.Code)
			}
			return fmt.Errorf("Request error, status code: %s", strconv.Itoa(resp.StatusCode))
		}
	}
	return nil
}
