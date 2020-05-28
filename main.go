package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_test/config"
	"grpc_test/controllers"
	_ "grpc_test/models"
	"grpc_test/proto"
	"log"
	"net"
)

func main() {
	appPort := config.Default("APP_PORT", "8888")
	lis, err := net.Listen("tcp", ":"+appPort)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer() //创建gRPC服务

	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */
	proto.RegisterWaiterServer(s, &controllers.Server{})
	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(s)

	fmt.Println("Now listening on: http://127.0.0.1:" + appPort)
	fmt.Println("Application started. Press CTRL+C to shut down.")
	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	/*客户端调用示例
	conn, _ := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	defer conn.Close()

	client := proto.NewWaiterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r1, err := client.DeletePddSessions(ctx, &proto.RequestDeletePddSessions{Id: 23})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		log.Printf("taoke is %+v", r1)
	}*/
}
