/**
 * server cli
 */
package service

import (
	"fmt"
	"github.com/go-redis/redis"
	// "io/ioutil"
	"encoding/base64"
	"encoding/json"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

var RConn *redis.Client
var Rerr error

// 睡眠等等时间
// var sleep_ int = 2

/**
 * 后台脚本server 入口
 * @return {[type]} [description]
 */
func Server_cli(params map[string]string) {

	// string to int
	redisdb, err0 := strconv.Atoi(params["redisdb"])
	if err0 != nil {
		redisdb = 0
	}
	RConn, Rerr = RedisClient(params["redishost"], params["redispass"], redisdb)

	if Rerr == nil {

		// 按CPU核数 设置并行数量
		runtime.GOMAXPROCS(runtime.NumCPU())

		for {
			// err := rConn.RPsh("RedisKey", "value").Err()
			d, err := RConn.LPop(RedisKey).Result()

			// fmt.Println(time.Now(), " Lpop.")

			// 判断重连
			if err != nil {
				if err == redis.Nil {
					// fmt.Println(time.Now(), " does not exists:", err)
				} else {
					fmt.Println(time.Now(), "Redis connection error:", err, reflect.TypeOf(err), ".\n")
					RConn, Rerr = RedisClient(params["redishost"], params["redispass"], redisdb)
				}

				//将CPU时间片让给其它goroutine
				// runtime.Gosched()

			} else {
				if len(d) > 0 {
					fmt.Println(time.Now(), " - redis LPop:", reflect.TypeOf(d), d, int(time.Second), ".\n")
					// defer RConn.Close()
					go func() {
						ack := consu_data(d)
						fmt.Println(time.Now(), " results:", ack)
					}()
				}
			}

			// 阻塞 2s
			// time.Sleep(time.Second * 2)
			// time.Sleep(time.Second)
			// time.Sleep(1000 * time.Millisecond)
			time.Sleep(1e9) // sleep one second

			// 非阻塞
			// time.After(time.Second + 10)
			continue
		}

	} else {
		fmt.Println("Redis connection error.\n")
	}

}

/**
 * [consu_data 消费数据]
 * @param  {[type]} d string)       (ack bool [description]
 * @return {[type]}      [description]
 */
func consu_data(d string) (ack bool) {

	shell := ""
	// base64_decode
	decodeBytes, errBase := base64.StdEncoding.DecodeString(d)
	if errBase != nil {
		shell = d
	} else {
		shell = string(decodeBytes)
	}

	dMaps := loads_json(shell)
	v, ok := dMaps["shell"]
	if ok {
		fmt.Println("Run shell:", v, "\n")
		// 执行shell
		go func() {
			err := run_shell(v)
			if err != nil {
				fmt.Println(err, ".\n")
			} else {
				fmt.Println("success.\n")
			}
		}()
		// 异步处理(test)

		return true
	}

	return false
}

/**
 * 解析json
 */
// type d_json struct {
// 	Shell string
// 	Time  string
// }

func loads_json(jsonStr string) (maps map[string]string) {
	// var s d_json
	// jsonStr := `{"shell":"ls -la","time":"2017-09-21"}`
	// json.Unmarshal([]byte(jsonStr), &s)

	maps_ := make(map[string]string)

	var s interface{}
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		return maps_
	}

	m := s.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			// fmt.Println(k, "is string", vv)
			maps_[k] = vv

		case int:
			// fmt.Println(k, "is int", vv)
			maps_[k] = strconv.Itoa(vv)

		case float64:
			// fmt.Println(k, "is float64", vv)
			maps_[k] = strconv.FormatFloat(vv, 'E', -1, 32)

		case []interface{}:
			// 数组字典
			// fmt.Println(k, "is an array:")
			// for i, u := range vv {
			// 	fmt.Println(i, u)
			// }
		case map[string]interface{}:
			// 字典字典

		default:
			// fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	// fmt.Println(maps_)
	return maps_
}

/**
 f = map[string]interface{}{
    "Name": "Wednesday",
    "Age":  6,
    "Parents": []interface{}{
        "Gomez",
        "Morticia",
    },
    "Parents": {"sss"}interface{}{
        "Gomez",
        "Morticia",
    },
}
*/

/**
 * 执行shell命令
 */
func run_shell(shell string) (msg error) {
	cmd := exec.Command("/bin/sh", "-c", shell)
	_, err := cmd.Output()
	if err != nil {
		// panic(err.Error())
	}
	if err := cmd.Start(); err != nil {
		// panic(err.Error())
	}
	if err := cmd.Wait(); err != nil {
		// panic(err.Error())
	}
	return err
}
