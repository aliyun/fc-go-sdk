package fc

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

//MD5 :Encoding MD5
func MD5(b []byte) string {
	ctx := md5.New()
	ctx.Write(b)
	return hex.EncodeToString(ctx.Sum(nil))
}

// HasPrefix check endpoint prefix
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// GetAccessPoint get correct endpoint and host
func GetAccessPoint(endpointInput string) (endpoint, host string) {
	unsecuredPrefix := "http://"
	securedPrefix := "https://"
	if HasPrefix(endpointInput, unsecuredPrefix) {
		host = endpointInput[len(unsecuredPrefix):]
		return endpointInput, host
	} else if HasPrefix(endpointInput, securedPrefix) {
		host = endpointInput[len(securedPrefix):]
		return endpointInput, host
	}
	return unsecuredPrefix + endpointInput, endpointInput
}

// IsBlank :check string pointer is nil or empty
func IsBlank(s *string) bool {
	if s == nil {
		return true
	}
	if len(*s) == 0 {
		return true
	}
	return false
}

type responseHeader struct {
	Header http.Header
}

func (h *responseHeader) GetRequestID() string {
	return h.Header.Get(HTTPHeaderRequestID)
}

func (h *responseHeader) GetErrorType() string {
	return h.Header.Get(HTTPHeaderFCErrorType)
}

func (h *responseHeader) GetEtag() string {
	return h.Header.Get(HTTPHeaderEtag)
}

type outputDecorator struct {
	responseHeader
}

func (o *outputDecorator) SetHeader(header http.Header) {
	o.SetHeader(header)
}

func (o outputDecorator) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func pathEscape(s string) string {
	return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}
