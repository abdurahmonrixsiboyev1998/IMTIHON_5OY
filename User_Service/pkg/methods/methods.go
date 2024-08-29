package methods

import (
	"context"
	"log"
	"user_service/internal/service"
	"user_service/pkg/protos/user"
)

type Service struct {
	user.UnimplementedUserServer
	S *service.Adjust
}

func (u *Service) Register(ctx context.Context, req *user.RegisterUserRequest) (*user.GeneralResponse, error) {
	r, err := u.S.AddUser(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}

func (u *Service) LogIn(ctx context.Context, req *user.LogInRequest) (*user.LogInResposne, error) {
	r, err := u.S.LogIn(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}

func (u *Service) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	r, err := u.S.GetUser(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}

func (u *Service) LastInserted(ctx context.Context, req *user.LastInsertedUser) (*user.GetUserResponse, error) {
	r, err := u.S.LastOne(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}

func (u *Service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.GeneralResponse, error) {
	r, err := u.S.Update(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}

func (u *Service) LogOut(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error) {
	r, err := u.S.Logout(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}

func (u *Service) DeleteUser(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error) {
	r, err := u.S.Delete(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
