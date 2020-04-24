package test

import (
	"testing"
	"weint"
)

func TestCSVOut(t *testing.T) {
	spider := weint.NewSpider()
	spider.Uid("1784537661")
	spider.Type(weint.TYPE_INFO)
	spider.Out(&weint.FileCSVOut{FileOut: weint.FileOut{
		UserFileName:  "1784537661-user.csv",
		WeiboFileName: "1784537661-weibo.csv",
	}})
	if err := spider.Run(); err != nil {
		t.Error(err)
	}
}
