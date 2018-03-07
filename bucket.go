package s3client

import( 
	"net/http"
	"github.com/webrtcn/s3client/models"
)

type bucket struct { 
	client *client
}

func (b *bucket) List() (*models.BucketsResult, error) {
	req := &request {
		method: get, 
	}
	buckets := &models.BucketsResult{}
	err := b.client.do(req, buckets)
	return buckets, err
}

func(b *bucket) Create(name string, perm models.ACL) error {
	req := &request {
		method: put,
		bucket: name,
		headers: http.Header {
			"x-amz-acl": {string(perm)},
		},
	}  
	err := b.client.do(req, nil)
	return err
}