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
	"strconv"

	netv1 "k8s.io/api/networking/v1"

	"github.com/erda-project/hepa-operator/pkg/middleware"
)

type ServerGuard struct {
	Switch `json:",inline"`

	RouteLimitRpm             int `json:"routeLimitRpm,omitempty"`
	RouteLimitRps             int `json:"routeLimitRps,omitempty"`
	RouteLimitBurstMultiplier int `json:"routeLimitBurstMultiplier,omitempty"`
}

func (in *ServerGuard) SwitchOnForIngress(ctx context.Context, ingress *netv1.Ingress) error {
	if in.RouteLimitRpm < 0 {
		in.RouteLimitRpm = 0
	}
	if in.RouteLimitRps < 0 {
		in.RouteLimitRps = 0
	}
	if in.RouteLimitBurstMultiplier == 0 {
		in.RouteLimitBurstMultiplier = 5
	}
	var (
		rpm   = strconv.Itoa(in.RouteLimitRpm)
		rps   = strconv.Itoa(in.RouteLimitRps)
		multi = strconv.Itoa(in.RouteLimitBurstMultiplier)
	)
	ingress.Annotations[RouteLimitBurstMultiplier.String()] = multi
	if in.RouteLimitRpm == 0 && in.RouteLimitRps == 0 {
		ingress.Annotations[RouteLimitRpm.String()] = rpm
		ingress.Annotations[RouteLimitRps.String()] = rps
		return nil
	}
	if in.RouteLimitRpm != 0 {
		ingress.Annotations[RouteLimitRpm.String()] = rpm
	}
	if in.RouteLimitRps != 0 {
		ingress.Annotations[RouteLimitRps.String()] = rps
	}
	return nil
}

func (in *ServerGuard) SwitchOffForIngress(ctx context.Context, ingress *netv1.Ingress) error {
	if len(ingress.Annotations) == 0 {
		return nil
	}
	for _, annotation := range []Annotation{
		RouteLimitRpm,
		RouteLimitRps,
		RouteLimitBurstMultiplier,
	} {
		delete(ingress.Annotations, annotation.String())
	}
	return nil
}

func (in *ServerGuard) SwitchOnForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
func (in *ServerGuard) SwitcherOffForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
