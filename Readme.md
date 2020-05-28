# 微服务示例
将proto文件编译为go文件

protoc --go_out=plugins=grpc:./test/ ./test.proto

搭建方法

环境：2018-05-12  protoc 3.5.1  go1.10.1 windows

gRPC: Google主导开发的RPC框架，这里不再赘述。

准备工作

先安装Protobuf 编译器 protoc，下载地址：https://github.com/google/protobuf/releases

我的是windows，将压缩包bin目录下的exe放到环境PATH目录中即可。

然后获取插件支持库

 // gRPC运行时接口编解码支持库

  go get -u github.com/golang/protobuf/proto

  // 从 Proto文件(gRPC接口描述文件) 生成 go文件 的编译器插件

  go get -u github.com/golang/protobuf/protoc-gen-go



获取go的gRPC包(网络问题可参阅https://www.jianshu.com/p/6392cb9dc38f)

go get google.golang.org/grpc



接口文件 /src/

新建test.proto示例：

syntax = "proto3";

  // 定义包名

  package test;

  

  // 可以定义多个服务，每个服务内可以定义多个接口

  service Waiter {

    // 定义接口 (结构体可以复用)

    // 方法 (请求消息结构体) returns (返回消息结构体) {}

    rpc DoMD5 (Req) returns (Res) {}

  }



  // 定义 Req 消息结构

  message Req {

    // 类型 字段 = 标识号

    string jsonStr = 1;

  }



  // 定义 Res 消息结构

  message Res {

    string backJson = 1;

  }

// PS：jsonStr和backJson只是随手写的名字，并没有用json



proto文件语法详解参阅：https://blog.csdn.net/u014308482/article/details/52958148

然后将proto文件编译为go文件

 // protoc --go_out=plugins=grpc:{输出目录}  {proto文件}

  protoc --go_out=plugins=grpc:./test/ ./test.proto

注意：原则上不要修改编译出来的*.bp.go文件的代码，因为双方接口基于同一个proto文件编译成自己的语言源码，此文件只作为接口数据处理，业务具体实现不在*.bp.go中。

服务端 /src/server/

本人也是刚接触Go，基于https://github.com/freewebsys/grpc-go-demo的Demo在修改中理解gRPC
其中中文注释均为个人理解笔记，若有不严谨的地方，还望指正。



package main

    import (

        "log"

        "net"

        "golang.org/x/net/context"

        "google.golang.org/grpc"

        "test"

        "google.golang.org/grpc/reflection"

        "fmt"

        "crypto/md5"

    )



    // 业务实现方法的容器

    type server struct{}



    // 为server定义 DoMD5 方法 内部处理请求并返回结果

    // 参数 (context.Context[固定], *test.Req[相应接口定义的请求参数])

    // 返回 (*test.Res[相应接口定义的返回参数，必须用指针], error)

    func (s *server) DoMD5(ctx context.Context, in *test.Req) (*test.Res, error) {

        fmt.Println("MD5方法请求JSON:"+in.JsonStr)

        return &test.Res{BackJson: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(in.JsonStr)))}, nil

    }



    func main() {

        lis, err := net.Listen("tcp", ":8028")  //监听所有网卡8028端口的TCP连接

        if err != nil {

            log.Fatalf("监听失败: %v", err)

        }

        s := grpc.NewServer() //创建gRPC服务



        /**注册接口服务

         * 以定义proto时的service为单位注册，服务中可以有多个方法

         * (proto编译时会为每个service生成Register***Server方法)

         * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])

         */

        test.RegisterWaiterServer(s, &server{})

        /**如果有可以注册多个接口服务,结构体要实现对应的接口方法

         * user.RegisterLoginServer(s, &server{})

         * minMovie.RegisterFbiServer(s, &server{})

         */

        // 在gRPC服务器上注册反射服务

        reflection.Register(s)

        // 将监听交给gRPC服务处理

        err = s.Serve(lis)

        if  err != nil {

            log.Fatalf("failed to serve: %v", err)

        }

    }

客户端 /src/client/

package main

        import (

            "log"

            "os"

            "golang.org/x/net/context"

            "google.golang.org/grpc"

            "test"

        )

    

    

        func main() {

            // 建立连接到gRPC服务

            conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())

            if err != nil {

                log.Fatalf("did not connect: %v", err)

            }

            // 函数结束时关闭连接

            defer conn.Close()

    

            // 创建Waiter服务的客户端

            t := test.NewWaiterClient(conn)

    

            // 模拟请求数据

            res := "test123"

            // os.Args[1] 为用户执行输入的参数 如：go run ***.go 123

            if len(os.Args) > 1 {

                res = os.Args[1]

            }

    

            // 调用gRPC接口

            tr, err := t.DoMD5(context.Background(), &test.Req{JsonStr: res})

            if err != nil {

                log.Fatalf("could not greet: %v", err)

            }

            log.Printf("服务端响应: %s", tr.BackJson)

        }


启动服务端监听，运行客户端即可达成远程调用