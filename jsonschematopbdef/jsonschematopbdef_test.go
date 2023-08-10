package jsonschematopbdef

import (
	"testing"
)

// 测试jsonschema转pbdef
func TestConvert(t *testing.T) {
	str := `{"$schema":"http://json-schema.org/draft-07/schema#","type":"object","properties":{"name":{"type":"string"},"age":{"type":"integer"}},"required":["name","age"]}`
	result, err := Convert(str)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", result)
}
