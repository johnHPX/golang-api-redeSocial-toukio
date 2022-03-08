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
		ReqAuthentication: true,
	},
	{
		Path:              "/users/listByNameOrNick",
		Method:            "GET",
		Handler:           resource.ListByNameOrNickUsers,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/find/{userId}",
		Method:            "GET",
		Handler:           resource.FindUsers,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/update/{userId}",
		Method:            "PUT",
		Handler:           resource.UpdateUser,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/delete/{userId}",
		Method:            "DELETE",
		Handler:           resource.DeletarUser,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/login",
		Method:            "POST",
		Handler:           resource.LoginUser,
		ReqAuthentication: false,
	},
	{
		Path:              "/users/seguir/{userId}",
		Method:            "POST",
		Handler:           resource.SeguirUser,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/parar-seguir/{userId}",
		Method:            "POST",
		Handler:           resource.PararSeguirUser,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/seguidores/{userId}",
		Method:            "GET",
		Handler:           resource.ListSeguidoresUser,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/seguindo/{userId}",
		Method:            "GET",
		Handler:           resource.ListSeguindoUser,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/seguindo/{userId}",
		Method:            "GET",
		Handler:           resource.ListSeguindoUser,
		ReqAuthentication: true,
	},
	{
		Path:              "/users/update-password/{userId}",
		Method:            "POST",
		Handler:           resource.UpdatePasswordUser,
		ReqAuthentication: true,
	},
}
