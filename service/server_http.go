/**
 * http server
 */
package service

import (
	"fmt"
	// "html/template"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
	"github.com/go-redis/redis"
	// "reflect"
	"strconv"
	// "strings"
)

var RConn0 *redis.Client
var Rerr0 error
var Params map[string]string

/**
 * web server 入口
 * @return {[type]} [description]
 */
func Server_http(params map[string]string) {

	Params = params

	redisdb, err0 := strconv.Atoi(params["redisdb"])
	if err0 != nil {
		redisdb = 0
	}
	RConn0, Rerr0 = RedisClient(params["redishost"], params["redispass"], redisdb)

	if Rerr0 == nil {

		http.HandleFunc("/", server_http_action)

		// strconv.Itoa(port)
		portStr := ":" + params["port"]

		fmt.Println("your portStr is:", portStr, "\n")

		// mux := http.NewServeMux()
		// err := http.ListenAndServe(portStr, mux)
		err := http.ListenAndServe(portStr, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
		// go mux.Run()
	}
	fmt.Println("RedisClient connection error.\n")
}

/**
 * [web_server_action 请求业务处理]
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func server_http_action(w http.ResponseWriter, r *http.Request) {
	// 解析参数, 默认是不会解析的
	r.ParseForm()

	// d := r.Form["d"]
	d := r.FormValue("d")
	key := r.FormValue("sign")
	if len(Params["signKey"]) > 0 && Params["signKey"] != key {

		fmt.Fprintf(w, "signKey error.")

	} else {

		fmt.Println("body:", r.Form)

		err := RConn0.RPush(RedisKey, d).Err()

		// 判断重连
		if err != nil {
			redisdb, err0 := strconv.Atoi(Params["redisdb"])
			if err0 != nil {
				redisdb = 0
			}
			RConn0, Rerr0 = RedisClient(Params["redishost"], Params["redispass"], redisdb)
			if Rerr0 == nil {
				fmt.Println("RedisClient connection error2.\n")
				fmt.Fprintf(w, "RedisClient connection error2.")
			} else {
				fmt.Println("redis RPush error:", err)
				fmt.Fprintf(w, "redis RPush error.")
			}
		} else {
			fmt.Fprintf(w, "success.")
			// io.WriteString(w, "success.")
		}
	}

}
