package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	file, _ := os.OpenFile("logrus_study/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logrus.SetOutput(io.MultiWriter(file, os.Stdout)) // 同时输出屏幕和文件
	logrus.Infof("你好")
	logrus.Error("出错了")
	logrus.Errorf("出错了")
	logrus.Errorln("出错了")
}
