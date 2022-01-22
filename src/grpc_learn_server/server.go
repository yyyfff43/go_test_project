package main

import (
"log"
"net"

"golang.org/x/net/context"
// 导入grpc包
"google.golang.org/grpc"
// 导入刚才我们生成的代码所在的proto包。
"google.golang.org/grpc/reflection"

pb "go_test_project/src/protobuf"
)

//生成pb.go时使用命令 protoc -I protobuf/ --go_out=plugins=grpc:protobuf protobuf/helloworld.proto 注意protobuf是包名

// 定义server，用来实现proto文件，里面实现的Greeter服务里面的接口
type server struct{}

// 实现SayHello接口
// 第一个参数是上下文参数，所有接口默认都要必填
// 第二个参数是我们定义的HelloRequest消息
// 返回值是我们定义的HelloReply消息，error返回值也是必须的。
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 创建一个HelloReply消息，设置Message字段，然后直接返回。
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 监听127.0.0.1:50051地址
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	pb.RegisterGreeterServer(s, &server{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}