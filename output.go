package Weint

import "fmt"

type OutInterface interface {
	WriteUserInfo(info *UserInfo)
	WriteWeiboInfo(info *WeiboInfo)
}

type ConsoleOut struct {
}

type SQLiteOut struct {
	DBName string
}

type ElasticOut struct {
	Host string
}

type FileOut struct {
	Filename string
}

type FileCSVOut struct {
	FileOut
}

type FileJsonOut struct {
	FileOut
}

func (o *ConsoleOut) WriteUserInfo(info *UserInfo) {
	fmt.Println()
}

func (o *ConsoleOut) WriteWeiboInfo(info *WeiboInfo) {
	fmt.Println()
}
