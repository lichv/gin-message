package models

import (
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type MessageConfig struct {
	Id int `json:"id" form:"id" gorm:"id"`
	Code string `json:"code" form:"code" gorm:"code"`
	Name string `json:"name" form:"name" gorm:"name"`
	Owner string `json:"owner" form:"owner" gorm:"owner"`
	Provider string `json:"provider" form:"provider" gorm:"provider"`
	Type string `json:"type" form:"type" gorm:"type"`
	Data string `json:"data" form:"data" gorm:"data"`
	Datatype string `json:"datatype" form:"datatype" gorm:"datatype"`
	Options string `json:"options" form:"options" gorm:"options"`
	Flag bool `json:"flag" form:"flag" gorm:"flag"`
	State bool `json:"state" form:"state" gorm:"state"`
}

var MessageConfigFields = []string{"id", "code", "name", "owner", "provider", "type", "data", "datatype", "options", "flag", "state"}

func FindMessageConfigByCode( code string) ( messageConfig *MessageConfig, err error) {
	err = db.Model(&MessageConfig{}).Where("code = ? ",code).First(&messageConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &MessageConfig{},err
	}
	return
}

func GetMessageConfigOne( query map[string]interface{},orderBy interface{}) ( *MessageConfig,error) {
	var messageConfig MessageConfig
	model := db.Model(&MessageConfig{})
	for key, value := range query {
		b,err := lichv.In (MessageConfigFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&messageConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &MessageConfig{},nil
	}
	return &messageConfig, nil
}

func ExistMessageConfigByCode(code string) (b bool,err error) {
	var messageConfig MessageConfig
	err = db.Model(&MessageConfig{}).Select("code").Where("code = ? ",code).First(&messageConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetMessageConfigPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*MessageConfig, []error) {
	var messageConfigs []*MessageConfig
	var errs []error
	model := db.Model(&MessageConfig{})
	for key, value := range query {
		b,err := lichv.In (MessageConfigFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&messageConfigs)
	errs = model.GetErrors()

	return messageConfigs, errs
}

func GetAllMessageConfigCode( query map[string]interface{},orderBy interface{},limit int) ( []string, []error) {
	var messageConfigs []MessageConfig
	var errs []error
	var result []string

	model := db.Table("demo_whitelist_user").Select("code")
	for key, value := range query {
		b,err := lichv.In (MessageConfigFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	model =model.Order(orderBy).Find(&messageConfigs)
	errs = model.GetErrors()

	for _, v := range messageConfigs {
		result = append(result, v.Code)
	}

	return result, errs
}

func GetMessageConfigs( query map[string]interface{},orderBy interface{},limit int) ( []*MessageConfig, []error) {
	var MessageConfigs []*MessageConfig
	var errs []error
	model := db.Model(&MessageConfig{})
	for key, value := range query {
		b,err := lichv.In (MessageConfigFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&MessageConfigs).GetErrors()

	return MessageConfigs, errs
}

func GetMessageConfigTotal(maps interface{}) (count int,err error) {
	err = db.Model(&MessageConfig{}).Where("is_active = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func AddMessageConfig( data map[string]interface{}) error {
	messageConfig := MessageConfig{
		Id:data["id"].(int),
		Code:data["code"].(string),
		Name:data["name"].(string),
		Owner:data["owner"].(string),
		Provider:data["provider"].(string),
		Type:data["type"].(string),
		Data:data["data"].(string),
		Datatype:data["datatype"].(string),
		Options:data["options"].(string),
		Flag:data["flag"].(bool),
		State:data["state"].(bool),

	}
	if err:= db.Create(&messageConfig).Error;err != nil{
		return err
	}
	return nil
}

func EditMessageConfig( code string,data map[string]interface{}) error {
	if err:= db.Model(&MessageConfig{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteMessageConfigByCode(code string) error {
	if err := db.Where("code=?",code).Delete(MessageConfig{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteMessageConfigs(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := lichv.In (MessageConfigFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&MessageConfig{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllMessageConfig() error {
	if err := db.Unscoped().Delete(&MessageConfig{}).Error; err != nil {
		return err
	}
	return nil
}
