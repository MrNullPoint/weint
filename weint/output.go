package weint

import "fmt"

type OutInterface interface {
	WriteUserInfo(info *UserInfo) error
	WriteWeiboInfo(info *WeiboInfo) error
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

func (o *ConsoleOut) WriteUserInfo(info *UserInfo) error {
	fmt.Println()
	return nil
}

func (o *ConsoleOut) WriteWeiboInfo(info *WeiboInfo) error {
	fmt.Println()
	return nil
}
