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
package ctl

import (
	"github.com/gin-gonic/gin"
	"knife-panel/internal/app/bll"
	"knife-panel/internal/app/ginplus"
)

func NewSystemMonitor(monitor bll.ISystemMonitor) *SystemMonitor {
	return &SystemMonitor{
		SystemMonitorBll: monitor,
	}
}

type SystemMonitor struct {
	SystemMonitorBll bll.ISystemMonitor
}

func (a *SystemMonitor) GetSystemInfo(c *gin.Context) {
	result, err := a.SystemMonitorBll.GetSystemInfo()
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResSuccess(c, result)
}
