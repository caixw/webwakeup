// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/issue9/logs"
	"github.com/issue9/utils"
)

const (
	version = "0.1.0"

	logFile  = "logs.xml"
	confFile = "config.json"
)

func main() {
	confdir := flag.String("conf", "./", "指定配置文件目录")
	v := flag.Bool("v", false, "版本号")
	flag.Parse()

	if *v {
		fmt.Println("version:", version)
		return
	}

	if err := initLogs(filepath.Join(*confdir, logFile)); err != nil {
		panic(err)
	}

	conf, err := loadConfig(filepath.Join(*confdir, confFile))
	if err != nil {
		panic(err)
	}

	for _, task := range conf.Tasks {
		http.HandleFunc(task.Path, task.ServeHTTP)
	}

	if conf.HTTPS {
		http.ListenAndServeTLS(conf.Port, conf.CertFile, conf.KeyFile, nil)
	} else {
		http.ListenAndServe(conf.Port, nil)
	}
}

func initLogs(path string) error {
	if utils.FileExists(path) {
		return logs.InitFromXMLFile(path)
	}

	logs.SetWriter(logs.LevelCritical, os.Stderr, "", log.LstdFlags)
	logs.SetWriter(logs.LevelError, os.Stderr, "", log.LstdFlags)
	logs.SetWriter(logs.LevelWarn, os.Stderr, "", log.LstdFlags)
	return nil
}
