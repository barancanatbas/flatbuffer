package jsonxml

type UserType int32

const (
	Admin       UserType = 0
	RegularUser UserType = 1
	Guest       UserType = 2
)

type Position struct {
	Latitude  float32 `json:"latitude" xml:"latitude"`
	Longitude float32 `json:"longitude" xml:"longitude"`
}

type User struct {
	ID       int32    `json:"id" xml:"id"`
	Name     string   `json:"name" xml:"name"`
	Age      int64    `json:"age" xml:"age"`
	Email    string   `json:"email" xml:"email"`
	IsActive bool     `json:"is_active" xml:"is_active"`
	UserType UserType `json:"user_type" xml:"user_type"`
	Position Position `json:"position" xml:"position"`
	Friends  []string `json:"friends" xml:"friends"`
}
