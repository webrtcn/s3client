package s3client

import (
	"testing"
)

func TestInitiateUpload(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// metas := make(map[string]string)
	// metas["cs"] = "100kb"
	// option := &CreateUploadsOption{
	// 	Metadatas: metas,
	// }
	// fmt.Println(option.Perm)
	// uploader := object.NewUploads("qp.txt", option)
	// e, err := uploader.Initiate(nil)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(e)
	// }
}

func TestUploadPart(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// metas := make(map[string]string)
	// metas["cs"] = "100kb"
	// option := &CreateUploadsOption{
	// 	Metadatas: metas,
	// }
	// fmt.Println(option.Perm)
	// uploader := object.NewUploads("qp.txt", option)
	// ret, err := uploader.Initiate(nil)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// data := []byte("<root>hello root</root>")
	// md5 := md5Content(data)
	// length := int64(len(data))
	// body := ioutil.NopCloser(bytes.NewReader(data))
	// part, err := uploader.UploadPart(0, ret.UploadID, md5, "text/xml", length, body)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(part)
	// }
}

func TestListPart(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// uploader := object.NewUploads("qp.txt", nil)
	// uploadID := "2~v6_Y1j-0Kc46GKnJUJANYz6E-N5VIRS"
	// for i := 0; i < 4; i++ {
	// 	data := []byte("<root>" + strconv.Itoa(i) + "</root>")
	// 	md5 := md5Content(data)
	// 	length := int64(len(data))
	// 	body := ioutil.NopCloser(bytes.NewReader(data))
	// 	part, err := uploader.UploadPart(i, uploadID, md5, "text/xml", length, body)
	// 	if err != nil {
	// 		t.Error(err)
	// 	} else {
	// 		fmt.Println(part)
	// 	}
	// }
	// value, err := uploader.ListPart("2~v6_Y1j-0Kc46GKnJUJANYz6E-N5VIRS")
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(value)
	// }
}

func TestComplete(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// uploader := object.NewUploads("qp0.txt", nil)
	// value, err := uploader.Initiate(nil)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// uploadID := value.UploadID
	// var parts CompleteParts
	// for i := 1; i < 2; i++ {
	// 	data := []byte("<root>" + strconv.Itoa(i) + "</root>")
	// 	md5 := md5Content(data)
	// 	length := int64(len(data))
	// 	body := ioutil.NopCloser(bytes.NewReader(data))
	// 	part, err := uploader.UploadPart(i, uploadID, md5, "text/xml", length, body)
	// 	if err != nil {
	// 		t.Error(err)
	// 	} else {
	// 		parts = append(parts, *part)
	// 	}
	// }
	// value1, err := uploader.ListPart(uploadID)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(value1)
	// }
	// // err = uploader.RemoveUploads(uploadID)
	// // if err != nil {
	// // 	t.Error(err)
	// // 	return
	// // }
	// value2, err := uploader.Complete(uploadID, parts)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Println(value2)
	// }
}

func TestRemoveUploads(t *testing.T) {
	// client := NewClient("http://example.com", "accessKey", "secretAccesskey")
	// bucket := client.NewBucket()
	// object := bucket.NewObject("mdh-test2")
	// uploader := object.NewUploads("qp0.txt", nil)
	// value, err := uploader.Initiate(nil)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// err = uploader.RemoveUploads(value.UploadID)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
}
