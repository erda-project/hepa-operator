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

type BaseStat struct {
	Hosts           string                    `json:"hosts,omitempty"`
	Path            string                    `json:"path,omitempty"`
	RedirectBy      RedirectBy                `json:"redirectBy,omitempty"`
	Policy          Policy                    `json:"policy,omitempty"`
	ResourceVersion HapiStatusResourceVersion `json:"resourceVersion"`
}

type RedirectByServiceStat struct {
	BaseStat `json:",inline"`

	ServiceName string `json:"serviceName,omitempty"`
	ServicePort int    `json:"servicePort,omitempty"`
}

type RedirectByUrlStat struct {
	BaseStat `json:",inline"`

	UpstreamHost  string `json:"upstreamHost,omitempty"`
	RewriteTarget string `json:"rewriteTarget,omitempty"`
}

func StatEqual(a1, a2 interface{}) ([]string, bool) {
	if a1 == nil || a2 == nil {
		return []string{"invalid stat"}, a1 == a1
	}
	return a1.(interface {
		DeepEqual(interface{}) ([]string, bool)
	}).DeepEqual(a2)
}
