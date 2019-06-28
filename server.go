package main

import (
	"github.com/gorilla/mux"
	"github.com/iswangwenbin/golion/api"
	"github.com/iswangwenbin/golion/pkg/sonyflakex"
	"github.com/micro/go-config"

	cf "github.com/micro/go-config/source/file"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const (
	version string = "0.01"
)

func init() {
	// 日志设置
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(log.DebugLevel)

	// 根据运行环境加载不同的配置文件
	filePath := "config/default.json"
	if _, err := os.Stat("/www/web/development"); !os.IsNotExist(err) {
		filePath = "config/development.json"
	}
	config.Load(cf.NewSource(cf.WithPath(filePath)))

	// MySQL 初始化
	//databasex.NewDatabase()

	// Sonyflake 自增初始化
	sonyflakex.NewSonyflake()
}

func main() {
	// HTTP
	router := mux.NewRouter()
	router.HandleFunc("/", api.Hello).Methods("GET")
	router.HandleFunc("/user", api.User).Methods("GET")
	// 监听端口
	port := config.Get("ENV", "port").String("8000")
	// 输出服务运行版本
	log.Infof("server version:%s", version)
	// 输出服务监听端口
	log.Infof("listen port:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
