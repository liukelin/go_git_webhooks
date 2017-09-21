/**
 * server cli
 */
package service

import (
	"fmt"
	"github.com/go-redis/redis"
	// "io/ioutil"
	"encoding/json"
	"os/exec"
	"reflect"
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
		for {
			// err := rConn.RPsh("RedisKey", "value").Err()
			d, err := RConn.LPop(RedisKey).Result()

			// 判断重连
			if err == redis.Nil {
				// fmt.Println(" does not exists:", RedisKey)

			} else if err != nil {
				fmt.Println("redis LPop err:", err, reflect.TypeOf(err), ".\n")
				RConn, Rerr = RedisClient(params["redishost"], params["redispass"], redisdb)

			} else {
				if len(d) > 0 {
					fmt.Println("redis LPop:", reflect.TypeOf(d), d, int(time.Second), ".\n")

					dMaps := loads_json(d)

					// 执行shell
					v, ok := dMaps["shell"]
					if ok {

						// 异步处理
						fmt.Println("Run shell:", v, "\n")
						err := run_shell(v)
						if err != nil {
							fmt.Println(err, ".\n")
						} else {
							fmt.Println("success.\n")
						}
					}
				}
			}

			// 阻塞 2s
			time.Sleep(time.Second * 2)

			// 非阻塞
			// time.After(time.Second + 10)
		}
	}
	fmt.Println("RedisClient connection error.\n")

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
