
BINARY_NAME = RestaurantUserService



run:
	go run main.go

clean:
	rm -f $(BINARY_NAME)

generate-pb:
	protoc \
	  --go_out=. --go_opt=paths=source_relative \
	  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	  restaurantuserservicerpb/userservice.proto