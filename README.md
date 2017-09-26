git web hooks

    1、搭建一个web server来接收web hooks的http请求

    2、将接收到的请求push到消息队列

    3、搭建一个自动运行脚本，异步处理接收的请求，在服务端执行命令行


    http request -> web server -> message queue  ->cli server 1
                                                 -> cli server 2
                                                 -> cli server ...

    4、这里的单机案例，如果多机需要执行的话，可用rabbitmq 队列的订阅模式，在各个机器运行 -server cli 服务

    
## 目录结构

~~~

├─service/              处理方法
│  ├─redis.go           redis操作类
│  ├─server_cli.go      脚本运行文件
│  └─server_http.go     http服务文件
└─main.go               入口文件

~~~


    下载安装包（此操作将包安装在gopath/src 下）
    Install: 
    $ go get -u github.com/go-redis/redis

    $ go build go_git_webhooks


    运行 http server： （为了防止恶意请求，加一个sign key验证 http请求合法性。）
    $ path/go_git_webhooks -server http -port 8888 -sign 123456

    运行 cli server：
    $ path/go_git_webhooks -server cli


    更多参数
    $ path/go_git_webhooks -help
      -debug
            Are you debug?
      -port string
            http服务端口，8888 (default "8888")
      -process string
            process member ? (default "1")
      -redishost string
            redis ip:端口 ? (default "localhost:6379")
      -server string
            服务类型 (http/cli) ? (default "http")
      -sign string
            http请求鉴权key, 设置后请求需要验证 &sign=key

    配置nginx:
![image](https://raw.githubusercontent.com/liukelin/go_git_webhooks/master/img/2.png)
    
    配置git web hooks:
![image](https://raw.githubusercontent.com/liukelin/go_git_webhooks/master/img/1.png)

    测试:
    curl -XGET http://webhooks.liukelin.top?sign=123456&d={"shell":"cd /var/www/obj/ && git pull"}

    d:{
        "shell":"" // 服务端所需要执行的命令
    }

    相当于:
    /bin/sh cd /var/www/obj/ && git pull

    也可以将参数d的值base64_encode后传递：
    curl -XGET http://webhooks.liukelin.top?sign=123456&d=eyJzaGVsbCI6ImNkIC92YXIvd3d3L29iai8gJiYgZ2l0IHB1bGwifQ==

    {"shell":"cd /var/www/object && expect /var/log/web_hooks/git_shell.exp"}
    {"shell":"cd /var/www/object2 && expect /var/log/web_hooks/git_shell.exp"}

总结:  
    
    延伸下去，就是个分布式任务调度系统。

    优化项：
        对于shell值，为了安全，可以做下命令过滤
        server cli 是消费队列数据，消费模式下单机里可以开启多个进程消费
        server cli 执行命令可使用异步操作，
        消费任务，可加入ACK机制（比如配合使用rabbitmq），允许消费失败归队
        end...

    代码优化：
        and...







