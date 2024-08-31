package myapp

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"testing"
)

func BenchmarkMarshalFlatbuffers(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)

	for i := 0; i < b.N; i++ {
		name := builder.CreateString("John Doe")
		email := builder.CreateString("john.doe@example.com")
		friends := []flatbuffers.UOffsetT{
			builder.CreateString("Alice"),
			builder.CreateString("Bob"),
		}

		UserStartFriendsVector(builder, len(friends))
		for i := len(friends) - 1; i >= 0; i-- {
			builder.PrependUOffsetT(friends[i])
		}
		friendsVector := builder.EndVector(len(friends))

		UserStart(builder)
		UserAddId(builder, 1)
		UserAddName(builder, name)
		UserAddAge(builder, 30)
		UserAddEmail(builder, email)
		UserAddIsActive(builder, true)
		UserAddUserType(builder, UserTypeRegularUser)
		UserAddPosition(builder, CreatePosition(builder, 12.34, 56.78))
		UserAddFriends(builder, friendsVector)
		user := UserEnd(builder)

		builder.Finish(user)
		builder.Reset()
	}
}

func BenchmarkUnmarshalFlatbuffers(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)

	name := builder.CreateString("John Doe")
	email := builder.CreateString("john.doe@example.com")
	friends := []flatbuffers.UOffsetT{
		builder.CreateString("Alice"),
		builder.CreateString("Bob"),
	}
	UserStartFriendsVector(builder, len(friends))
	for i := len(friends) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(friends[i])
	}
	friendsVector := builder.EndVector(len(friends))

	UserStart(builder)
	UserAddId(builder, 1)
	UserAddName(builder, name)
	UserAddAge(builder, 30)
	UserAddEmail(builder, email)
	UserAddIsActive(builder, true)
	UserAddUserType(builder, UserTypeRegularUser)
	UserAddPosition(builder, CreatePosition(builder, 12.34, 56.78))
	UserAddFriends(builder, friendsVector)
	user := UserEnd(builder)

	builder.Finish(user)
	data := builder.Bytes[builder.Head():]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := GetRootAsUser(data, 0)
		if buf == nil {
			b.Errorf("failed to unmarshal FlatBuffers")
		}
	}
}
