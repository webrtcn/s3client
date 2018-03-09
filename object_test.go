package s3client

import (
	"crypto/md5"
	"encoding/base64"
	"testing"
)

func TestCreateObject(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// data := []byte("<root>hello root</root>")
	// md5 := md5Content(data)
	// length := int64(len(data))
	// body := ioutil.NopCloser(bytes.NewReader(data))
	// err := object.Create("lyl.txt", md5, "text/xml", length, body, models.PublicReadWrite)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
}

func md5Content(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	value := base64.StdEncoding.EncodeToString(cipherStr)
	return value
}

func TestGetObject(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// resp, err := object.Get("lyl.txt", nil)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(resp)
	// }
}

func TestCopyObject(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// value, err := object.Copy("mdh-test2", "lyl.txt", "lyl1.txt", nil)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(value)
	// }
}

func TestRemoveObject(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// err := object.Remove("lyl1.txt")
	// if err != nil {
	// 	t.Error(err)
	// }
}

func TestGetHeader(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// resp, err := object.GetHeader("lyl.txt", &GetObjectOption{
	// 	IfMatchTag:      "e6056bba267f36bca2e6cd5c949f756c",
	// 	IfModifiedSince: "Fri, 09 Mar 2017 07:14:53 GMT",
	// })
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(resp)
	// }
}

func TestGetObjectACL(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// value, err := object.GetACL("lyl.txt")
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(value)
	// }
}

func TestSetObjectACL(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// err := object.SetACL("lyl.txt", models.PublicRead)
	// if err != nil {
	// 	t.Error(err)
	// }
}
