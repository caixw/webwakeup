webwakeup [![Build Status](https://travis-ci.org/caixw/webwakeup.svg?branch=master)](https://travis-ci.org/caixw/webwakeup)
======

一个小玩意儿，可以处理类似 webhook 等远程唤醒相应程序的功能。
可以通过自定义的配置文件，指定路径的执行的命令。



### 安装

1. 安装程序: `go get github.com/caixw/webwakeup`；
1. 编写配置文件，可直接复代码下的 config.json，在此基础上进行修改；

##### 关于日志输出

和 config.json 同目录下存在 logs.xml 文件，则会将此文件作为日志配置文件进行加载。
否则自动将 ERROR, CRITICAL, WARN 三个通道的内容输出到 stderr 中，其它通道的忽略。
日志采用 [logs](https://github.com/issue9/logs)，相关的配置文件可参考其内容。


### 版权

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
