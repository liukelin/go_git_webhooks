    git web hooks 脚本md

    1、搭建一个web server来接收web hooks的http请求

    2、搭建一个自动运行脚本，异步处理接收的请求


    http request -> web server -> message queue  ->cli server 1
                                                 -> cli server 2
                                                 -> cli server ...



    运行 http server：
    $ path/go_git_webhooks -server http -port 8888 

    运行 cli server：
    $ path/go_git_webhooks -server cli