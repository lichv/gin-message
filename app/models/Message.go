package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type Message struct {
	Id int `json:"id" form:"id" gorm:"id"`
	Type string `json:"type" form:"type" gorm:"type"`
	Provider string `json:"provider" form:"provider" gorm:"provider"`
	Target string `json:"target" form:"target" gorm:"target"`
	Template string `json:"template" form:"template" gorm:"template"`
	Params string `json:"params" form:"params" gorm:"params"`
	Content string `json:"content" form:"content" gorm:"content"`
	Result string `json:"result" form:"result" gorm:"result"`
	Ptime int `json:"ptime" form:"ptime" gorm:"ptime"`
	Dtime int `json:"dtime" form:"dtime" gorm:"dtime"`
	Flag int `json:"flag" form:"flag" gorm:"flag"`
	State int `json:"state" form:"state" gorm:"state"`
}

var MessageFields = []string{"id", "type", "provider", "target", "template", "params", "content", "result", "ptime", "dtime", "flag", "state"}

func FindMessageByCode( code string) ( message *Message, err error) {
	err = db.Model(&Message{}).Where("code = ? ",code).First(&message).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Message{},err
	}
	return
}

func GetMessageOne( query map[string]interface{},orderBy interface{}) ( *Message,error) {
	var message Message
	model := db.Model(&Message{})
	for key, value := range query {
		b,err := lichv.In (MessageFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&message).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &Message{},nil
	}
	return &message, nil
}

func ExistMessageByCode(code string) (b bool,err error) {
	var message Message
	err = db.Model(&Message{}).Select("code").Where("code = ? ",code).First(&message).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetMessagePages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*Message, []error) {
	var messages []*Message
	var errs []error
	model := db.Model(&Message{})
	for key, value := range query {
		b,err := lichv.In (MessageFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&messages)
	errs = model.GetErrors()

	return messages, errs
}

func GetMessages( query map[string]interface{},orderBy interface{},limit int) ( []*Message, []error) {
	var Messages []*Message
	var errs []error
	model := db.Model(&Message{})
	for key, value := range query {
		b,err := lichv.In (MessageFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&Messages).GetErrors()

	return Messages, errs
}

func GetMessageTotal(maps interface{}) (count int,err error) {
	err = db.Model(&Message{}).Where("is_active = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func AddMessage( data map[string]interface{}) (int,error) {
	message := Message{}
	_,ok := data["id"]
	if ok {
		message.Id = lichv.IntVal(data["id"])
	}
	_,ok = data["type"]
	if ok {
		message.Type = lichv.Strval(data["type"])
	}
	_,ok = data["provider"]
	if ok {
		message.Provider = lichv.Strval(data["provider"])
	}
	_,ok = data["target"]
	if ok {
		message.Target = lichv.Strval(data["target"])
	}
	_,ok = data["template"]
	if ok{
		message.Template = lichv.Strval(data["template"])
	}
	_,ok = data["params"]
	if ok {
		message.Params = lichv.Strval(data["params"])
	}
	_,ok = data["content"]
	if ok {
		message.Content = lichv.Strval(data["content"])
	}
	_,ok = data["result"]
	if ok {
		message.Result = lichv.Strval(data["result"])
	}
	_,ok = data["ptime"]
	if ok {
		message.Ptime = lichv.IntVal(data["ptime"])
	}
	_,ok = data["dtime"]
	if ok {
		message.Dtime = lichv.IntVal(data["dtime"])
	}
	_,ok = data["flag"]
	if ok {
		message.Flag = lichv.IntVal(data["flag"])
	}
	_,ok = data["state"]
	if ok {
		message.State = lichv.IntVal(data["state"])
	}

	if message.Target == "" {
		return 0,errors.New("目标为空")
	}
	if message.Template=="" && message.Params=="" && message.Content=="" {
		return 0,errors.New("内容为空")
	}

	if err:= db.Create(&message).Error;err != nil{
		return 0,err
	}
	return message.Id,nil
}

func ModifyMessage( id int,data map[string]interface{}) error {
	if err:= db.Model(&Message{}).Where("id=?",id).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteMessageByCode(code string) error {
	if err := db.Where("code=?",code).Delete(Message{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteMessages(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := lichv.In (MessageFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&Message{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllMessage() error {
	if err := db.Unscoped().Delete(&Message{}).Error; err != nil {
		return err
	}
	return nil
}
