package routers

import (
	"API-RS-TOUKIO/internal/interf/resource"
)

var routerUsers = []Router{
	{
		Path:              "/users",
		Method:            "POST",
		Handler:           resource.CreateUser,
		ReqAuthentication: false,
	},
}
