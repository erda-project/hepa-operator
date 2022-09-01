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
	"net/url"
	"strconv"
)

type Backend struct {
	RedirectBy    RedirectBy `json:"redirectBy,omitempty"`
	ServiceName   string     `json:"serviceName,omitempty"`
	ServicePort   int        `json:"servicePort,omitempty"`
	UpstreamHost  string     `json:"upstreamHost,omitempty"`
	RewriteTarget string     `json:"rewriteTarget,omitempty"`
}

func (in Backend) GetUpstreamHostName() string {
	return (&url.URL{Host: in.UpstreamHost}).Hostname()
}

func (in Backend) GetUpstreamHostPort(defaults int) string {
	if port := (&url.URL{Host: in.UpstreamHost}).Port(); port != "" {
		return port
	}
	if defaults > 0 {
		return strconv.Itoa(defaults)
	}
	return ""
}
