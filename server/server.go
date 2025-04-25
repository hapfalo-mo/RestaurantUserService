package server

import (
	"context"
	"github.com/hapfalo-mo/RestaurantUserService/restaurantuserservicerpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type userserver struct {
	restaurantuserservicerpb.UnimplementedRestaurantUserServiceServer
}

func Sum(ctx context.Context, req *restaurantuserservicerpb.SumRequest) (res *restaurantuserservicerpb.SumReponse, err error) {
	log.Println("Processing in Sum Method")
	res = &restaurantuserservicerpb.SumReponse{
		Result: req.GetNum1() + req.GetNum2(),
	}
	return res, nil
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
