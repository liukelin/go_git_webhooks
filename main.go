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
var server *string = flag.String("server", "http", "server type (http/cli) ?")
var port *string = flag.String("port", "8888", "port ?")
var process *string = flag.String("process", "1", "process member ?")
var redishost *string = flag.String("redishost", "localhost:6379", "redis host ?")

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

	/**
	if *server == "cli_server" {
		cli_server(*process)
	} else if *server == "web_server" {
		web_server(*port)
	} else {
		web_server(*port)
	}**/

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
