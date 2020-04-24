package weint

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

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
	Host string
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
	text += "微博数量: " + strconv.FormatInt(info.StatuesCount, 10) + " | "
	return nil
}

func (o *ConsoleOut) WriteWeiboInfo(info *WeiboInfo) error {
	return nil
}

func (o *SQLiteOut) WriteUserInfo(info *UserInfo) error {
	panic("implement me")
}

func (o *SQLiteOut) WriteWeiboInfo(info *WeiboInfo) error {
	panic("implement me")
}

func (o *ElasticOut) WriteUserInfo(info *UserInfo) error {
	panic("implement me")
}

func (o *ElasticOut) WriteWeiboInfo(info *WeiboInfo) error {
	panic("implement me")
}

func (o *FileCSVOut) WriteUserInfo(info *UserInfo) error {
	f, _ := os.OpenFile(o.WeiboFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	defer w.Flush()

	return w.Write([]string{strconv.FormatInt(info.Id, 10), info.ScreenName, info.Description, info.Gender,
		strconv.FormatInt(info.FollowCount, 10), strconv.FormatInt(info.FollowersCount, 10), strconv.FormatInt(info.StatuesCount, 10),
		strconv.FormatBool(info.Verified), info.VerifiedReason})
}

func (o *FileCSVOut) WriteWeiboInfo(info *WeiboInfo) error {
	panic("implement me")
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
