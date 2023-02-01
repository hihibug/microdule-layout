package config

import (
	"google.golang.org/grpc"
	"microdule-layout/internal/global"
)

//grpc 注册
func RegisterGrpc(s *grpc.Server) error {
	//pb.RegisterGrpcAppsServer(s, &grpc_server.GrpcCommon{})
	//pb.RegisterGrpcUploadsServer(s, &grpc_server.GrpcUploads{})

	//reflection.Register(s)
	register, err := global.Data.Rpc.Register(global.Data.Etcd.Clients())
	if err != nil {
		return err
	}
	go register.ListenLeaseRespChan()

	return nil
}
