# json-schema转换成pb的定义
因为后续还需要自己进行部分的DSL，以及对于内容的解析，所以这部分短期来看，还是自己对json-schema进行解析比较合适
1、将json-schema 解析，在 golang 中使用使用 map[string]any 的格式来承载



## 参考文档
1. https://transform.tools/json-schema-to-protobuf
2. https://github.com/okdistribute/jsonschema-protobuf/tree/master