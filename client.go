package s3client

import( 
	"fmt" 
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"io/ioutil"
	"encoding/xml"
	"github.com/webrtcn/s3client/models"
)

type client struct {
	endPoint string
	accessKey string
	secretAccessKey string
}

//NewClient init client 
func NewClient(endPoint, accessKey, secretAccessKey string) *client {
	return &client {
		endPoint: endPoint,
		secretAccessKey: secretAccessKey,
		accessKey: accessKey,
	}
}

//NewBucket new bucket service
func (c *client) NewBucket() *bucket {
	return &bucket {
		client: c,
	}
} 

func (c *client) do(req *request, entity interface{}) error {
	httpClient := &http.Client {}
	reqURL, err := url.Parse(c.endPoint + req.getQueryString())
	if err != nil {
		return err
	} 
	sign := req.sign(c.accessKey, c.secretAccessKey)
	reqHeaders := http.Header {  
		"Authorization": { sign },
		"Host": { reqURL.Host }, 
	}  
	for k, v := range req.headers {
		reqHeaders.Add(k, v[0])
	} 
	httpRequest := &http.Request {
		URL: reqURL,
		Method: string(req.method),
		Header: reqHeaders,
	}
	if v, ok := req.headers["Content-Length"]; ok {
		httpRequest.ContentLength, _ = strconv.ParseInt(v[0], 10, 64)  
		if httpRequest.ContentLength == 0 {
			req.playload = nil
		} else {
			httpRequest.Body = ioutil.NopCloser(req.playload)
		}
	}
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer resp.Body.Close() 
	fmt.Println("Status Code: " + strconv.Itoa(resp.StatusCode))
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("######################")
	fmt.Println(string(body))
	fmt.Println("######################")  
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
		return errors.New(fmt.Sprintf("Request error, status code: %s", strconv.Itoa(resp.StatusCode)))
	}
	return nil
}