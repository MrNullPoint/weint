package weint

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/olivere/elastic"
	"io/ioutil"
	"os"
	"strconv"
)

const ES_USER_INFO = "user_infos"
const ES_WEIBO_INFO = "weibo_infos"

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
	Host   string
	client *elastic.Client
}

type FileOut struct {
	UserFileName  string
	WeiboFileName string
}

type FileCSVOut struct {
	FileOut
}

type FileJsonOut struct {
	FileOut
}

func (o *ConsoleOut) WriteUserInfo(info *UserInfo) error {
	text := ""
	text += "id: " + strconv.FormatInt(info.Id, 10) + " | "
	text += "用户名: " + info.ScreenName + " | "
	text += "性别: " + info.Gender + " | "
	text += "简介: " + info.Description + " | "
	text += "关注者数量: " + strconv.FormatInt(info.FollowCount, 10) + " | "
	text += "粉丝数量: " + strconv.FormatInt(info.FollowersCount, 10) + " | "
	text += "微博数量: " + strconv.FormatInt(info.StatusesCount, 10) + " | "
	return nil
}

func (o *ConsoleOut) WriteWeiboInfo(info *WeiboInfo) error {
	return nil
}

func (o *SQLiteOut) WriteUserInfo(info *UserInfo) error {
	db, err := gorm.Open("sqlite3", o.DBName)
	if err != nil {
		return err
	}

	db.AutoMigrate(&UserInfo{})

	return db.Create(info).Error
}

func (o *SQLiteOut) WriteWeiboInfo(info *WeiboInfo) error {
	db, err := gorm.Open("sqlite3", o.DBName)
	if err != nil {
		return err
	}

	db.AutoMigrate(&WeiboInfo{})

	return db.Create(info.Build()).Error
}

func (o *ElasticOut) SetUpClient() error {
	url := elastic.DefaultURL

	if o.Host != "" {
		url = "http://" + o.Host
	}

	var err error

	o.client, err = elastic.NewSimpleClient(elastic.SetURL(url))

	return err
}

func (o *ElasticOut) WriteUserInfo(info *UserInfo) error {
	if err := o.SetUpClient(); err != nil {
		return err
	}
	_, err := o.client.Index().Index(ES_USER_INFO).BodyJson(info).Do(context.Background())
	return err
}

func (o *ElasticOut) WriteWeiboInfo(info *WeiboInfo) error {
	if err := o.SetUpClient(); err != nil {
		return err
	}
	_, err := o.client.Index().Index(ES_WEIBO_INFO).BodyJson(info).Do(context.Background())
	return err
}

func (o *FileCSVOut) WriteUserInfo(info *UserInfo) error {
	f, _ := os.OpenFile(o.UserFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	defer w.Flush()

	return w.Write([]string{strconv.FormatInt(info.Id, 10), info.ScreenName, info.Description, info.ProfileUrl, info.Gender,
		strconv.FormatInt(info.FollowCount, 10), strconv.FormatInt(info.FollowersCount, 10), strconv.FormatInt(info.StatusesCount, 10),
		strconv.FormatBool(info.Verified), info.VerifiedReason})
}

func (o *FileCSVOut) WriteWeiboInfo(info *WeiboInfo) error {
	if info == nil {
		return nil
	}

	f, _ := os.OpenFile(o.WeiboFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	pb, _ := json.Marshal(info.Pics)

	return w.Write([]string{info.Idstr, info.CreatedAt, info.Source, info.Text, string(pb),
		info.CommentsCount.String(), info.AttitudesCount.String(), info.RepostsCount.String(),
		strconv.FormatInt(info.User.Id, 10), info.User.ScreenName})
}

func (o *FileJsonOut) WriteUserInfo(info *UserInfo) error {
	if b, err := json.Marshal(info); err != nil {
		return err
	} else {
		return ioutil.WriteFile(o.UserFileName, b, os.ModePerm)
	}
}

func (o *FileJsonOut) WriteWeiboInfo(info *WeiboInfo) error {
	fd, _ := os.OpenFile(o.WeiboFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer fd.Close()

	if info != nil {
		if b, err := json.Marshal(info); err != nil {
			return err
		} else {
			_, err := fd.Write(b)
			fd.WriteString("\n")
			return err
		}
	}

	return nil
}
