package entity

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/microservice-ec-site/user/pkg/proto"
)

type UpdateUserInput struct {
	ID        string
	Email     *string
	Password  *string
	FirstName *string
	LastName  *string
}

func (u *UpdateUserInput) ToProto() *proto.UpdateUserInput {
	return &proto.UpdateUserInput{
		Id:        u.ID,
		Email:     stringToWrapper(u.Email),
		Password:  stringToWrapper(u.Password),
		FirstName: stringToWrapper(u.FirstName),
		LastName:  stringToWrapper(u.LastName),
	}
}

func FromProtoToUpdateUserInput(p *proto.UpdateUserInput) *UpdateUserInput {
	return &UpdateUserInput{
		ID:        p.Id,
		Email:     wrapperToString(p.Email),
		Password:  wrapperToString(p.Password),
		FirstName: wrapperToString(p.FirstName),
		LastName:  wrapperToString(p.LastName),
	}
}

func stringToWrapper(s *string) *wrappers.StringValue {
	if s == nil {
		return nil
	}

	return &wrappers.StringValue{Value: *s}
}

func wrapperToString(w *wrappers.StringValue) *string {
	if w == nil {
		return nil
	}

	return &w.Value
}
