## 目录结构

``` lua
|-- layout-go
    |-- bootstrap               #引导,用于注册各类服务
    |-- build                   #编译后的目录
    |-- cmd                     #命令
    |   |-- make                #代码生成命令
    |   |-- migrate             #数据迁移命令
    |   |-- seed                #数据填充命令
    |   |-- service             #业务层命令
    |-- config                  #配置
    |-- cron                    #定时
    |-- database                #数据相关
    |   |-- dao                 #gen生成
    |   |-- factory             #数据工厂
    |   |-- migration           #迁移文件
    |   |-- model               #模型
    |   |-- seeder              #填充文件
    |-- docs                    #文档
    |-- grpc                    #微服务
    |-- http                    #web
    |   |-- api1                #web端,如user,merchant,admin
    |       |-- controller      #控制器
    |       |-- request         #表单
    |       |-- route           #路由
    |-- inlet                   #入口
    |   |-- cmd                 #命令入口
    |   |-- cron                #定时入口
    |   |-- grpc                #微服务入口
    |   |-- http                #web入口
    |       |-- api1            #web的api1入口
    |-- pkg                     #自定义package
    |   |-- app                 #综合
    |   |-- console             #打印
    |   |-- database            #数据库连接
    |   |-- gin                 #web服务
    |   |-- jwt                 #Auth验证
    |   |-- logger              #日志
    |   |-- migrate             #数据迁移
    |   |-- queue               #队列
    |   |-- redis               #redis
    |   |-- seed                #数据填充
    |   |-- shutdown            #服务关闭
    |   |-- timer               #定时器
    |   |-- validator           #表单验证
    |-- service                 #业务抽象层
    |-- storage                 #文件目录
        |-- log                 #日志目录

```

## 路由规则

| 请求方法   | API 示例             | 说明    |
|--------|--------------------|-------|
| GET    | /api/user          | 获取列表  |
| GET    | /api/user/{id}     | 获取详情  |
| POST   | /api/user          | 新增    |
| PUT    | /api/user/{id}     | 修改    |
| DELETE | /api/user/{id}     | 删除    |


## 所有命令

启动web服务：

```
$ go run ./inlet/http/api1/main.go server
```

启动定时服务：

```
$ go run ./inlet/cron/main.go cron
```

生成web文档：

```
$ swag init --parseDependency --parseInternal -g inlet/http/api1/main.go
```

cmd命令：

```
$ go run ./inlet/cmd -h
Default will run "serve" command, you can use "-h" flag to see all subcommands

Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  make        Generate file and code
  migrate     Run database migration
  seed        Insert fake data to the database
  service     业务层cmd

Flags:
  -h, --help          help for this command
  -p, --path string   example: --path=./config/setting.yaml (default "./config/setting.yaml")

Use " [command] --help" for more information about a command.
```

make 命令：

```
$ go run ./inlet/cmd/main.go make -h
Generate file and code

Usage:
   make [command]

Available Commands:
  all         Crate all file, example: make all api1 user
  cmd         Create a command, should be snake_case, example: make cmd buckup_database
  controller  Create api controller, example: make controller api1 user
  factory     Create model's factory file, example: make factory user
  gen         Generate file and code, example: make gen
  migration   Create a migration file, example: make migration user
  model       Crate model file, example: make model user
  request     Create request file, example make request user
  route       Crate route file, example: make route api1 user
  seeder      Create seeder file, example: make seeder user

Flags:
  -h, --help   help for make

Global Flags:
  -p, --path string   example: --path=./config/setting.yaml (default "./config/setting.yaml")

Use " make [command] --help" for more information about a command.
```

migrate 命令：

```
$ go run ./inlet/cmd/main.go migrate -h
Run database migration

Usage:
   migrate [command]

Available Commands:
  down        Reverse the up command
  fresh       Drop all tables and re-run all migrations
  refresh     Reset and re-run all migrations
  reset       Rollback all database migrations
  up          Run unMigrated migrations

Flags:
  -h, --help   help for migrate

Global Flags:
  -p, --path string   example: --path=./config/setting.yaml (default "./config/setting.yaml")

Use " migrate [command] --help" for more information about a command.
```

seed 命令：

```
$ go run ./inlet/cmd/main.go seed -h
Insert fake data to the database

Usage:
   seed [flags]

Flags:
  -h, --help   help for seed

Global Flags:
  -p, --path string   example: --path=./config/setting.yaml (default "./config/setting.yaml")
```
