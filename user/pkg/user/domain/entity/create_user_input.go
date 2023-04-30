package entity

import "github.com/microservice-ec-site/user/pkg/proto"

type CreateUserInput struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func (c *CreateUserInput) ToProto() *proto.CreateUserInput {
	return &proto.CreateUserInput{
		Email:     c.Email,
		Password:  c.Password,
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}
}

func FromProtoToCreateUserEntity(p *proto.CreateUserInput) *CreateUserInput {
	return &CreateUserInput{
		Email:     p.Email,
		Password:  p.Password,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
}
