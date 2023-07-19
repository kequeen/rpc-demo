package json2pbdata

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func testUnmarshalString() {
	// 定义 JSON 数据
	jsonData := `{"name": "John", "age": 30}`
	// 创建 结构体
	pbMessage := &MyMessage{}

	// 将 JSON 转换为 结构体
	err := jsonpb.UnmarshalString(jsonData, pbMessage)
	if err != nil {
		log.Fatalf("Error while unmarshaling JSON: %v", err)
	}
	// 打印结构体对象
	fmt.Println(pbMessage)
}

// 定义反序列化的golang结构体， 使用Protocol Buffers（protobuf）的标记格式指定了字段的序列化和反序列化规则。
type MyMessage struct {
	Name string `protobuf:"bytes,1,name=name"`
	Age  int32  `protobuf:"varint,2,name=age"`
}

// 实现 proto.Message 接口的方法
func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
