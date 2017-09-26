/**
 * liukelin
 *
 */
package main

import (
	"fmt"
	// "html/template"
	// "io/ioutil"
	"flag"
	// "log"
	// "os"
	// "github.com/go-redis/redis"
	service "go_git_webhooks/service"
	// "strings"
)

var debug *bool = flag.Bool("debug", false, "Are you debug?")
var server *string = flag.String("server", "http", "服务类型 (http/cli) ?")
var port *string = flag.String("port", "8888", "http服务端口，8888")
var process *string = flag.String("process", "1", "process member ?")
var redishost *string = flag.String("redishost", "localhost:6379", "redis ip:端口 ?")
var signKey *string = flag.String("sign", "", "http请求鉴权key, 设置后请求需要验证 &sign=key")

// var rConn *redis.Client

/**
 *  main
 * @return {[type]} [description]
 * -debug xx -server xx -port xx -process xx
 * -debug
 * -help
 */
func main() {
	flag.Parse()

	fmt.Println("your server is:", *server, "\n")

	params := make(map[string]string)
	params["server"] = *server
	params["port"] = *port
	params["process"] = *process
	params["redishost"] = *redishost
	params["redispass"] = ""
	params["redisdb"] = "0"
	params["signKey"] = *signKey

	switch *server {

	case "cli":

		service.Server_cli(params)

	case "http":

		service.Server_http(params)

	}

	if *debug == true {
		fmt.Println("your debug is:", debug, "\n")
		// fmt.Println("redis:", rClient, "\n")
	}
}

/**
 * [load_json_conf 加载获取配置]
 * @return {[type]} [description]
 */
func load_json_conf() {

}
