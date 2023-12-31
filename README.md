# 关于rpc
本项目是想用golang去构建一些常见的rpc场景，期望能够加深自己对于各种rpc场景的了解

## pbrpc
pb作为目前应用最为广泛的rpc通信编码格式

protoc --go_out=. --go-grpc_out=. demo.proto
在使用 Protocol Buffers 编译器生成 Golang 代码时，使用 --go-grpc_out 标志指定该插件的使用。该插件会读取 .proto 文件，并生成相应的 Golang 代码，包括 gRPC 服务端和客户端的代码。
具体而言，go-grpc_out 插件会生成以下内容：
服务端代码：生成一个服务实现的结构体，该结构体实现了由 .proto 文件中定义的 gRPC 服务接口。服务实现中包含接收 gRPC 请求的方法，以及对应的请求和响应消息类型。
客户端代码：生成一个用于与 gRPC 服务端进行通信的客户端代码。客户端代码提供了与服务端相匹配的方法，用于发送请求并接收响应。
注册代码：生成一个用于将服务实现注册到 gRPC 服务器的函数，通常是以 Register[ServiceName]Server 的形式命名。这个注册函数需要传入 gRPC 服务器对象和服务实现对象，以建立服务端的 RPC 服务。

生成代码需要安装protoc-gen-go 和 protoc-gen-go-grpc 插件

### 关于pb
全称 Protocol Buffers，
```
Protocol Buffers（简称：ProtoBuf）是一种开源跨平台的序列化数据结构的协议。其对于存储资料或在网络上进行通信的程序是很有用的。这个方法包含一个接口描述语言，描述一些数据结构，并提供程序工具根据这些描述产生代码，这些代码将用来生成或解析代表这些数据结构的字节流。
``` 
