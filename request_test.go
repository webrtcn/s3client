package s3client

import (   
	"net/http"
	"testing"
)

func TestGetCanonicalizedResource(t *testing.T) {
	req := &request { }
	path := req.getCanonicalizedResource()
	if (path != "/") {
		t.Error("Path error")
	}
	req = &request {
		bucket: "bucket",
	}
	path = req.getCanonicalizedResource()
	if (path != "/bucket/") {
		t.Error("/bucket/ path error")
	}
	req = &request {
		bucket: "bucket",
		object: "test.txt",
	}
	path = req.getCanonicalizedResource()
	if (path != "/bucket/test.txt") {
		t.Error("/bucket/test.txt path error")
	}
	req = &request {
		bucket: "bucket",
		object: "",
		suffixs: []suffix {
			suffix {
				key: "acl",
				flag: true,
			},
		},
	}
	path = req.getCanonicalizedResource()
	if (path != "/bucket?acl") {
		t.Error("/bucket?acl path error")
	}
	req = &request {
		bucket: "bucket",
		object: "test.txt",
		suffixs: []suffix {
			suffix {
				key: "uploads",
				flag: true,
			},
		},
	}
	path = req.getCanonicalizedResource()
	if (path != "/bucket/test.txt?uploads") {
		t.Error("/bucket/test.txt?uploads path error")
	}
	req = &request {
		bucket: "bucket",
		object: "this/is/a/test.txt",
		suffixs: []suffix {
			suffix {
				key: "partNumber",
				value: "1",
				flag: false,
			},
			suffix {
				key: "uploadId",
				value: "2~0_EI37ycVV1zh-KAjDZgn15VtyoM7KK",
				flag: false,
			},
		}, 
	}
	path = req.getCanonicalizedResource() 
	if (path != "/bucket/this/is/a/test.txt?partNumber=1&uploadId=2~0_EI37ycVV1zh-KAjDZgn15VtyoM7KK") {
		t.Error("/bucket/this/is/a/test.txt?partNumber=1&uploadId=2~0_EI37ycVV1zh-KAjDZgn15VtyoM7KK path error")
	}

	suffixs := []suffix {
		suffix {
			key: "max-keys",
			value: "1000",
			flag: false,
		},
		suffix {
			key: "marker",
			value: "1",
			flag: false,
		},
		suffix {
			key: "delimiter",
			value: "/",
			flag: false,
		},
		suffix {
			key: "prefix",
			value: "0",
			flag: false,
		},
	} 
	req = &request {
		bucket: "bucket",
		object: "test.txt",
		suffixs: suffixs,
	}

	path = req.getCanonicalizedResource()
	if (path != "/bucket/test.txt") {
		t.Error("suffixs test error") 
	}

	path = req.getQueryString()
	if (path != "/bucket/test.txt?max-keys=1000&marker=1&delimiter=/&prefix=0") {
		t.Error("suffixs query error") 
	}
}

func TestGetContentMD5(t *testing.T) {
	req := &request {
		headers: http.Header{ },
	}
	md5 := req.getContentMD5() 
	if (md5 != empty) {
		t.Error("md5 faild.")
	}
	req = &request {
		headers: http.Header{
			"Content-MD5": {"123456"},
		},
	}
	md5 = req.getContentMD5() 
	if (md5 != "123456") {
		t.Error("md5 faild.")
	}
}

func TestGetContentType(t *testing.T) {
	req := &request {
		headers: http.Header {},
	}
	ct := req.getContentType()
	if ct != empty {
		t.Error("content type error")
	} 
	req = &request {
		headers: http.Header {
			"Content-Type": {"text/plain"},
		},
	}
	ct = req.getContentType()
	if ct != "text/plain" {
		t.Error("content type error")
	}
}

func TestGetCanonicalizedAmzHeaders(t *testing.T) {
	req := &request {}
	xamz := req.getCanonicalizedAmzHeaders()
	if empty != xamz {
		t.Error("AmzHeaders error")
	}
	req = &request {
		headers: http.Header {
			"X-Amz-Meta-ReviewedBy" : {"joe@johnsmith.net", "jane@johnsmith.net"},
			"X-Amz-Meta-FileChecksum": {"0x02661779"},
			"X-Amz-Meta-ChecksumAlgorithm": {"crc32"},
		},
	}
	xamz = req.getCanonicalizedAmzHeaders()
	if empty == xamz {
		t.Error("AmzHeaders error")
	}
	if xamz != "x-amz-meta-checksumalgorithm:crc32\nx-amz-meta-filechecksum:0x02661779\nx-amz-meta-reviewedby:joe@johnsmith.net,jane@johnsmith.net\n" {
		t.Error("AmzHeaders error")
	}
	req = &request {
		headers: http.Header {
			"x-amz-acl" : {"bucket-owner-full-control"},
		},
	}
	xamz = req.getCanonicalizedAmzHeaders()
	if empty == xamz {
		t.Error("AmzHeaders error")
	}
	if xamz != "x-amz-acl:bucket-owner-full-control\n" {
		t.Error("AmzHeaders error")
	}
}

func TestStringToSign(t *testing.T) {
	req := &request {}
	sign := req.stringToSign()
	if empty == sign {
		t.Error("sign error")
	} 
}

func TestSign(t *testing.T) {
	req := &request {}
	signValue := req.sign("accessKey", "secretAccessKey")
	if empty == signValue {
		t.Error("sign error")
	}  
}