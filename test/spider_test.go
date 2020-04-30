package test

import (
	"github.com/MrNullPoint/weint"
	"net/http"
	"net/url"
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

func TestSpiderClient(t *testing.T) {
	spider := weint.NewSpider()

	uri := url.URL{}
	proxy, _ := uri.Parse("http://127.0.0.1:1080")

	spider.Client(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	})

	spider.Type(weint.TYPE_INFO)
	spider.Type(weint.TYPE_WEIBO)

	if err := spider.Run(); err != nil {
		t.Error(err)
	}
}
