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
package schema

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	InfoStat    *host.InfoStat         `json:"info_stat"`
	VMemStat    *mem.VirtualMemoryStat `json:"v_mem_stat"`
	SwapMemStat *mem.SwapMemoryStat    `json:"swap_mem_stat"`
	LoadStat    *load.AvgStat          `json:"load_stat"`
}
