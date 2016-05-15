// Copyright (c) 2016, Gerasimos Maropoulos
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//	  this list of conditions and the following disclaimer
//    in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse
//    or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER AND CONTRIBUTOR, GERASIMOS MAROPOULOS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package main v0.0.1
package main

import (
	"os"

	"github.com/kataras/cli"
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/utils"
)

func main() {
	app := cli.NewApp("httpserver", "converts any directory into a fast http(s) website", "0.0.1")
	app.Flag("dir", "."+string(os.PathSeparator), "the current working directory")
	app.Flag("gzip", false, "enable serving with gzip compression & use file cache")
	app.Flag("host", "0.0.0.0:8080", "the server addr to listen for")
	app.Flag("log", false, "enable requests logging")
	app.Flag("cert", "", "cert file for https")
	app.Flag("key", "", "key file for https")

	app.Run(run)

}

func run(args cli.Flags) error {
	startServer(args.String("host"), args.String("cert"), args.String("key"), args.String("dir"), args.Bool("log"), args.Bool("gzip"))
	return nil
}

func startServer(host, certFile, keyFile, dir string, enableLogs, enableCompression bool) {
	server := iris.New()
	server.Config().Render.Template.Engine = config.NoEngine

	if enableLogs {
		server.Use(logger.Default())
	}

	hasIndex := utils.Exists(dir + utils.PathSeparator + "index.html")

	// if enable compression then cache gzip files, if this dir doesn't contains an index.html then serve as fileserver
	serveHandler := iris.StaticHandlerFunc(dir, 0, enableCompression, !hasIndex)

	server.Get("/*filepath", func(ctx *iris.Context) {
		if len(ctx.Param("filepath")) < 2 && hasIndex {
			ctx.Request.SetRequestURI("index.html")
		}
		ctx.Next()

	}, serveHandler)

	cli.Printf("Server is running at: %s\n", host)
	if certFile != "" && keyFile != "" {
		server.ListenTLS(host, certFile, keyFile)
	} else {
		server.Listen(host)
	}

}
