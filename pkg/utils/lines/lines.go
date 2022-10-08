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

package lines

import (
	"strings"
)

type Lines struct {
	lines []string
}

func From(s string) *Lines {
	return &Lines{lines: strings.Split(s, "\n")}
}

func (l *Lines) Set(line string) *Lines {
	l.Delete(line)
	l.lines = append([]string{line}, l.lines...)
	return l
}

func (l *Lines) Delete(line string) *Lines {
	for i := 0; i < len(l.lines); i++ {
		if l.lines[i] == line {
			l.lines = append(l.lines[:i], l.lines[i+1:]...)
		}
	}
	return l
}

func (l *Lines) String() string {
	return strings.Join(l.lines, "\n")
}
