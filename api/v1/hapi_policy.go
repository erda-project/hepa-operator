// Copyright (c) 2022 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package v1

import (
	"reflect"
	"strings"

	"github.com/erda-project/hepa-operator/pkg/interfaces"
)

type Policy struct {
	Auth        Auth        `json:"auth,omitempty"`
	CORS        CORS        `json:"cors,omitempty"`
	Metric      Metric      `json:"metric,omitempty"`
	Proxy       Proxy       `json:"proxy,omitempty"`
	SafetyCSRF  SafetyCSRF  `json:"safetyCSRF,omitempty"`
	SafetyIP    SafetyIP    `json:"safetyIP,omitempty"`
	SafetySBAC  SafetySBAC  `json:"safetySBAC,omitempty"`
	ServerGuard ServerGuard `json:"serverGuard,omitempty"`
}

type Switch struct {
	Global bool `json:"global,omitempty"`
	Switch bool `json:"switch,omitempty"`
}

func (in *Switch) GetGlobal() bool {
	return in.Global
}

func (in *Switch) GetSwitch() bool {
	return in.Switch
}

func (in *Policy) ListAll() map[string]interface {
	interfaces.Switcher
	interfaces.GlobalGetter
} {
	var m = make(map[string]interface {
		interfaces.Switcher
		interfaces.GlobalGetter
	})
	tof := reflect.TypeOf(*in)
	vof := reflect.ValueOf(in)
	for i := 0; i < tof.NumField(); i++ {
		jTag := tof.Field(i).Tag.Get("json")
		jTag = jTag[:strings.Index(jTag, ",")]
		v := vof.Elem().Field(i).Addr().Interface().(interface {
			interfaces.Switcher
			interfaces.GlobalGetter
		})
		m[jTag] = v
	}
	return m
}
