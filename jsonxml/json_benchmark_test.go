package jsonxml

import (
	"encoding/json"
	"testing"
)

func BenchmarkUnmarshalJSON(b *testing.B) {
	data := `{
		"id": 1,
		"name": "John Doe",
		"age": 30,
		"email": "john.doe@example.com",
		"is_active": true,
		"user_type": 1,
		"position": {
			"latitude": 12.34,
			"longitude": 56.78
		},
		"friends": ["Alice", "Bob"]
	}`

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user User
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			b.Errorf("failed to unmarshal JSON: %v", err)
		}
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
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
		_, err := json.Marshal(user)
		if err != nil {
			b.Errorf("failed to marshal JSON: %v", err)
		}
	}
}
