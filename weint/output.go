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

func (o *SQLiteOut) WriteUserInfo(info *UserInfo) error {
	panic("implement me")
}

func (o *SQLiteOut) WriteWeiboInfo(info *WeiboInfo) error {
	panic("implement me")
}

func (o *ElasticOut) WriteUserInfo(info *UserInfo) error {
	panic("implement me")
}

func (o *ElasticOut) WriteWeiboInfo(info *WeiboInfo) error {
	panic("implement me")
}

func (o *FileCSVOut) WriteUserInfo(info *UserInfo) error {
	panic("implement me")
}

func (o *FileCSVOut) WriteWeiboInfo(info *WeiboInfo) error {
	panic("implement me")
}

func (o *FileJsonOut) WriteUserInfo(info *UserInfo) error {
	panic("implement me")
}

func (o *FileJsonOut) WriteWeiboInfo(info *WeiboInfo) error {
	panic("implement me")
}
