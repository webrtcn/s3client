package models

import(
	"encoding/xml"
	"time"
)

type BucketsResult struct {
	XMLName xml.Name `xml:"ListAllMyBucketsResult"`
	Owner Owner
	Buckets BucketDetail
}

type Owner struct {
	OwnerID string 		`xml:"ID"`
	DisplayName string
}

type BucketDetail struct {
	Bucket []BucketItem
}

type BucketItem struct {
	Name string
	CreationDate time.Time
}