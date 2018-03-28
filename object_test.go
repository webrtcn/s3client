package s3client

import (
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/webrtcn/s3client/models"
)

func TestCreateObject(t *testing.T) {
	filename := "/Users/rain/Downloads/3.jpg"
	client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	bucket := client.NewBucket()
	object := bucket.NewObject("mdh-test2")
	md5, _ := hash(filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fi, _ := f.Stat()
	length := fi.Size()
	defer timeCost(time.Now(), "TestCreateObject")
	err = object.Create("/3.jpg", md5, "application/octet-stream", length, f, models.PublicReadWrite)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("create object successful.")
}

func timeCost(start time.Time, method string) {
	terminal := time.Since(start)
	fmt.Println(method, " 耗时: ", terminal)
}

func md5Content(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	value := base64.StdEncoding.EncodeToString(cipherStr)
	return value
}

func hash(filename string) (md5String string, err error) {
	defer timeCost(time.Now(), "Hash")
	fi, err := os.Open(filename)
	if err != nil {
		return
	}
	defer fi.Close()
	reader := bufio.NewReader(fi)
	md5Ctx := md5.New()
	_, err = io.Copy(md5Ctx, reader)
	if err != nil {
		return
	}
	cipherStr := md5Ctx.Sum(nil)
	value := base64.StdEncoding.EncodeToString(cipherStr)
	return value, nil
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
