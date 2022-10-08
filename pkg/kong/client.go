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

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Client struct {
	Version string `json:"version,omitempty"`
	Openapi string `json:"openapi,omitempty"`
}

func NewClient(openapi string) (*Client, error) {
	resp, err := http.Get(openapi)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.Errorf("invalid response: %s %s", resp.Status, string(data))
	}
	var info Info
	if err = json.Unmarshal(data, &info); err != nil {
		return nil, errors.Wrapf(err, "failed to Unmarshal kong info: %s", string(data))
	}
	return &Client{
		Version: info.Version,
		Openapi: openapi,
	}, nil
}

func (cli *Client) ServicePager() *ServicePager {
	return &ServicePager{
		cli:   cli,
		next:  "",
		valid: true,
	}
}

func (cli *Client) RoutePager() *RoutePager {
	return &RoutePager{
		cli:   cli,
		next:  "",
		valid: true,
	}
}

func (cli *Client) PluginPager() *PluginPager {
	return &PluginPager{
		cli:   cli,
		next:  "",
		valid: true,
	}
}

func (cli *Client) ServiceDeleter() *ServiceDeleter {
	return &ServiceDeleter{cli: cli}
}

func (cli *Client) RouteDeleter() *RouteDeleter {
	return &RouteDeleter{cli: cli}
}

func (cli *Client) PluginDeleter() *PluginDeleter {
	return &PluginDeleter{cli: cli}
}

func (cli *Client) IsV1() bool {
	return strings.HasPrefix(cli.Version, "1.")
}

func (cli *Client) IsV2() bool {
	return strings.HasPrefix(cli.Version, "2.")
}

type ServicePager struct {
	cli   *Client
	next  string
	valid bool
}

func (pager *ServicePager) Next() (page ServicePage, ok bool, err error) {
	defer func() {
		if err != nil {
			pager.next = ""
			pager.valid = false
		}
	}()
	if !pager.valid {
		return ServicePage{}, false, nil
	}
	var openapi = pager.cli.Openapi + "/services"
	if pager.next != "" {
		openapi = pager.cli.Openapi + pager.next
	}
	request, err := http.NewRequest(http.MethodGet, openapi, nil)
	if err != nil {
		return ServicePage{}, false, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return ServicePage{}, false, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ServicePage{}, false, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return ServicePage{}, false, errors.Errorf("response error: %s %s", response.Status, string(data))
	}

	if err = json.Unmarshal(data, &page); err != nil {
		return ServicePage{}, false, errors.Wrap(err, "failed to Unmarshal ServicePage")
	}
	if page.Next == nil || *page.Next == "" {
		pager.next = ""
		pager.valid = false
	} else {
		pager.next = *page.Next
		pager.valid = true
	}
	if len(page.Data) == 0 {
		return ServicePage{}, false, nil
	}
	return page, true, nil
}

type ServicePage struct {
	Next   *string   `json:"next,omitempty"`
	Data   []Service `json:"data,omitempty"`
	Offset *string   `json:"offset,omitempty"`
}

type RoutePager struct {
	cli   *Client
	next  string
	valid bool
}

func (pager *RoutePager) Next() (page RoutePage, ok bool, err error) {
	defer func() {
		if err != nil {
			pager.next = ""
			pager.valid = false
		}
	}()
	if !pager.valid {
		return RoutePage{}, false, nil
	}
	var openapi = pager.cli.Openapi + "/routes"
	if pager.next != "" {
		openapi = pager.cli.Openapi + pager.next
	}
	request, err := http.NewRequest(http.MethodGet, openapi, nil)
	if err != nil {
		return RoutePage{}, false, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return RoutePage{}, false, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return RoutePage{}, false, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return RoutePage{}, false, errors.Errorf("response error: %s %s", response.Status, string(data))
	}

	if err = json.Unmarshal(data, &page); err != nil {
		return RoutePage{}, false, errors.Wrap(err, "failed to Unmarshal RoutePage")
	}
	if page.Next == nil || *page.Next == "" {
		pager.next = ""
		pager.valid = false
	} else {
		pager.next = *page.Next
		pager.valid = true
	}
	if len(page.Data) == 0 {
		return RoutePage{}, false, nil
	}
	return page, true, nil
}

type RoutePage struct {
	Next *string `json:"next,omitempty"`
	Data []Route `json:"data,omitempty"`
}

type Info struct {
	Version string `json:"version"`
}

type PluginPager struct {
	cli   *Client
	next  string
	valid bool
}

func (pager *PluginPager) Next() (page PluginPage, ok bool, err error) {
	defer func() {
		if err != nil {
			pager.next = ""
			pager.valid = false
		}
	}()
	if !pager.valid {
		return PluginPage{}, false, nil
	}
	var openapi = pager.cli.Openapi + "/plugins"
	if pager.next != "" {
		openapi = pager.cli.Openapi + pager.next
	}
	request, err := http.NewRequest(http.MethodGet, openapi, nil)
	if err != nil {
		return PluginPage{}, false, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return PluginPage{}, false, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return PluginPage{}, false, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return PluginPage{}, false, errors.Errorf("response error: %s %s", response.Status, string(data))
	}

	if err = json.Unmarshal(data, &page); err != nil {
		return PluginPage{}, false, errors.Wrap(err, "failed to Unmarshal ServicePage")
	}
	if page.Next == nil || *page.Next == "" {
		pager.next = ""
		pager.valid = false
	} else {
		pager.next = *page.Next
		pager.valid = true
	}
	if len(page.Data) == 0 {
		return PluginPage{}, false, nil
	}
	return page, true, nil
}

type PluginPage struct {
	Next   *string  `json:"next,omitempty"`
	Data   []Plugin `json:"data,omitempty"`
	Offset *string  `json:"offset,omitempty"`
}

type ServiceDeleter struct {
	cli *Client
}

func (in *ServiceDeleter) Delete(id string) error {
	return deleteKongObject(in.cli.Openapi+"/services", id)
}

type RouteDeleter struct {
	cli *Client
}

func (in *RouteDeleter) Delete(id string) error {
	return deleteKongObject(in.cli.Openapi+"/routes", id)
}

type PluginDeleter struct {
	cli *Client
}

func (in *PluginDeleter) Delete(id string) error {
	return deleteKongObject(in.cli.Openapi+"/plugins", id)
}

func deleteKongObject(openapi, id string) error {
	if id == "" {
		return errors.New("empty id")
	}
	request, err := http.NewRequest(http.MethodDelete, openapi+"/"+id, nil)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		data, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		return errors.Errorf("response error: %s %s", response.Status, string(data))
	}
	return nil
}
