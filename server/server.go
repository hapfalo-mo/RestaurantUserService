package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"restaurantuserservice/restaurantuserservicerpb"
)

type userserver struct {
	restaurantuserservicerpb.UnimplementedRestaurantUserServiceServer
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50016")
	if err != nil {
		log.Fatalf("Failed to listen server : %v", err)
	}
	s := grpc.NewServer()
	restaurantuserservicerpb.RegisterRestaurantUserServiceServer(s, &userserver{})
	log.Println("Service is running in port : 50016")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start server : %v", err)
	}
}
