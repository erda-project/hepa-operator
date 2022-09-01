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

package hepa

type Interface interface {
	GetClusterInfo()
	GetGlobalPolicy(name string) ([]byte, error)
}

type mock struct {
}

func (m mock) GetClusterInfo() {

}

func (m mock) GetGlobalPolicy(name string) ([]byte, error) {
	return []byte(
		`mseAuth:
  authType: hmac-auth
  switch:
    switch: true
mseSafetyIP:
  blackListSourceRange: ""
  domainBlackListSourceRange: ""
  domainWhiteListSourceRange: ""
  ipType: x-peer-ip
  keyRateLimitingValue: 10 query_per_second
  switch:
	switch: true
  whiteListSourceRange: 123.45.67.1/16,10.10.10.10`), nil
}

func Mock() Interface {
	return mock{}
}
