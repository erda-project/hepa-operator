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

package kong

type Route struct {
	Id                      string       `json:"id"`
	CreatedAt               int          `json:"created_at"`
	UpdatedAt               int          `json:"updated_at"`
	Name                    string       `json:"name"`
	Protocols               []string     `json:"protocols"`
	Methods                 []string     `json:"methods"`
	Hosts                   []string     `json:"hosts"`
	Paths                   []string     `json:"paths"`
	Headers                 RouteHeaders `json:"headers"`
	HttpsRedirectStatusCode int          `json:"https_redirect_status_code"`
	RegexPriority           int          `json:"regex_priority"`
	StripPath               bool         `json:"strip_path"`
	PathHandling            string       `json:"path_handling"`
	PreserveHost            bool         `json:"preserve_host"`
	RequestBuffering        bool         `json:"request_buffering"`
	ResponseBuffering       bool         `json:"response_buffering"`
	Tags                    []string     `json:"tags"`
	Service                 RouteService `json:"service"`
}

type RouteHeaders struct {
	XMyHeader      []string `json:"x-my-header"`
	XAnotherHeader []string `json:"x-another-header"`
}

type RouteService struct {
	Id string `json:"id"`
}
