package test

import (
	"github.com/MrNullPoint/weint"
	"testing"
)

func TestSpiderOut(t *testing.T) {
	spider := weint.NewSpider()
	spider.Uid("1784537661")
	spider.Type(weint.TYPE_INFO)
	spider.Type(weint.TYPE_WEIBO)
	spider.Out(&weint.FileJsonOut{FileOut: weint.FileOut{
		UserFileName:  "1784537661-user.json",
		WeiboFileName: "1784537661-weibo.json",
	}})
	if err := spider.Run(); err != nil {
		t.Error(err)
	}
}
