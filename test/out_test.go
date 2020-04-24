package test

import (
	"github.com/MrNullPoint/weint"
	"testing"
)

func TestCSVOut(t *testing.T) {
	spider := weint.NewSpider()
	spider.Uid("5644764907")
	spider.Type(weint.TYPE_INFO)
	spider.Type(weint.TYPE_WEIBO)
	spider.Out(&weint.FileCSVOut{FileOut: weint.FileOut{
		UserFileName:  "5644764907-user.csv",
		WeiboFileName: "5644764907-weibo.csv",
	}})
	if err := spider.Run(); err != nil {
		t.Error(err)
	}
}

func TestSQLiteOut(t *testing.T) {
	spider := weint.NewSpider()
	spider.Uid("5644764907")
	spider.Type(weint.TYPE_INFO)
	spider.Type(weint.TYPE_WEIBO)
	spider.Out(&weint.SQLiteOut{DBName: "5644764907.db"})
	if err := spider.Run(); err != nil {
		t.Error(err)
	}
}

func TestJsonOut(t *testing.T) {
	spider := weint.NewSpider()
	spider.Uid("5644764907")
	spider.Type(weint.TYPE_INFO)
	spider.Type(weint.TYPE_WEIBO)
	spider.Out(&weint.FileJsonOut{FileOut: weint.FileOut{
		UserFileName:  "5644764907-user.json",
		WeiboFileName: "5644764907-weibo.json",
	}})
	if err := spider.Run(); err != nil {
		t.Error(err)
	}
}

func TestElasticOut(t *testing.T) {
	spider := weint.NewSpider()
	spider.Uid("5644764907")
	spider.Type(weint.TYPE_INFO)
	spider.Type(weint.TYPE_WEIBO)
	spider.Out(&weint.ElasticOut{Host: "127.0.0.1:9200"})
	if err := spider.Run(); err != nil {
		t.Error(err)
	}
}
