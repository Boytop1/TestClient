package main

import (
	"log"
	"net"
	"testClient/msg"

	"google.golang.org/protobuf/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 创建一个消息 Test
	test := &msg.Hello{
		// 使用辅助函数设置域的值
		Name: proto.String("hello"),
	}

	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// 进行解码
	newTest := &msg.Hello{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// 测试结果
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}

	// 发送消息
	conn.Write(data)
}
