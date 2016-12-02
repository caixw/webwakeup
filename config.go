// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/issue9/logs"
	"github.com/issue9/utils"
)

type config struct {
	Port     string  `json:"port"`
	HTTPS    bool    `json:"https"`
	CertFile string  `json:"certFile"`
	KeyFile  string  `json:"keyFile"`
	Tasks    []*task `json:"tasks"`
}

type task struct {
	Path     string   `json:"path"`     // 访问路径
	Password string   `json:"password"` // 密码
	Command  string   `json:"command"`  // 执行的命令
	Args     []string `json:"args"`     // 参数
	cmd      *exec.Cmd
}

// 加载配置文件
func loadConfig(path string) (*config, error) {
	conf := &config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, conf); err != nil {
		return nil, err
	}

	if len(conf.Port) == 0 {
		return nil, errors.New("必须指定端口号")
	}

	if conf.HTTPS {
		if !utils.FileExists(conf.KeyFile) {
			return nil, errors.New("keyFile 并不存在")
		}
		if !utils.FileExists(conf.CertFile) {
			return nil, errors.New("certFile 并不存在")
		}
	}

	if len(conf.Tasks) == 0 {
		return nil, errors.New("无任何任务执行")
	}

	for _, task := range conf.Tasks {
		if len(task.Path) == 0 {
			return nil, errors.New("path不能为空")
		}

		if len(task.Command) == 0 {
			return nil, errors.New("command不能为空")
		}

		task.cmd = exec.Command(task.Command, task.Args...)
	}

	return conf, nil
}

func (t *task) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("password") != t.Password {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err := t.cmd.Run(); err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
