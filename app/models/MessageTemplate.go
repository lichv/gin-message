package models

import (
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type MessageTemplate struct {
	Id int `json:"id" form:"id" gorm:"id"`
	Code string `json:"code" form:"code" gorm:"code"`
	Name string `json:"name" form:"name" gorm:"name"`
	Type string `json:"type" form:"type" gorm:"type"`
	Owner string `json:"owner" form:"owner" gorm:"owner"`
	Scope string `json:"scope" form:"scope" gorm:"scope"`
	Content string `json:"content" form:"content" gorm:"content"`
	Flag bool `json:"flag" form:"flag" gorm:"flag"`
	State bool `json:"state" form:"state" gorm:"state"`
}

var MessageTemplateFields = []string{"id", "code", "name", "type", "owner", "scope", "content", "flag", "state"}

func FindMessageTemplateByCode( code string) ( messageTemplate *MessageTemplate, err error) {
	err = db.Model(&MessageTemplate{}).Where("code = ? ",code).First(&messageTemplate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &MessageTemplate{},err
	}
	return
}

func GetMessageTemplateOne( query map[string]interface{},orderBy interface{}) ( *MessageTemplate,error) {
	var messageTemplate MessageTemplate
	model := db.Model(&MessageTemplate{})
	for key, value := range query {
		b,err := lichv.In (MessageTemplateFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&messageTemplate).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &MessageTemplate{},nil
	}
	return &messageTemplate, nil
}

func ExistMessageTemplateByCode(code string) (b bool,err error) {
	var messageTemplate MessageTemplate
	err = db.Model(&MessageTemplate{}).Select("code").Where("code = ? ",code).First(&messageTemplate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetMessageTemplatePages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*MessageTemplate, []error) {
	var messageTemplates []*MessageTemplate
	var errs []error
	model := db.Model(&MessageTemplate{})
	for key, value := range query {
		b,err := lichv.In (MessageTemplateFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&messageTemplates)
	errs = model.GetErrors()

	return messageTemplates, errs
}

func GetAllMessageTemplateCode( query map[string]interface{},orderBy interface{},limit int) ( []string, []error) {
	var messageTemplates []MessageTemplate
	var errs []error
	var result []string

	model := db.Table("demo_whitelist_user").Select("code")
	for key, value := range query {
		b,err := lichv.In (MessageTemplateFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	model =model.Order(orderBy).Find(&messageTemplates)
	errs = model.GetErrors()

	for _, v := range messageTemplates {
		result = append(result, v.Code)
	}

	return result, errs
}

func GetMessageTemplates( query map[string]interface{},orderBy interface{},limit int) ( []*MessageTemplate, []error) {
	var MessageTemplates []*MessageTemplate
	var errs []error
	model := db.Model(&MessageTemplate{})
	for key, value := range query {
		b,err := lichv.In (MessageTemplateFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&MessageTemplates).GetErrors()

	return MessageTemplates, errs
}

func GetMessageTemplateTotal(maps interface{}) (count int,err error) {
	err = db.Model(&MessageTemplate{}).Where("is_active = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func AddMessageTemplate( data map[string]interface{}) error {
	messageTemplate := MessageTemplate{
		Id:data["id"].(int),
		Code:data["code"].(string),
		Name:data["name"].(string),
		Type:data["type"].(string),
		Owner:data["owner"].(string),
		Scope:data["scope"].(string),
		Content:data["content"].(string),
		Flag:data["flag"].(bool),
		State:data["state"].(bool),

	}
	if err:= db.Create(&messageTemplate).Error;err != nil{
		return err
	}
	return nil
}

func EditMessageTemplate( code string,data map[string]interface{}) error {
	if err:= db.Model(&MessageTemplate{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteMessageTemplateByCode(code string) error {
	if err := db.Where("code=?",code).Delete(MessageTemplate{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteMessageTemplates(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := lichv.In (MessageTemplateFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&MessageTemplate{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllMessageTemplate() error {
	if err := db.Unscoped().Delete(&MessageTemplate{}).Error; err != nil {
		return err
	}
	return nil
}
