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
	"context"

	netv1 "k8s.io/api/networking/v1"

	"github.com/erda-project/hepa-operator/pkg/middleware"
)

type SafetySBAC struct {
	Switch `json:",inline"`

	AccessControlAPI string   `json:"accessControlAPI,omitempty"`
	Global           bool     `json:"global,omitempty"`
	Methods          []string `json:"methods,omitempty"`
	Patterns         []string `json:"patterns,omitempty"`
	WithBody         bool     `json:"withBody,omitempty"`
	WithCookie       bool     `json:"withCookie,omitempty"`
	WithHeaders      []string `json:"withHeaders,omitempty"`
}

func (in *SafetySBAC) SwitchOnForIngress(ctx context.Context, ingress *netv1.Ingress) error { /*todo: not implement*/
	return nil
}
func (in *SafetySBAC) SwitchOffForIngress(ctx context.Context, ingress *netv1.Ingress) error { /*todo: not implement*/
	return nil
}
func (in *SafetySBAC) SwitchOnForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
func (in *SafetySBAC) SwitcherOffForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
