// Copyright 2019 liuxiaodong Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ws

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"knife-panel/backend/localcommand"
	"knife-panel/internal/app/routers/api/ctl"
	"knife-panel/pkg/auth"
	"knife-panel/server"
	"knife-panel/utils"
	"os"
)

func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	return container.Invoke(func(
		a auth.Auther,
		e *casbin.SyncedEnforcer,
		cDemo *ctl.Demo,
		cLogin *ctl.Login,
		cMenu *ctl.Menu,
		cRole *ctl.Role,
		cUser *ctl.User,
	) error {
		g := app.Group("/ws")
		tty := g.Group("/tty")
		{
			tty.GET("/", func(ctx *gin.Context) {

				appOptions := &server.Options{
					PermitWrite: true,
					Width:       600,
					Height:      800,
				}
				if err := utils.ApplyDefaultValues(appOptions); err != nil {
					exit(err, 1)
				}
				backendOptions := &localcommand.Options{}
				if err := utils.ApplyDefaultValues(backendOptions); err != nil {
					exit(err, 1)
				}

				factory, err := localcommand.NewFactory("/bin/bash", nil, backendOptions)
				if err != nil {
					exit(err, 3)
				}
				gCtx, _ := context.WithCancel(context.Background())
				srv, err := server.New(factory, appOptions)
				if err != nil {
					exit(err, 3)
				}
				srv.Run(gCtx, ctx, server.WithGracefullContext(gCtx))
			})
		}
		return nil
	})
}
func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}
