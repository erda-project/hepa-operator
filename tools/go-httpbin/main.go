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

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	expose = os.Getenv("EXPOSE_PORT")
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.Infoln(r.Method, r.URL.String())
		if printHeaders := r.URL.Query().Get("print-headers"); strings.EqualFold(printHeaders, "true") {
			data, err := json.MarshalIndent(r.Header, "", "  ")
			if err != nil {
				logrus.WithError(err).Errorln("failed to json.MarshalIndent")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logrus.Println("print headers:", string(data))
		}
		if printQuery := r.URL.Query().Get("print-query"); strings.EqualFold(printQuery, "true") {
			data, err := json.MarshalIndent(r.URL.Query(), "", "  ")
			if err != nil {
				logrus.WithError(err).Errorln("failed to json.MarshalIndent")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logrus.Println("print query:", string(data))
		}
		if printBody := r.URL.Query().Get("print-body"); strings.EqualFold(printBody, "true") {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				logrus.WithError(err).Errorln("failed to io.ReadAll")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logrus.Println("print request body:", string(data))
		}
	})
	logrus.Infoln("ListenAndServe :" + expose)
	if err := http.ListenAndServe(":"+expose, nil); err != nil {
		panic(err)
	}
}
