package fc

import (
	"testing"
)

func TestGetSignResourceWithQueries(t *testing.T) {
	path := "/path/action with space"
	queries := map[string][]string{
		"xyz":             {},
		"foo":             {"bar"},
		"key2":            {"123"},
		"key1":            {"xyz", "abc"},
		"key3/~x-y_z.a#b": {"value/~x-y_z.a#b"},
	}
	resource := GetSignResourceWithQueries(path, queries)

	expectedResource := "/path/action with space\nfoo=bar\nkey1=abc\nkey1=xyz\nkey2=123\nkey3/~x-y_z.a#b=value/~x-y_z.a#b\nxyz"
	if resource != expectedResource {
		t.Fatalf("%s expected but %s in actual", expectedResource, resource)
	}

}

func TestGetSignature(t *testing.T) {
	path := "/2016-08-15/proxy/service.LATEST/func/abc 123"
	queries := map[string][]string{
		"x":                   {"xyz", "456"},
		"a":                   {"123 45"},
		"x-fc-expires":        {"1583221068"},
		"x-fc-access-key-id":  {"akID"},
		"x-fc-security-token": {"stsToken"},
	}

	headers := map[string]string{
		"x-fc-trace-id": "trace-id",
		"content-type":  "text/json",
		"content-md5":   "ef5b43a0fe1d5b401b62c3ba33a7a3d6",
	}

	mockAkSec := "S9MWOBG0AOWHSO9HPASP216QOPHT5YR4NLH3A"
	fcResource := GetSignResourceWithQueries(path, queries)
	signature := GetSignature(mockAkSec, "GET", headers, fcResource)

	expectedSign := "QFdNhz0Dl+onUM/MvpX940WFH3H8OHayZHbby1c01aI="
	expectedSignStr := "GET\nef5b43a0fe1d5b401b62c3ba33a7a3d6\ntext/json\n1583221068\nx-fc-trace-id:trace-id\n/2016-08-15/proxy/service.LATEST/func/abc 123\na=123 45\nx-fc-access-key-id=akID\nx-fc-expires=1583221068\nx-fc-security-token=stsToken\nx=456\nx=xyz"

	if signature != expectedSign {
		t.Fatalf("%s expected but %s in actual, signStr:%s", expectedSign, signature, expectedSignStr)
	}
}
