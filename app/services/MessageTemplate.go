package services

import (
	"gin-message/app/models"
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

func ExistMessageTemplateByCode(code string) (b bool,err error) {
	b,err = models.ExistMessageTemplateByCode(code)
	return b, err
}

func GetMessageTemplateTotal(maps interface{}) (count int,err error) {
	count,err = models.GetMessageTemplateTotal(map[string]interface{}{})
	return count, err
}
func GetMessageTemplateOne( query map[string]interface{},orderBy interface{}) (*MessageTemplate, error) {
	var nu *models.MessageTemplate
	nu,err := models.GetMessageTemplateOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferMessageTemplateModel(nu),nil
}

func GetMessageTemplatePages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (messageTemplates []*MessageTemplate, total int, errs []error) {
	total,err := models.GetMessageTemplateTotal(query)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetMessageTemplatePages(query,orderBy,pageNum,pageSize)
	messageTemplates = TransferMessageTemplates(us)
	return messageTemplates,total,nil
}
func GetAllMessageTemplateCode( query map[string]interface{},orderBy interface{},limit int)([]string,[]error){
	codes, errors := models.GetAllMessageTemplateCode(query, orderBy, limit)
	return codes,errors
}
func GetMessageTemplates( query map[string]interface{},orderBy interface{},limit int) ([]*MessageTemplate,[]error) {
	users, errors := models.GetMessageTemplates(query, orderBy, limit)
	messageTemplates := TransferMessageTemplates(users)
	return messageTemplates,errors
}

func AddMessageTemplate( data map[string]interface{}) (err error ){
	err = models.AddMessageTemplate(data)
	return err
}

func EditMessageTemplate( code string,data map[string]interface{}) (err error) {
	err = models.EditMessageTemplate(code,data)
	return err
}

func DeleteMessageTemplate(maps map[string]interface{}) (err error) {
	err = models.DeleteMessageTemplates(maps)
	return nil
}

func ClearAllMessageTemplate() (err error) {
	err = models.ClearAllMessageTemplate()
	return err
}

func TransferMessageTemplateModel(u *models.MessageTemplate)(messageTemplate *MessageTemplate){
	messageTemplate =  &MessageTemplate{
		Id:u.Id,
		Code:u.Code,
		Name:u.Name,
		Type: u.Type,
		Owner:u.Owner,
		Scope:u.Scope,
		Content:u.Type,
		Flag:u.Flag,
		State:u.State,
	}
	return
}
func TransferMessageTemplates(us []*models.MessageTemplate) (messageTemplates []*MessageTemplate) {
	for _,value := range us {
		messageTemplate := TransferMessageTemplateModel(value)
		messageTemplates = append(messageTemplates, messageTemplate)
	}
	return messageTemplates
}
