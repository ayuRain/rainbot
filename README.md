
## Build & Run 

```shell
sh build.sh 
./output/bin/rainbot
```
## Code Structure
```
.
├── biz //核心业务逻辑
│   ├── client //调用第三方服务封装
│   ├── config //配置文件逻辑
│   │   ├── global.go
│   │   └── init.go
│   ├── handler //接口层逻辑
│   │   ├── chat.go
│   │   ├── init.go
│   │   ├── larkbot.go
│   │   └── ping.go
│   ├── model 
│   │   ├── chat.go
│   │   └── larkbot.go
│   ├── router
│   │   └── register.go
│   ├── service
│   │   ├── chat
│   │   │   ├── chat.go
│   │   │   └── util.go
│   │   ├── init.go
│   │   └── larkbot
│   └── util
│       └── http.go
├── build.sh
├── conf
│   ├── config.example
│   └── global.yaml
├── go.mod
├── go.sum
├── main.go
├── output
│   ├── bin
│   │   └── rainbot
│   └── bootstrap.sh
├── router.go
├── router_gen.go
└── script
└── bootstrap.sh
```