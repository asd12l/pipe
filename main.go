// Solo.go - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/b3log/solo.go/controller"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var Version = "1.0.0"

// The only one init function in Solo.
func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	util.InitConf()
}

// Entry point.
func main() {
	service.ConnectDB()

	router := controller.MapRoutes()
	server := &http.Server{
		Addr:    util.Conf.Server,
		Handler: router,
	}

	handleSignal(server)

	log.Infof("Solo.go (v%s) is running [%s]", Version, util.Conf.Server)

	server.ListenAndServe()
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		log.Infof("got signal [%s], exiting Solo.go now", s)
		if err := server.Close(); nil != err {
			log.Error("server close failed: ", err)
		}

		service.DisconnectDB()

		log.Info("Solo exited")
		os.Exit(0)
	}()
}
