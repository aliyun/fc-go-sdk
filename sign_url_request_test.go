package fc

import (
	"testing"
	"time"
)

func TestSignURL(t *testing.T) {
	input := NewSignURLInput("GET", "service", "func", time.Unix(1583339600, 0))
	url, err := input.signURL("2016-08-15", "http://localhost", "akid", "akSec", "stsToken")
	expectedURL := "http://localhost/2016-08-15/proxy/service/func/?x-fc-access-key-id=akid&x-fc-expires=1583339600&x-fc-security-token=stsToken&x-fc-signature=Yi0jCBQ6cTZc%2BtN8wtYUz9FbYNijx5jDiE3sVNW6XwA%3D"
	if url != expectedURL {
		t.Fatalf("%s expected but %s in actual due to %+v", expectedURL, url, err)
	}

	// make sure stsToken should be signed
	nURL, nErr := input.signURL("2016-08-15", "http://localhost", "akid", "akSec", "")
	nExpectedURL := "http://localhost/2016-08-15/proxy/service/func/?x-fc-access-key-id=akid&x-fc-expires=1583339600&x-fc-signature=6Rc4KxEAv3%2BlHpbNQWceIUqWVWq3ewxhadlJBRzZ0XQ%3D"
	if nURL != nExpectedURL {
		t.Fatalf("%s expected but %s in actual due to %+v", nExpectedURL, nURL, nErr)
	}
}

