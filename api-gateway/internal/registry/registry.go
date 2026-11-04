package registry

import (
	"os"

	"api-gateway/internal/constant"
	"api-gateway/internal/rpctypes"
)

type Route struct {
	Address  string
	RPCName  string
	NewArgs  func() any
	NewReply func() rpctypes.Result
}

func authAddr() string {
	if v := os.Getenv("AUTH_SERVICE_ADDR"); v != "" {
		return v
	}
	return "auth-service:1234"
}

func Lookup(service, method string) (Route, bool) {
	routes := map[string]map[string]Route{
		constant.ServiceAuth: {
			constant.MethodRegister: {
				Address:  authAddr(),
				RPCName:  "AuthService.Register",
				NewArgs:  func() any { return &rpctypes.RegisterArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.RegisterReply{} },
			},
			constant.MethodLogin: {
				Address:  authAddr(),
				RPCName:  "AuthService.Login",
				NewArgs:  func() any { return &rpctypes.LoginArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.LoginReply{} },
			},
		},
	}

	r, ok := routes[service][method]
	return r, ok
}
