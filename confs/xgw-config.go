package confs

import "os"

func GetHttpAddr() string {

	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")

	return httpHost + ":" + httpPort
}

func GetGrpcAddr() string {

	grpcHost := os.Getenv("GRPC_HOST")
	grpcPort := os.Getenv("GRPC_PORT")

	return grpcHost + ":" + grpcPort
}
