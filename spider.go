package weint

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	TYPE_INFO  = 1
	TYPE_WEIBO = 1 << 1
)

type Type int64

const (
	Int64 Type = iota
	String
)

const WEIBO_URL = "https://m.weibo.cn/api/container/getIndex"

type WeiboResp struct {
	Ok   int `json:"ok"`
	Data struct {
		UserInfo *UserInfo    `json:"userInfo"`
		Cards    []*WeiboInfo `json:"cards"`
		TabsInfo struct {
			Tabs []struct {
				TabKey      string `json:"tabKey"`
				ContainerId string `json:"containerid"`
			} `json:"tabs"`
		} `json:"tabsInfo"`
		CardListInfo struct {
			SinceId uint64 `json:"since_id"`
		} `json:"cardlistInfo"`
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
	CardType int    `json:"card_type"`
	Itemid   string `json:"itemid"`
	Scheme   string `json:"scheme"`
	Mblog    struct {
		User           UserInfo   `json:"user"`
		Idstr          string     `json:"idstr"`
		Mid            string     `json:"mid"`
		Text           string     `json:"text"`
		Source         string     `json:"source"`
		OriginalPic    string     `json:"original_pic"`
		MblogVipType   int        `json:"mblog_vip_type"`
		CreatedAt      string     `json:"created_at"`
		RepostsCount   WeiboCount `json:"reposts_count"`
		CommentsCount  WeiboCount `json:"comments_count"`
		AttitudesCount WeiboCount `json:"attitudes_count"`
		Pics           []struct {
			Pid string `json:"pid"`
			Url string `json:"url"`
		} `json:"pics"`
	} `json:"mblog"`
}

type WeiboCount struct {
	Type   Type
	IntVal int
	StrVal string
}

// @implement: json unmarshal
func (count *WeiboCount) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		count.Type = String
		return json.Unmarshal(value, &count.StrVal)
	}
	count.Type = Int64
	return json.Unmarshal(value, &count.IntVal)
}

// 实现 json.Marshaller 接口
func (count *WeiboCount) MarshalJSON() ([]byte, error) {
	switch count.Type {
	case Int64:
		return json.Marshal(count.IntVal)
	case String:
		return json.Marshal(count.StrVal)
	default:
		return []byte{}, fmt.Errorf("impossible Weibo.Type")
	}
}

type Spider struct {
	req        *http.Request
	client     *http.Client
	_type      int
	quick      bool
	uid        string
	limit      int
	container  string
	since      uint64
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

// @function: 设置爬虫速度
func (s *Spider) Quick(quick bool) *Spider {
	s.quick = quick
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
		if err := s.fetchWeiboInfo(); err != nil {
			return err
		}
	}

	return nil
}

// @function: 获取用户信息
func (s *Spider) fetchUserInfo() error {
	q := s.req.URL.Query()
	q.Set("is_all", "1")
	q.Set("type", "uid")
	q.Set("value", s.uid)

	s.req.URL.RawQuery = q.Encode()

	if err := s.doRequest(); err != nil {
		return err
	}

	return s.outProfile()
}

// @function: 获取微博信息
func (s *Spider) fetchWeiboInfo() error {
	q := s.req.URL.Query()
	q.Set("is_all", "1")
	q.Set("type", "uid")
	q.Set("value", s.uid)
	q.Set("containerid", s.container)

	if s.since != 0 {
		q.Set("since_id", strconv.FormatUint(s.since, 10))
	}

	s.req.URL.RawQuery = q.Encode()

	if err := s.doRequest(); err != nil {
		return err
	}

	if err := s.outWeibo(); err != nil {
		return err
	}

	if s.since != 0 {
		return s.fetchWeiboInfo()
	}

	return nil
}

// @function: 执行请求获取微博响应数据
func (s *Spider) doRequest() error {
	rs := rand.NewSource(time.Now().Unix())
	ra := rand.New(rs)

	if !s.quick {
		time.Sleep(time.Duration(ra.Intn(5)+1) * time.Second)
	}

	s.req.Header.Set("User-Agent", uaList[ra.Intn(len(uaList))])

	resp, err := s.client.Do(s.req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK || string(b) == "" {
		return errors.New("weibo has no resp due to some reason such as request rate limit")
	}

	var data WeiboResp
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	s.profile = data.Data.UserInfo
	s.weibos = data.Data.Cards
	s.since = data.Data.CardListInfo.SinceId

	for _, t := range data.Data.TabsInfo.Tabs {
		if t.TabKey == "weibo" {
			s.container = t.ContainerId
		}
	}

	return nil
}

func (s *Spider) outProfile() error {
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

func (s *Spider) outWeibo() error {
	for _, w := range s.weibos {
		if s.defaultOut != nil {
			if err := s.defaultOut.WriteWeiboInfo(w); err != nil {
				return err
			}
		}

		if s.userOut != nil {
			if err := s.userOut.WriteWeiboInfo(w); err != nil {
				return err
			}
		}
	}

	return nil
}
