package jsonxml

import (
	"encoding/xml"
	"testing"
)

func BenchmarkUnmarshalXML(b *testing.B) {
	data := `<User>
		<id>1</id>
		<name>John Doe</name>
		<age>30</age>
		<email>john.doe@example.com</email>
		<is_active>true</is_active>
		<user_type>1</user_type>
		<position>
			<latitude>12.34</latitude>
			<longitude>56.78</longitude>
		</position>
		<friends>Alice</friends>
		<friends>Bob</friends>
	</User>`

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user User
		err := xml.Unmarshal([]byte(data), &user)
		if err != nil {
			b.Errorf("failed to unmarshal XML: %v", err)
		}
	}
}

func BenchmarkMarshalXML(b *testing.B) {
	user := User{
		ID:       1,
		Name:     "John Doe",
		Age:      30,
		Email:    "john.doe@example.com",
		IsActive: true,
		UserType: RegularUser,
		Position: Position{Latitude: 12.34, Longitude: 56.78},
		Friends:  []string{"Alice", "Bob"},
	}

	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(user)
		if err != nil {
			b.Errorf("failed to marshal XML: %v", err)
		}
	}
}
