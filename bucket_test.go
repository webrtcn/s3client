package s3client

import (
	"fmt"
	"testing"
	"github.com/webrtcn/s3client/models"
)

func TestList(t *testing.T) { 
	client := NewClient("http://example.com", "abc", "1234567")
	bucket := client.NewBucket()
	values, err := bucket.List()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(values.Owner.OwnerID)
	}
}

func TestCreate(t *testing.T) {
	client := NewClient("http://example.com", "abc", "1234567")
	bucket := client.NewBucket()
	err := bucket.Create("-hello-world-", models.PublicReadWrite)
	if err != nil {
		t.Error(err.Error())
	} 
}