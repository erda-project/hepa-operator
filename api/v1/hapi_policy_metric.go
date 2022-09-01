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

type Metric struct {
	Switch `json:",inline"`
}

func (in *Metric) SwitchOnForIngress(ctx context.Context, ingress *netv1.Ingress) error { /*todo: not implement*/
	return nil
}
func (in *Metric) SwitchOffForIngress(ctx context.Context, ingress *netv1.Ingress) error { /*todo: not implement*/
	return nil
}
func (in *Metric) SwitchOnForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
func (in *Metric) SwitcherOffForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
