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

package v1_test

import (
	"testing"

	hapiv1 "github.com/erda-project/hepa-operator/api/v1"
)

func TestBackend_GetUpstreamHostName(t *testing.T) {
	var backend = hapiv1.Backend{
		RedirectBy:    "url",
		ServiceName:   "",
		ServicePort:   0,
		UpstreamHost:  "erda.cloud",
		RewriteTarget: "/s",
	}
	if backend.GetUpstreamHostName() != "erda.cloud" {
		t.Fatal("error host name")
	}
	if backend.GetUpstreamHostPort(0) != "" {
		t.Fatal("error port")
	}
	if port := backend.GetUpstreamHostPort(80); port != "80" {
		t.Fatalf("error port, expect: %s, got: %s", "80", port)
	}

	backend.UpstreamHost = "erda.cloud:80"
	if backend.GetUpstreamHostName() != "erda.cloud" {
		t.Fatal("error host name")
	}
	if backend.GetUpstreamHostPort(0) != "80" {
		t.Fatal("error port")
	}
}
