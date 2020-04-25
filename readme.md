# 微博爬虫非官方 API

golang 版本非官方新浪微博采集 API，不需要微博 api 也不需要登录之后的 cookie

## 安装方法

命令行版本直接从 release 处下载，golang 方式通过下载依赖

```shell
go get github.com/MrNullPoint/weint
```

## 使用方法

### 命令行版本

从 release 中下载对应操作系统的二进制文件，指定参数运行，下面以 linux 二进制可执行文件为例：

```shell script
NAME:
   A simple tool to get somebody's weibo data - A new cli application

USAGE:
   weint-v0.0.1-linux-amd64 [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --user value, -u value     set weibo user id, must set
   --info, -i                 set to get user's profile (default: false)
   --weibo, -w                set to get user's weibo list (default: false)
   --quick, -q                set to use quick mode, best practice is to use a proxy pool when set this flag (default: false)
   --proxy value, -p value    set proxy[暂不支持]
   --out value, -o value      set output type, csv/json/db/elastic
   --file value, -f value     set output filename
   --elastic value, -e value  set elastic search address (default: 127.0.0.1:9200)
   --help, -h                 show help (default: false)
   --version, -v              print the version (default: false)
```

### 代码实现版本

参考样例

## 样例

### 命令行版本

- 获取指定用户信息

```shell
$ weint -u 用户id -i
```

- 获取指定用户微博

```shell
$ weint -u 用户id -w
```

- 结果保存为 CSV

```shell
$ weint -u 用户id -w -o csv -f output.csv
```

- 结果保存为 JSON 文件

```shell
$ weint -u 用户id -w -o json -f output.json
```

- 结果保存到 SQLite

```shell
$ weint -u 用户id -w -o db -f output.db
```

- 结果保存到 Elasticsearch

```shell
$ weint -u 用户id -w -o elastic -e 127.0.0.1:9200
```

- 快速采集模式（不推荐）

```shell
$ weint -u 用户id -i -w -q
```

### 代码实现版本

- 获取指定用户信息

```go
spider := weint.NewSpider().Uid("微博账号id").Type(weint.TYPE_INFO)
if err := spider.Run(); err != nil {
    log.Panic(err)
}
```

- 获取指定用户微博

```go
spider := weint.NewSpider().Uid("微博账号id").Type(weint.TYPE_WEIBO)
if err := spider.Run(); err != nil {
    log.Panic(err)
}
```

- 结果保存为 CSV

```go
spider := weint.NewSpider()
spider.Uid("微博账号id")
spider.Type(weint.TYPE_INFO)
spider.Type(weint.TYPE_WEIBO)
spider.Out(&weint.FileCSVOut{FileOut: weint.FileOut{
	UserFileName:  "profile.csv",
	WeiboFileName: "weibo.csv",
}})
if err := spider.Run(); err != nil {
    log.Panic(err)
}
```

- 结果保存为 JSON 文件

```go
spider := weint.NewSpider()
spider.Uid("微博账号id")
spider.Type(weint.TYPE_INFO)
spider.Type(weint.TYPE_WEIBO)
spider.Out(&weint.FileJsonOut{FileOut: weint.FileOut{
	UserFileName:  "user.json",
	WeiboFileName: "weibo.json",
}})
if err := spider.Run(); err != nil {
    log.Panic(err)
}
```

- 结果保存到 SQLite

```go
spider := weint.NewSpider()
spider.Uid("微博账号id")
spider.Type(weint.TYPE_INFO)
spider.Type(weint.TYPE_WEIBO)
spider.Out(&weint.SQLiteOut{DBName: "db.db"})
if err := spider.Run(); err != nil {
    log.Panic(err)
}
```

- 结果保存到 Elasticsearch

```go
spider := weint.NewSpider()
spider.Uid("微博账号id")
spider.Type(weint.TYPE_INFO)
spider.Type(weint.TYPE_WEIBO)
spider.Out(&weint.ElasticOut{Host: "127.0.0.1:9200"})
if err := spider.Run(); err != nil {
    log.Panic(err)
}
```

## TODO

- [ ] 获取指定用户粉丝和关注者
- [ ] 支持限制返回
- [ ] 支持代理
- [ ] 增加 godoc

## Inspired By

- python 版本微博爬虫 - [weiboSpider](https://github.com/dataabc/weiboSpider)
- python 版本推文采集 uofficial api - [twint](https://github.com/twintproject/twint)
