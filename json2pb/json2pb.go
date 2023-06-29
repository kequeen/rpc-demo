package json2pbdata

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	// 定义 JSON 数据
	jsonData := `{"name": "John", "age": 30}`

	// 创建 PB 消息对象
	pbMessage := &MyMessage{}

	// 将 JSON 转换为 PB
	err := jsonpb.UnmarshalString(jsonData, pbMessage)
	if err != nil {
		log.Fatalf("Error while unmarshaling JSON: %v", err)
	}

	// 打印 PB 对象
	fmt.Println(pbMessage)
}

// 定义 PB 消息类型
type MyMessage struct {
	Name string `protobuf:"bytes,1,name=name"`
	Age  int32  `protobuf:"varint,2,name=age"`
}

// 实现 proto.Message 接口的方法
func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
