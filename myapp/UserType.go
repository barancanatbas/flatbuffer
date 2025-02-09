// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package myapp

import "strconv"

type UserType int8

const (
	UserTypeAdmin       UserType = 0
	UserTypeRegularUser UserType = 1
	UserTypeGuest       UserType = 2
)

var EnumNamesUserType = map[UserType]string{
	UserTypeAdmin:       "Admin",
	UserTypeRegularUser: "RegularUser",
	UserTypeGuest:       "Guest",
}

var EnumValuesUserType = map[string]UserType{
	"Admin":       UserTypeAdmin,
	"RegularUser": UserTypeRegularUser,
	"Guest":       UserTypeGuest,
}

func (v UserType) String() string {
	if s, ok := EnumNamesUserType[v]; ok {
		return s
	}
	return "UserType(" + strconv.FormatInt(int64(v), 10) + ")"
}
