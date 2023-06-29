package json2pbdef

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bitly/go-simplejson"
)

func TestMessageFromJsonObject(t *testing.T) {
	data := `
{
  "hello": "world",
  "num":1,
  "num2":1.2,
  "array":[1,2], 
  "str_array":["a","b"],
  "obj": {
    "k1": "v1",
    "k2": 2,
    "k3": false
  }
}
`
	js, _ := simplejson.NewJson([]byte(data))
	code, err := messageFromJsonObject("Hello", js.MustMap(), 0, []string{})
	fmt.Printf("code %s, err:%v\n", strings.Join(code, "\n"), err)
}
