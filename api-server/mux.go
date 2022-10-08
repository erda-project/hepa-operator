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

package api_server

import (
	"net/http"
)

var (
	Address = ":7999"
)

var (
	mux = http.NewServeMux()
)

func Handle(pattern string, handler http.Handler) {
	mux.Handle(pattern, handler)
}

func HandleFunc(pattern string, f http.HandlerFunc) {
	mux.HandleFunc(pattern, f)
}

func ListenAndServe() error {
	return (&http.Server{Addr: Address, Handler: mux}).ListenAndServe()
}
