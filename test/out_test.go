package test

import (
	"testing"
	"weint"
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
