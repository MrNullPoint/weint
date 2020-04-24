# 微博爬虫非官方 API

golang 版本非官方新浪微博采集 API，不需要微博 api 也不需要登录之后的 cookie

## 安装方法

命令行版本直接从 release 处下载，golang 方式通过下载依赖

```shell
go get github.com/MrNullPoint/weint
```

## 使用方法

### 命令行版本

从 release 中下载对应操作系统的二进制文件，指定参数运行

```shell

```

### 代码实现版本

参考 [godoc](http://demo.com) 或者样例

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
$ weint -u 用户id -w -o elastic -h 127.0.0.1:9200
```

### 代码实现版本

- 获取指定用户信息
- 获取指定用户微博
- 指定代理
- 限制返回微博数量
- 结果保存为 CSV
- 结果保存为 JSON 文件
- 结果保存到 SQLite
- 结果保存到 Elasticsearch

## TODO

- [ ] 获取指定用户粉丝和关注者
- [ ] 支持多种代理模式

## Inspired By

- python 版本微博爬虫 - [weiboSpider](https://github.com/dataabc/weiboSpider)
- python 版本推文采集 uofficial api - [twint](https://github.com/twintproject/twint)
