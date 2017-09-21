    git web hooks 脚本md

    1、搭建一个web server来接收web hooks的http请求

    2、将接收到的请求push到消息队列

    3、搭建一个自动运行脚本，异步处理接收的请求


    http request -> web server -> message queue  ->cli server 1
                                                 -> cli server 2
                                                 -> cli server ...


    下载安装包（此操作将包安装在gopath/src 下）
    Install: 
    $ go get -u github.com/go-redis/redis


    运行 http server：
    $ path/go_git_webhooks -server http -port 8888 

    运行 cli server：
    $ path/go_git_webhooks -server cli


    为了防止恶意请求，可以加一个sign key验证 http请求合法性。
    $ path/go_git_webhooks -server http -port 8888 -sign 123456

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









