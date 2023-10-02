// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Login "github.com/onlyLTY/oneKeyUpdate/UGREEN/internal/handler/Login"
	api "github.com/onlyLTY/oneKeyUpdate/UGREEN/internal/handler/api"
	containersManager "github.com/onlyLTY/oneKeyUpdate/UGREEN/internal/handler/containersManager"
	imagesManager "github.com/onlyLTY/oneKeyUpdate/UGREEN/internal/handler/imagesManager"
	version "github.com/onlyLTY/oneKeyUpdate/UGREEN/internal/handler/version"
	"github.com/onlyLTY/oneKeyUpdate/UGREEN/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: webindexHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.IndexCheckMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/login",
					Handler: Login.DoLoginHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/login",
					Handler: Login.LoginIndexHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CookieCheckMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: containersManager.ContainersManagerIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/start_container",
					Handler: containersManager.StartContainerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/stop_container",
					Handler: containersManager.StopContainerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/rename_container",
					Handler: containersManager.RenameContainerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create_container",
					Handler: containersManager.CreateContainerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/remove_container",
					Handler: containersManager.RemoveContainerHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/containersManager"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CookieCheckMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: imagesManager.ImagesManagerIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/get_new_image",
					Handler: imagesManager.GetNewImageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/remove_image",
					Handler: imagesManager.RemoveImageHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/imagesManager"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CookieCheckMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: version.VersionIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/get_version",
					Handler: version.GetVersionsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/updateprogram",
					Handler: version.UpdateprogramHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/checkprogramupdate",
					Handler: version.CheckprogramupdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/version"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.IndexCheckMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/auth",
					Handler: api.LoginHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)
}
