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
package internal

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"knife-panel/internal/app/schema"
)

func NewSystemMonitor() *SystemMonitor {
	return &SystemMonitor{}
}

type SystemMonitor struct {
}

func (a *SystemMonitor) GetSystemInfo() (*schema.SystemInfo, error) {
	systemInfo := schema.SystemInfo{}
	if infoStat, err := host.Info(); err != nil {
		return nil, err
	} else {
		systemInfo.InfoStat = infoStat
	}
	if vMemStat, err := mem.VirtualMemory(); err != nil {
		return nil, err
	} else {
		systemInfo.VMemStat = vMemStat
	}
	if swapMemStat, err := mem.SwapMemory(); err != nil {
		return nil, err
	} else {
		systemInfo.SwapMemStat = swapMemStat
	}
	if loadStat, err := load.Avg(); err != nil {
		return nil, err
	} else {
		systemInfo.LoadStat = loadStat
	}

	return &systemInfo, nil
}
