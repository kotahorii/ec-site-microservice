package entity

import "github.com/microservice-ec-site/user/pkg/proto"

type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	CreatedAt string
	UpdatedAt string
}

func (u *User) ToProto() *proto.User {
	return &proto.User{
		Id:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func FromProtoToUser(p *proto.User) *User {
	return &User{
		ID:        p.Id,
		Email:     p.Email,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
