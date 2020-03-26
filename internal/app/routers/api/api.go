package api

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"knife-panel/backend/localcommand"
	"knife-panel/internal/app/middleware"
	"knife-panel/internal/app/routers/api/ctl"
	"knife-panel/pkg/auth"
	"knife-panel/server"
	"knife-panel/utils"
	"knife-panel/webtty"
	"net/http"
	"os"
)

// RegisterRouter 注册/api路由
func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	err := ctl.Inject(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(
		a auth.Auther,
		e *casbin.SyncedEnforcer,
		cDemo *ctl.Demo,
		cLogin *ctl.Login,
		cMenu *ctl.Menu,
		cRole *ctl.Role,
		cUser *ctl.User,
	) error {

		g := app.Group("/api")

		// 用户身份授权
		g.Use(middleware.UserAuthMiddleware(a,
			middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
			middleware.AllowPathPrefixSkipper("/api/ws"),
		))

		// casbin权限校验中间件
		g.Use(middleware.CasbinMiddleware(e,
			middleware.AllowPathPrefixSkipper("/api/v1/pub"),
			middleware.AllowPathPrefixSkipper("/api/ws"),
		))

		// 请求频率限制中间件
		g.Use(middleware.RateLimiterMiddleware())

		v1 := g.Group("/v1")
		{
			pub := v1.Group("/pub")
			{
				// 注册/api/v1/pub/login
				gLogin := pub.Group("login")
				{
					gLogin.GET("captchaid", cLogin.GetCaptcha)
					gLogin.GET("captcha", cLogin.ResCaptcha)
					gLogin.POST("", cLogin.Login)
					gLogin.POST("exit", cLogin.Logout)
				}

				// 注册/api/v1/pub/refresh-token
				pub.POST("/refresh-token", cLogin.RefreshToken)

				// 注册/api/v1/pub/current
				gCurrent := pub.Group("current")
				{
					gCurrent.PUT("password", cLogin.UpdatePassword)
					gCurrent.GET("user", cLogin.GetUserInfo)
					gCurrent.GET("menutree", cLogin.QueryUserMenuTree)
				}

			}

			// 注册/api/v1/demos
			gDemo := v1.Group("demos")
			{
				gDemo.GET("", cDemo.Query)
				gDemo.GET(":id", cDemo.Get)
				gDemo.POST("", cDemo.Create)
				gDemo.PUT(":id", cDemo.Update)
				gDemo.DELETE(":id", cDemo.Delete)
				gDemo.PATCH(":id/enable", cDemo.Enable)
				gDemo.PATCH(":id/disable", cDemo.Disable)
			}

			// 注册/api/v1/menus
			gMenu := v1.Group("menus")
			{
				gMenu.GET("", cMenu.Query)
				gMenu.GET(":id", cMenu.Get)
				gMenu.POST("", cMenu.Create)
				gMenu.PUT(":id", cMenu.Update)
				gMenu.DELETE(":id", cMenu.Delete)
			}
			v1.GET("/menus.tree", cMenu.QueryTree)

			// 注册/api/v1/roles
			gRole := v1.Group("roles")
			{
				gRole.GET("", cRole.Query)
				gRole.GET(":id", cRole.Get)
				gRole.POST("", cRole.Create)
				gRole.PUT(":id", cRole.Update)
				gRole.DELETE(":id", cRole.Delete)
			}
			v1.GET("/roles.select", cRole.QuerySelect)

			// 注册/api/v1/users
			gUser := v1.Group("users")
			{
				gUser.GET("", cUser.Query)
				gUser.GET(":id", cUser.Get)
				gUser.POST("", cUser.Create)
				gUser.PUT(":id", cUser.Update)
				gUser.DELETE(":id", cUser.Delete)
				gUser.PATCH(":id/enable", cUser.Enable)
				gUser.PATCH(":id/disable", cUser.Disable)
			}
		}

		ws := g.Group("/ws")
		{
			tty := ws.Group("/tty")
			{
				tty.GET("/", func(ctx *gin.Context) {
					upgrader := &websocket.Upgrader{
						ReadBufferSize:  1024,
						WriteBufferSize: 1024,
						Subprotocols:    webtty.Protocols,
						CheckOrigin: func(r *http.Request) bool {
							return true
						},
					}
					conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
					if err != nil {
						fmt.Println(err)
						return
					}
					defer conn.Close()
					err = processWSConn(ctx, conn)
					var closeReason string
					switch err {
					case ctx.Err():
						closeReason = "cancelation"
					case webtty.ErrSlaveClosed:
						closeReason = "err slave closed"
					case webtty.ErrMasterClosed:
						closeReason = "client"
					default:
						closeReason = fmt.Sprintf("an error: %s", err)
					}
					fmt.Println(closeReason)
					//ctx.Writer.WriteString(closeReason)
				})
			}
		}

		return nil
	})

}
func processWSConn(ctx context.Context, conn *websocket.Conn) error {
	//typ, initLine, err := conn.ReadMessage()
	//if err != nil {
	//	return errors.Wrapf(err, "failed to authenticate websocket connection")
	//}
	//if typ != websocket.TextMessage {
	//	return errors.New("failed to authenticate websocket connection: invalid message type")
	//}
	appOptions := &server.Options{}
	if err := utils.ApplyDefaultValues(appOptions); err != nil {
		os.Exit(1)
	}
	backendOptions := &localcommand.Options{}
	if err := utils.ApplyDefaultValues(backendOptions); err != nil {
		os.Exit(1)
	}
	factory, err := localcommand.NewFactory("bash", nil, backendOptions)
	if err != nil {
		os.Exit(3)
	}

	slave, err := factory.New(nil)
	if err != nil {
		return errors.Wrapf(err, "failed to create backend")
	}
	defer slave.Close()

	//titleBuf := new(bytes.Buffer)
	//err = server.titleTemplate.Execute(titleBuf, titleVars)
	//if err != nil {
	//	return errors.Wrapf(err, "failed to fill window title template")
	//}

	opts := []webtty.Option{
		webtty.WithWindowTitle([]byte("hello world")),
	}
	//if server.options.PermitWrite {
	//	opts = append(opts, webtty.WithPermitWrite())
	//}
	//if server.options.EnableReconnect {
	//	opts = append(opts, webtty.WithReconnect(server.options.ReconnectTime))
	//}
	//if server.options.Width > 0 {
	//	opts = append(opts, webtty.WithFixedColumns(server.options.Width))
	//}
	//if server.options.Height > 0 {
	//	opts = append(opts, webtty.WithFixedRows(server.options.Height))
	//}
	//if server.options.Preferences != nil {
	//	opts = append(opts, webtty.WithMasterPreferences(server.options.Preferences))
	//}

	tty, err := webtty.New(&server.WsWrapper{conn}, slave, opts...)
	if err != nil {
		return errors.Wrapf(err, "failed to create webtty")
	}

	err = tty.Run(ctx)

	return err
}
