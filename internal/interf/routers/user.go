package routers

import (
	"API-RS-TOUKIO/internal/interf/resource"
)

var routerUsers = []Router{
	{
		Path:              "/users/create",
		Method:            "POST",
		Handler:           resource.CreateUser,
		ReqAuthentication: false,
	},
	{
		Path:              "/users/listALL",
		Method:            "GET",
		Handler:           resource.ListAllUsers,
		ReqAuthentication: false,
	},
	{
		Path:              "/users/listByNameOrNick",
		Method:            "GET",
		Handler:           resource.ListByNameOrNickUsers,
		ReqAuthentication: false,
	},
	{
		Path:              "/users/find/{userId}",
		Method:            "GET",
		Handler:           resource.FindUsers,
		ReqAuthentication: false,
	},
	{
		Path:              "/users/update/{userId}",
		Method:            "PUT",
		Handler:           resource.UpdateUser,
		ReqAuthentication: false,
	},
}
