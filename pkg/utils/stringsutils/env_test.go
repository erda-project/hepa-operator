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

package stringsutils_test

import (
	"testing"

	"github.com/erda-project/hepa-operator/pkg/utils/stringsutils"
)

func TestLoadEnv(t *testing.T) {
	t.Run("input a struct", func(t *testing.T) {
		type testStruct struct {
			Name string `env:"NAME"`
			Age  int    `env:"AGE"`
		}
		var c1 testStruct
		if err := stringsutils.LoadEnv(c1); err != nil {
			t.Log(err)
		} else {
			t.Fatal("err should not be nil")
		}
	})
	t.Run("input nil", func(t *testing.T) {
		if err := stringsutils.LoadEnv(nil); err != nil {
			t.Log(err)
		} else {
			t.Fatal("err should not be nil")
		}
	})
	t.Run("input nil interface", func(t *testing.T) {
		type testStruct struct {
			Name string `env:"NAME"`
			Age  int    `env:"AGE"`
		}
		var c *testStruct
		if err := stringsutils.LoadEnv(c); err != nil {
			t.Log(err)
		} else {
			t.Fatal("err should not be nil")
		}
	})
	t.Run("field type is not String", func(t *testing.T) {
		type testStruct struct {
			Name string `env:"NAME"`
			Age  int    `env:"AGE"`
		}
		var c testStruct
		if err := stringsutils.LoadEnv(&c); err != nil {
			t.Log(err)
		} else {
			t.Fatal("err should not be nil")
		}
	})
	t.Run("Normal", func(t *testing.T) {
		type testStruct struct {
			Name string `env:"NAME"`
			Age  string `env:"AGE"`
		}
		var c testStruct
		var result = testStruct{
			Name: "dspo",
			Age:  "29",
		}
		t.Setenv("NAME", result.Name)
		t.Setenv("AGE", result.Age)
		if err := stringsutils.LoadEnv(&c); err != nil {
			t.Fatal(err)
		}
		if c.Name != result.Name {
			t.Fatalf("expect .Name: %s, got: %s", result.Name, c.Name)
		}
		if c.Age != result.Age {
			t.Fatalf("expect .Age: %s, got: %s", result.Age, c.Age)
		}
	})
}
