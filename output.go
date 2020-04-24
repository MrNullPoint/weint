package weint

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/olivere/elastic"
	"io/ioutil"
	"os"
	"strings"
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
	fmt.Println(strings.Join(info.Slice(), ","))
	return nil
}

func (o *ConsoleOut) WriteWeiboInfo(info *WeiboInfo) error {
	fmt.Println(strings.Join(info.Slice(), ","))
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

	return w.Write(info.Slice())
}

func (o *FileCSVOut) WriteWeiboInfo(info *WeiboInfo) error {
	if info == nil {
		return nil
	}

	f, _ := os.OpenFile(o.WeiboFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	return w.Write(info.Slice())
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
