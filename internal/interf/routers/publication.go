package routers

import "API-RS-TOUKIO/internal/interf/resource"

var routerPublication = []Router{
	{
		Path:              "/publication/create",
		Method:            "POST",
		Handler:           resource.CreatePublication,
		ReqAuthentication: true,
	},
	{
		Path:              "/publication/listALL",
		Method:            "GET",
		Handler:           resource.ListAllPublication,
		ReqAuthentication: true,
	},
	{
		Path:              "/publication/find/{publicationID}",
		Method:            "GET",
		Handler:           resource.FindByIDPublication,
		ReqAuthentication: true,
	},
	{
		Path:              "/publication/update/{publicationID}",
		Method:            "PUT",
		Handler:           resource.UpdatePublication,
		ReqAuthentication: true,
	},
	{
		Path:              "/publication/delete/{publicationID}",
		Method:            "DELETE",
		Handler:           resource.DeletePublication,
		ReqAuthentication: true,
	},
	{
		Path:              "/publication/users/{userID}",
		Method:            "GET",
		Handler:           resource.ListByIDUserPublication,
		ReqAuthentication: true,
	},
	{
		Path:              "/publication/like/{publicationID}",
		Method:            "POST",
		Handler:           resource.LikePublication,
		ReqAuthentication: true,
	},
	{
		Path:              "/publication/deslike/{publicationID}",
		Method:            "POST",
		Handler:           resource.DeslikePublication,
		ReqAuthentication: true,
	},
}
