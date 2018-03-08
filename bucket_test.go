package s3client

import (
	"fmt"
	"testing"
	//"github.com/webrtcn/s3client/models"
)

func TestList(t *testing.T) {
	// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// values, err := bucket.List()
	// if err != nil {
	// 	t.Error(err.Error())
	// } else {
	// 	fmt.Println(values.Owner.OwnerID)
	// }
}

func TestCreate(t *testing.T) {
	// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// err := bucket.Create("-hello-world-", models.PublicReadWrite)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
}

func TestDelete(t *testing.T) {
	// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// err := bucket.Remove("hello-world")
	// if err != nil {
	// 	t.Error(err.Error())
	// }
}

func TestGet(t *testing.T) {
	fmt.Println("test get")
	// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// value, err := bucket.Get("mdh-test2", "", "", "0", 1000)
	// if err != nil {
	// 	t.Error(err.Error())
	// } else {
	// 	fmt.Println(value)
	// }
}

func TestGetLocation(t *testing.T) {
	// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// value, err := bucket.GetLocation("mdh-test2")
	// if err != nil {
	// 	t.Error(err.Error())
	// } else {
	// 	fmt.Println(value)
	// }
}

func TestGetACL(t *testing.T) {
	// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// err := bucket.GetACL("mdh-test2")
	// if err != nil {
	// 	t.Error(err.Error())
	// }
}
func TestSetACL(t *testing.T) {
	/// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// err := bucket.SetACL("mdh-test2", models.PublicReadWrite)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
}

func TestListUploads(t *testing.T) {
	// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// v, err := bucket.ListUploads("mdh-test2", "", "", "", "", 1000, 10)
	// if err != nil {
	// 	t.Error(err.Error())
	// } else {
	// 	fmt.Println(v)
	// }
}

func TestSetVersioning(t *testing.T) {
	// client := NewClient("http://example.com", "abc", "1234567")
	// bucket := client.NewBucket()
	// err := bucket.SetVersioning("mdh-test2", true)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
}
