package Weint

const (
	TYPE_INFO  = 1
	TYPE_WEIBO = 1 << 1
)

type UserInfo struct {
}

type WeiboInfo struct {
}

type Spider struct {
	_type      int
	uid        string
	limit      int
	defaultOut OutInterface
	userOut    OutInterface
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
	return nil
}

func NewSpider() *Spider {
	s := new(Spider)
	s.defaultOut = &ConsoleOut{}
	s.userOut = nil
	return s
}
