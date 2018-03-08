# go-ceph-s3-client 

## Installation

Install the package with:

```
go get github.com/webrtcn/s3client
```

Import it with:

```
import "github.com/webrtcn/s3client"
```

#### Support feature list

Feature | Status | Remark
---|---|---
List Buckets |  √
Delete Bucket|  √
Create Bucket|  √
Bucket ACLs (Get, Put) | √
Bucket Location | √
Bucket Object Versions | √
Get Bucket Info (HEAD) |√
Put Object|x|developing
Delete Object|x|developing
Get Object|x|developing
Object ACLs (Get, Put)|x|developing
Get Object Info (HEAD)|x|developing
POST Object|x|developing
Copy Object|x|developing
Multipart Uploads|x|developing
List Multipart Uploads|√

## Example


```

package main

import (
	"fmt"
	"testing"
	. "github.com/webrtcn/s3client"
	. "github.com/webrtcn/s3client/models"
)

func main {
    //list all buckets
	client := NewClient("http://example.com", "accessKey", "secretAccessKey")
	bucket := client.NewBucket()
	values, err := bucket.List()
	if err != nil {
	 	t.Error(err.Error())
	 } else {
	 	fmt.Println(values.Owner.OwnerID)
	 }
}

```
