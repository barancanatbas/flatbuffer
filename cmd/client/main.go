package main

import (
	"context"
	"log"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"flatbuffer-example/myapp"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.ForceCodec(flatbuffers.FlatbuffersCodec{})))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := myapp.NewUserServiceClient(conn)

	for {
		getUser(client)

		getAllUsers(client)

		time.Sleep(1 * time.Second)
	}
}

func getUser(client myapp.UserServiceClient) {
	builder := flatbuffers.NewBuilder(1024)

	myapp.UserRequestStart(builder)
	myapp.UserRequestAddId(builder, 1)
	req := myapp.UserRequestEnd(builder)
	builder.Finish(req)

	res, err := client.GetUser(context.Background(), builder)
	if err != nil {
		log.Printf("Error while calling GetUser: %v", err)
	} else {
		user := res.User(nil)
		log.Printf("User ID: %d, User Name: %s, User Age: %d, User Email: %s, Is Active: %v\n",
			user.Id(), string(user.Name()), user.Age(), string(user.Email()), user.IsActive())
	}
}

func getAllUsers(client myapp.UserServiceClient) {
	getAllUserReqBuilder := flatbuffers.NewBuilder(1024)

	myapp.GetAllUsersRequestStart(getAllUserReqBuilder)
	getAllUserReq := myapp.GetAllUsersRequestEnd(getAllUserReqBuilder)
	getAllUserReqBuilder.Finish(getAllUserReq)

	userList, err := client.GetAllUsers(context.Background(), getAllUserReqBuilder)
	if err != nil {
		log.Printf("Error while calling GetAllUsers: %v", err)
	} else {
		users := make([]myapp.User, userList.UsersLength())
		for i := 0; i < userList.UsersLength(); i++ {
			userList.Users(&users[i], i)
		}

		for _, user := range users {
			log.Printf("User ID: %d, User Name: %s, User Age: %d, User Email: %s, Is Active: %v\n",
				user.Id(), string(user.Name()), user.Age(), string(user.Email()), user.IsActive())
		}
	}
}
