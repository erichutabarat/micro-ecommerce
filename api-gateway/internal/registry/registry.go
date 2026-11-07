package registry

import (
	"os"

	"api-gateway/internal/constant"
	"api-gateway/internal/rpctypes"
)

type Route struct {
	Address  string
	RPCName  string
	Role     string
	NewArgs  func() any
	NewReply func() rpctypes.Result
}

func authAddr() string {
	if v := os.Getenv("AUTH_SERVICE_ADDR"); v != "" {
		return v
	}
	return "auth-service:1234"
}

func productAddr() string {
	if v := os.Getenv("PRODUCT_SERVICE_ADDR"); v != "" {
		return v
	}
	return "product-service:1235"
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
		constant.ServiceSeller: {
			constant.MethodRegister: {
				Address:  authAddr(),
				RPCName:  "SellerService.Register",
				NewArgs:  func() any { return &rpctypes.SellerRegisterArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.SellerRegisterReply{} },
			},
			constant.MethodLogin: {
				Address:  authAddr(),
				RPCName:  "SellerService.Login",
				NewArgs:  func() any { return &rpctypes.SellerLoginArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.SellerLoginReply{} },
			},
		},
		constant.ServiceProduct: {
			constant.MethodList: {
				Address:  productAddr(),
				RPCName:  "ProductService.List",
				NewArgs:  func() any { return &rpctypes.ListProductsArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.ListProductsReply{} },
			},
			constant.MethodGet: {
				Address:  productAddr(),
				RPCName:  "ProductService.Get",
				NewArgs:  func() any { return &rpctypes.GetProductArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.ProductReply{} },
			},
			constant.MethodCreate: {
				Address:  productAddr(),
				RPCName:  "ProductService.Create",
				Role:     "seller",
				NewArgs:  func() any { return &rpctypes.CreateProductArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.ProductReply{} },
			},
			constant.MethodListMine: {
				Address:  productAddr(),
				RPCName:  "ProductService.ListMine",
				Role:     "seller",
				NewArgs:  func() any { return &rpctypes.ListMyProductsArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.ListProductsReply{} },
			},
			constant.MethodUpdate: {
				Address:  productAddr(),
				RPCName:  "ProductService.Update",
				Role:     "seller",
				NewArgs:  func() any { return &rpctypes.UpdateProductArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.ProductReply{} },
			},
			constant.MethodDelete: {
				Address:  productAddr(),
				RPCName:  "ProductService.Delete",
				Role:     "seller",
				NewArgs:  func() any { return &rpctypes.DeleteProductArgs{} },
				NewReply: func() rpctypes.Result { return &rpctypes.DeleteProductReply{} },
			},
		},
	}

	r, ok := routes[service][method]
	return r, ok
}
