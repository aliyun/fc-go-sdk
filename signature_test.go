package fc

import (
	"testing"
)

const (
	expectedResource = "/path/action with space\n" +
		"foo=bar\n" +
		"key1=abc\n" +
		"key1=xyz\n" +
		"key2=123\n" +
		"key3/~x-y_z.a#b=value/~x-y_z.a#b\n" +
		"xyz"
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

	if resource != expectedResource {
		t.Fatalf("%s expected but %s in actual", expectedResource, resource)
	}
}
