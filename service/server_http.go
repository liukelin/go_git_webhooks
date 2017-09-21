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
	"reflect"
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

		err := http.ListenAndServe(portStr, nil)

		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
	fmt.Println("RedisClient connection error.\n")
}

/**
 * [web_server_action description]
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func server_http_action(w http.ResponseWriter, r *http.Request) {
	// 解析参数, 默认是不会解析的
	r.ParseForm()

	/**
	fmt.Println("request_map:", r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("url_long:", r.Form["url_long"])
	**/
	// body := make(map[string]string)
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ";"))
	// }

	// fmt.Println(r.Form["data"], r.Form["branch"], r.Form["c"])
	// msg := ""

	// d := r.Form["d"]
	d := r.FormValue("d")

	fmt.Println("body d:", d, reflect.TypeOf(d))

	err := RConn0.RPush(Key, d).Err()

	// 判断重连
	if err != nil {
		redisdb, err0 := strconv.Atoi(Params["redisdb"])
		if err0 != nil {
			redisdb = 0
		}
		RConn0, Rerr0 = RedisClient(Params["redishost"], Params["redispass"], redisdb)
		if Rerr0 == nil {
			fmt.Println("RedisClient connection error2.\n")
		}
	}

	// 这个写入到w输出到客户端
	fmt.Fprintf(w, "return:", err)
}
