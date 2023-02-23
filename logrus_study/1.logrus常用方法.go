package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	cBlack = 0
	cRed   = 1
	cGreen = 2
)

func PrintColor(colorCode int, text string) {
	fmt.Printf("\033[3%dm%s\033[0m", colorCode, text)
}

func main() {

	PrintColor(cRed, "红色")
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetLevel(logrus.WarnLevel) // 设置日志等级
	logrus.SetLevel(logrus.InfoLevel) // 设置日志等级
	logrus.Errorf("出错了")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("debug")
	logrus.Println("打印")

	fmt.Println(logrus.GetLevel())               // 显示等级
	logrus.SetFormatter(&logrus.JSONFormatter{}) // 用JSON显示
	logrus.SetFormatter(&logrus.TextFormatter{   // 设置颜色
		ForceColors: true,
	})

	log := logrus.WithField("app", "study").WithField("service", "logrus") // 可以链式调用
	log = logrus.WithFields(logrus.Fields{
		"user": "ww",
		"ip":   "192.168.200.254",
	})
	log.Errorf("你好") // 会统一加上上面的字段

	fmt.Println("\033[31m 红色 \033[0m") // 设置字体颜色
	fmt.Println("\033[41m 红色 \033[0m") // 设置背景颜色
}
