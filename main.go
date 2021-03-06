package main

import (
	"bdev/models"
	"bdev/controllers"
	"bdev/logger"
	_ "bdev/routers"
	_ "github.com/lib/pq"

	"bdev/config"
	pb "bdev/protos"
	"github.com/astaxie/beego"
	"sync"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
)

func init() {
	beego.BConfig.AppName = config.AppConf.AppName
	beego.BConfig.RunMode = config.AppConf.RunMode
	beego.BConfig.CopyRequestBody = config.AppConf.CopyRequestBody
	beego.BConfig.Listen.HTTPPort = config.AppConf.HttpPort
	beego.BConfig.WebConfig.EnableDocs = config.AppConf.EnableDocs
	beego.BConfig.WebConfig.AutoRender = config.AppConf.AutoRender
	models.InitDB()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		grpcStart()
		wg.Done()
	}()
	wg.Wait()
	logger.Debug("ser done")
}

func grpcStart() {
	lis, err := net.Listen("tcp", config.AppConf.GrpcListen)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	logger.Debugf("grpc-port(%v)", config.AppConf.GrpcListen)
	pb.RegisterDeveloperServiceServer(s, &controllers.DeveloperServiceController{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}


