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

## Example


```

package main

import (
	"fmt"
	"testing"
	. "github.com/webrtcn/s3client"
	"github.com/webrtcn/s3client/models"
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
