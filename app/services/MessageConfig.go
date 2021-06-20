package services

import (
	"encoding/json"
	"gin-message/app/models"
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

func ExistMessageConfigByCode(code string) (b bool,err error) {
	b,err = models.ExistMessageConfigByCode(code)
	return b, err
}

func GetMessageConfigTotal(maps interface{}) (count int,err error) {
	count,err = models.GetMessageConfigTotal(map[string]interface{}{})
	return count, err
}
func GetMessageConfigOne( query map[string]interface{},orderBy interface{}) (*MessageConfig, error) {
	var nu *models.MessageConfig
	nu,err := models.GetMessageConfigOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferMessageConfigModel(nu),nil
}

func FindMessageConfigValueByCode( code string) (map[string]interface{}, error) {
	var nu *models.MessageConfig
	var config map[string]interface{}
	nu,err := models.GetMessageConfigOne(map[string]interface{}{"code":code},"code asc")
	if err != nil {
		return map[string]interface{}{},err
	}
	var res map[string]interface{}
	if nu.Datatype== "json" {
		err := json.Unmarshal([]byte(nu.Data), &res)
		if err != nil{
			return nil,err
		}
		return res,nil
	}
	err = json.Unmarshal([]byte(nu.Data), &config)
	if err != nil {
		return map[string]interface{}{},err
	}
	return config,nil
}

func GetMessageConfigPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (messageConfigs []*MessageConfig, total int, errs []error) {
	total,err := models.GetMessageConfigTotal(query)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetMessageConfigPages(query,orderBy,pageNum,pageSize)
	messageConfigs = TransferMessageConfigs(us)
	return messageConfigs,total,nil
}
func GetAllMessageConfigCode( query map[string]interface{},orderBy interface{},limit int)([]string,[]error){
	codes, errors := models.GetAllMessageConfigCode(query, orderBy, limit)
	return codes,errors
}
func GetMessageConfigs( query map[string]interface{},orderBy interface{},limit int) ([]*MessageConfig,[]error) {
	users, errors := models.GetMessageConfigs(query, orderBy, limit)
	messageConfigs := TransferMessageConfigs(users)
	return messageConfigs,errors
}

func AddMessageConfig( data map[string]interface{}) (err error ){
	err = models.AddMessageConfig(data)
	return err
}

func EditMessageConfig( code string,data map[string]interface{}) (err error) {
	err = models.EditMessageConfig(code,data)
	return err
}

func DeleteMessageConfig(maps map[string]interface{}) (err error) {
	err = models.DeleteMessageConfigs(maps)
	return nil
}

func ClearAllMessageConfig() (err error) {
	err = models.ClearAllMessageConfig()
	return err
}

func TransferMessageConfigModel(u *models.MessageConfig)(messageConfig *MessageConfig){
	messageConfig =  &MessageConfig{
		Id:u.Id,
		Code:u.Code,
		Name:u.Name,
		Owner:u.Owner,
		Provider:u.Provider,
		Type:u.Type,
		Data:u.Data,
		Datatype:u.Datatype,
		Options:u.Options,
		Flag:u.Flag,
		State:u.State,
	}
	return
}
func TransferMessageConfigs(us []*models.MessageConfig) (messageConfigs []*MessageConfig) {
	for _,value := range us {
		messageConfig := TransferMessageConfigModel(value)
		messageConfigs = append(messageConfigs, messageConfig)
	}
	return messageConfigs
}
