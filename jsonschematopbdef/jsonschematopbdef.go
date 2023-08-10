package jsonschematopbdef

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 将json schema转换为pbdef
func Convert(schema string) (pbdef string, err error) {
	//首先判断是否json-schema
	if !isJsonSchema(schema) {
		return "", errors.New("not json schema")
	}
	//转换
	decoder := json.NewDecoder(strings.NewReader(schema))

	var data map[string]interface{}
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//可以只针对properties进行处理
	properties := data["properties"]
	fmt.Println("message Root {")
	convertJsonSchema(properties.(map[string]interface{}))
	fmt.Println("}")
	return "", nil
}

// 转换json-schema，支持递归转换
func convertJsonSchema(schema map[string]interface{}) (pbdef string, err error) {
	count := 0
	for key, value := range schema {
		count++
		valueMap := value.(map[string]interface{})
		switch valueMap["type"] {
		case "object":
			//递归处理
			convertJsonSchema(valueMap["properties"].(map[string]interface{}))
		case "array":
			//递归处理
			convertJsonSchema(valueMap["items"].(map[string]interface{}))
		case "string":
			// 对string进行转换
			fmt.Println("string " + key + " = " + strconv.Itoa(count) + ";")
		case "number":
			// 对number进行转换
			fmt.Println("float " + key + " = " + strconv.Itoa(count) + ";")
		case "integer":
			// 对integer进行转换
			fmt.Println("int32 " + key + " = " + strconv.Itoa(count) + ";")
		default:
			//对value的各种类型进行转换
		}
	}
	return "", nil
}

// 判断是否json-schema
func isJsonSchema(schema string) (is bool) {
	var jsonData map[string]interface{}
	err := json.Unmarshal([]byte(schema), &jsonData)
	if err != nil {
		return false
	}
	// 检查 JSON Schema 的特定关键字是否存在
	// 这里只列举了一些常见的关键字，实际中可以根据需要扩展
	requiredKeywords := []string{
		// "$schema",
		// "title",
		// "description",
		"type",
		"properties",
		// "required",
	}

	for _, keyword := range requiredKeywords {
		_, exists := jsonData[keyword]
		if !exists {
			return false
		}
	}
	// 如果以上关键字都存在，则认为是 JSON Schema
	return true
}
