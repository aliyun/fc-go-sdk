package fc

import (
	"net/http"
	"net/url"
)

// Credentials defines the returned credential
type Credentials struct {
	AccessKeyID     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Expiration      string `json:"Expiration"`
	SecurityToken   string `json:"SecurityToken"`
}

// GetTempBucketTokenOutput ...
type GetTempBucketTokenOutput struct {
	Header      http.Header `json:"header"`
	Credentials Credentials `json:"credentials"`
	OssRegion   string      `json:"ossRegion"`
	OssBucket   string      `json:"ossBucket"`
	ObjectName  string      `json:"objectName"`
}

// GetTempBucketTokenInput is empty
type GetTempBucketTokenInput struct{}

func (i GetTempBucketTokenInput) GetQueryParams() url.Values {
	return make(url.Values)
}

func (i GetTempBucketTokenInput) GetPath() string {
	return "/tempBucketToken"
}

func (i GetTempBucketTokenInput) GetHeaders() Header {
	return make(Header)
}

func (i GetTempBucketTokenInput) GetPayload() interface{} {
	return nil
}

func (i GetTempBucketTokenInput) Validate() error {
	return nil
}
