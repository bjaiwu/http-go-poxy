package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/bjaiwu/http-go-poxy/httpsserve"
)

var logger = log.New(os.Stderr, "httpsproxy:", log.Llongfile|log.LstdFlags)
var listenAdress string

func init() {
	// 	var configPath string
	// 	var debug bool // debug 模式
	// 	flag.StringVar(&configPath, "config", "config", "需要指点启动配置文件")
	flag.StringVar(&listenAdress, "L", "0.0.0.0:5589", "listen address.eg: 0.0.0.0:5589")
	// 	flag.BoolVar(&debug, "debug", false, "debug配置")
	flag.Parse()
	// 	// 绝对路径
	// 	config.Debug = debug
	if !checkAdress(listenAdress) {
		logger.Fatal("-L listen address format incorrect.Please check it")
	}

	// 	configPath, err := filepath.Abs(configPath)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		os.Exit(0)
	// 	}
	// 	osPath := "/" // 系统分割符
	// 	if runtime.GOOS == "windows" {
	// 		osPath = "\\"
	// 	}
	// 	config.OsPath = osPath
	// 	configPath += osPath // 配置文件路径
	// 	config.ConfigPath = configPath
	// 	byteJson, err := os.ReadFile(configPath + "config.yaml")
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		os.Exit(0)
	// 	}
	// 	err = yaml.Unmarshal(byteJson, &config.ConfigInfo)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		os.Exit(0)
	// 	}
	// 	if config.ConfigInfo.HomeId == "" {
	// 		log.Println("请配置homeId")
	// 		os.Exit(0)
	// 	}
	// 	// 获取所有rcu
	// 	if err = internal.InitSnCodeMap(); err != nil {
	// 		log.Println("获取组织所有rcu失败")
	// 		os.Exit(0)
	// 	}
}

func checkAdress(adress string) bool {
	_, err := net.ResolveTCPAddr("tcp", adress)
	return err == nil

}

func main() {
	// 运行
	go httpsserve.Serve(listenAdress)
	select {}
}
