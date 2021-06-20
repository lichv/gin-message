package models

import (
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type Message struct {
	Id int `json:"id" form:"id" gorm:"id"`
	Provider string `json:"provider" form:"provider" gorm:"provider"`
	Target string `json:"target" form:"target" gorm:"target"`
	Template string `json:"template" form:"template" gorm:"template"`
	Params string `json:"params" form:"params" gorm:"params"`
	Content string `json:"content" form:"content" gorm:"content"`
	Result string `json:"result" form:"result" gorm:"result"`
	Ptime int `json:"ptime" form:"ptime" gorm:"ptime"`
	Dtime int `json:"dtime" form:"dtime" gorm:"dtime"`
	Flag bool `json:"flag" form:"flag" gorm:"flag"`
	State bool `json:"state" form:"state" gorm:"state"`
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

func GetAllMessageCode( query map[string]interface{},orderBy interface{},limit int) ( []int, []error) {
	var messages []Message
	var errs []error
	var result []int

	model := db.Table("demo_whitelist_user").Select("code")
	for key, value := range query {
		b,err := lichv.In (MessageFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	model =model.Order(orderBy).Find(&messages)
	errs = model.GetErrors()

	for _, v := range messages {
		result = append(result, v.Id)
	}

	return result, errs
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

func AddMessage( data map[string]interface{}) error {
	message := Message{
		Id:data["id"].(int),
		Provider:data["provider"].(string),
		Target:data["target"].(string),
		Template:data["template"].(string),
		Params:data["params"].(string),
		Content:data["content"].(string),
		Result:data["result"].(string),
		Ptime:data["ptime"].(int),
		Dtime:data["dtime"].(int),
		Flag:data["flag"].(bool),
		State:data["state"].(bool),

	}
	if err:= db.Create(&message).Error;err != nil{
		return err
	}
	return nil
}

func EditMessage( code string,data map[string]interface{}) error {
	if err:= db.Model(&Message{}).Where("code=?",code).Updates(data).Error;err != nil{
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
