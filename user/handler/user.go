package handler

import (
	"context"
	"go-micro-service/domain/model"
	service "go-micro-service/domain/service"
	pb "go-micro-service/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

// Register(context.Context, *UserRegisterRequest, *UserRegisterResponse) error
// 	Login(context.Context, *UserLoginRequest, *UserLoginResponse) error
// 	GetUserInfo(context.Context, *UserInfoRequest, *Use	rInfoResponse) error

// Register
func (u *User) Register(ctx context.Context, req *pb.UserRegisterRequest, res *pb.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     req.UserName,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		HashPassword: req.Passward,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	res.Message = " Register Success"
	return nil
}

func (u *User) Login(ctx context.Context, req *pb.UserLoginRequest, res *pb.UserLoginResponse) error {
	ok, err := u.UserDataService.CheckPwd(req.UserName, req.Password)
	if err != nil {
		return err
	}
	res.IsSuccess = ok
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, req *pb.UserInfoRequest, res *pb.UserInfoResponse) error {
	info, err := u.UserDataService.FindUserByName(req.UserName)
	if err != nil {
		return err
	}
	res = UserEntity(info)
	return nil
}

// user Info format
func UserEntity(user *model.User) *pb.UserInfoResponse {
	res := &pb.UserInfoResponse{
		UserId:    user.ID,
		UserName:  user.UserName,
		LastName:  user.LastName,
		FirstName: user.FirstName,
	}
	return res
}

// // Call is a single request handler called via client.Call or the generated client code
// func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
// 	log.Info("Received User.Call request")
// 	rsp.Msg = "Hello " + req.Name
// 	return nil
// }

// // Stream is a server side stream handler called via client.Stream or the generated client code
// func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.User_StreamStream) error {
// 	log.Infof("Received User.Stream request with count: %d", req.Count)

// 	for i := 0; i < int(req.Count); i++ {
// 		log.Infof("Responding: %d", i)
// 		if err := stream.Send(&user.StreamingResponse{
// 			Count: int64(i),
// 		}); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // PingPong is a bidirectional stream handler called via client.Stream or the generated client code
// func (e *User) PingPong(ctx context.Context, stream user.User_PingPongStream) error {
// 	for {
// 		req, err := stream.Recv()
// 		if err != nil {
// 			return err
// 		}
// 		log.Infof("Got ping %v", req.Stroke)
// 		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
// 			return err
// 		}
// 	}
// }
