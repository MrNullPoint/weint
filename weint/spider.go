package weint

import (
	"encoding/json"
	"errors"
	browser "github.com/EDDYCJY/fake-useragent"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	TYPE_INFO  = 1
	TYPE_WEIBO = 1 << 1
)

const WEIBO_URL = "https://m.weibo.cn/api/container/getIndex"

type WeiboResp struct {
	Ok   int `json:"ok"`
	Data struct {
		UserInfo UserInfo `json:"userInfo"`
		Scheme   string   `json:"scheme"`
	} `json:"data"`
}

type UserInfo struct {
	Id              int64  `json:"id"`
	ScreenName      string `json:"screen_name"`
	ProfileImageUrl string `json:"profile_image_url"`
	ProfileUrl      string `json:"profile_url"`
	StatuesCount    int64  `json:"statues_count"`
	Verified        bool   `json:"verified"`
	VerifiedType    int    `json:"verified_type"`
	VerifiedTypeExt int    `json:"verified_type_ext"`
	VerifiedReason  string `json:"verified_reason"`
	Description     string `json:"description"`
	Gender          string `json:"gender"`
	FollowCount     int64  `json:"follow_count"`
	FollowersCount  int64  `json:"followers_count"`
	Mbtype          int    `json:"mbtype"`
	Mbrank          int64  `json:"mbrank"`
	Urank           int64  `json:"urank"`
}

type WeiboInfo struct {
}

type Spider struct {
	req        *http.Request
	client     *http.Client
	_type      int
	uid        string
	limit      int
	container  string
	profile    *UserInfo
	weibos     []*WeiboInfo
	defaultOut OutInterface
	userOut    OutInterface
}

// @function: 初始化一个爬虫对象
func NewSpider() *Spider {
	s := new(Spider)
	s.client = new(http.Client)
	s.req, _ = http.NewRequest("GET", WEIBO_URL, nil)
	s.req.Header.Set("user-agent", browser.Random())
	s.defaultOut = &ConsoleOut{}
	s.userOut = nil
	return s
}

// @function: 设置微博用户的 id
func (s *Spider) Uid(uid string) *Spider {
	s.uid = uid
	return s
}

// @function: 设置获取微博数量
func (s *Spider) Limit(limit int) *Spider {
	s.limit = limit
	return s
}

// @function: 设置爬虫类型
func (s *Spider) Type(_type int) *Spider {
	s._type += _type
	return s
}

// @function: 设置输出类型
func (s *Spider) Out(out OutInterface) *Spider {
	s.userOut = out
	return s
}

// @function: 运行爬虫
func (s *Spider) Run() error {
	if s.uid == "" {
		return errors.New("weibo user id is not correct")
	}

	if s._type&TYPE_INFO > 0 {
		if err := s.fetchUserInfo(); err != nil {
			return err
		}
	}

	if s._type&TYPE_WEIBO > 0 {
		if err := s.fetchWeiboinfo(); err != nil {
			return err
		}
	}

	return nil
}

// @function: 获取用户信息
func (s *Spider) fetchUserInfo() error {
	q := s.req.URL.Query()
	q.Add("type", "uid")
	q.Add("value", s.uid)
	s.req.URL.RawQuery = q.Encode()

	if err := s.doRequest(); err != nil {
		return err
	}

	if s.defaultOut != nil {
		if err := s.defaultOut.WriteUserInfo(s.profile); err != nil {
			return err
		}
	}

	if s.userOut != nil {
		if err := s.userOut.WriteUserInfo(s.profile); err != nil {
			return err
		}
	}

	return nil
}

// @function: 获取微博信息
func (s *Spider) fetchWeiboinfo() error {
	return nil
}

func (s *Spider) doRequest() error {
	resp, err := s.client.Do(s.req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	var data WeiboResp
	if err := json.Unmarshal(b, &data); err != nil {
		log.Fatal(err)
	}

	u, _ := url.Parse(data.Data.Scheme)
	m, _ := url.ParseQuery(u.RawQuery)

	s.profile = &data.Data.UserInfo
	s.container = m.Get("lfid")
}
