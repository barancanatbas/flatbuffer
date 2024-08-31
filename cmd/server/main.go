package main

import (
	"context"
	"flatbuffer-example/myapp"
	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	myapp.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *myapp.UserRequest) (*flatbuffers.Builder, error) {
	builder := flatbuffers.NewBuilder(0)

	name := builder.CreateString("John Doe")
	email := builder.CreateString("john.doe@example.com")

	myapp.UserStart(builder)
	myapp.UserAddId(builder, req.Id())
	myapp.UserAddName(builder, name)
	myapp.UserAddAge(builder, 30)
	myapp.UserAddEmail(builder, email)
	myapp.UserAddIsActive(builder, true)
	user := myapp.UserEnd(builder)

	myapp.UserResponseStart(builder)
	myapp.UserResponseAddUser(builder, user)
	response := myapp.UserResponseEnd(builder)

	builder.Finish(response)

	return builder, nil
}

func (s *server) GetAllUsers(ctx context.Context, req *myapp.GetAllUsersRequest) (*flatbuffers.Builder, error) {
	builder := flatbuffers.NewBuilder(0)

	var users []flatbuffers.UOffsetT
	for i := 0; i < 10; i++ {
		name := builder.CreateString("John Doe")
		email := builder.CreateString("j")

		myapp.UserStart(builder)
		myapp.UserAddId(builder, int32(i))
		myapp.UserAddName(builder, name)
		myapp.UserAddAge(builder, 30)
		myapp.UserAddEmail(builder, email)
		myapp.UserAddIsActive(builder, true)
		users = append(users, myapp.UserEnd(builder))
	}

	myapp.UsersListResponseStartUsersVector(builder, len(users))
	for i := len(users) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(users[i])
	}

	usersList := builder.EndVector(len(users))

	myapp.UsersListResponseStart(builder)
	myapp.UsersListResponseAddUsers(builder, usersList)
	response := myapp.UsersListResponseEnd(builder)

	builder.Finish(response)

	return builder, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ForceServerCodec(flatbuffers.FlatbuffersCodec{}))
	myapp.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("gRPC server is running on port :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
