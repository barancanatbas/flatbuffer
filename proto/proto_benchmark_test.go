package proto

import (
	"flatbuffer-example/proto/myappProto"
	"google.golang.org/protobuf/proto"
	"testing"
)

func BenchmarkUnmarshalProtobuf(b *testing.B) {
	userProto := &myappProto.User{
		Id:       1,
		Name:     "John Doe",
		Age:      30,
		Email:    "john.doe@example.com",
		IsActive: true,
		UserType: myappProto.UserType_REGULAR_USER,
		Position: &myappProto.Position{Latitude: 12.34, Longitude: 56.78},
		Friends:  []string{"Alice", "Bob"},
	}
	data, err := proto.Marshal(userProto)
	if err != nil {
		b.Fatalf("failed to marshal Protobuf: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user myappProto.User
		err := proto.Unmarshal(data, &user)
		if err != nil {
			b.Errorf("failed to unmarshal Protobuf: %v", err)
		}
	}
}

func BenchmarkMarshalProtobuf(b *testing.B) {
	user := &myappProto.User{
		Id:       1,
		Name:     "John Doe",
		Age:      30,
		Email:    "john.doe@example.com",
		IsActive: true,
		UserType: myappProto.UserType_REGULAR_USER,
		Position: &myappProto.Position{Latitude: 12.34, Longitude: 56.78},
		Friends:  []string{"Alice", "Bob"},
	}

	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(user)
		if err != nil {
			b.Errorf("failed to marshal Protobuf: %v", err)
		}
	}
}
