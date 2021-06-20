package models

import (
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type Syslog struct {
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

var SyslogFields = []string{"id", "code", "name", "owner", "provider", "type", "data", "datatype", "options", "flag", "state"}

func FindSyslogByCode( code string) ( syslog *Syslog, err error) {
	err = db.Model(&Syslog{}).Where("code = ? ",code).First(&syslog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Syslog{},err
	}
	return
}

func GetSyslogOne( query map[string]interface{},orderBy interface{}) ( *Syslog,error) {
	var syslog Syslog
	model := db.Model(&Syslog{})
	for key, value := range query {
		b,err := lichv.In (SyslogFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&syslog).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &Syslog{},nil
	}
	return &syslog, nil
}

func ExistSyslogByCode(code string) (b bool,err error) {
	var syslog Syslog
	err = db.Model(&Syslog{}).Select("code").Where("code = ? ",code).First(&syslog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetSyslogPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*Syslog, []error) {
	var syslogs []*Syslog
	var errs []error
	model := db.Model(&Syslog{})
	for key, value := range query {
		b,err := lichv.In (SyslogFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&syslogs)
	errs = model.GetErrors()

	return syslogs, errs
}

func GetAllSyslogCode( query map[string]interface{},orderBy interface{},limit int) ( []string, []error) {
	var syslogs []Syslog
	var errs []error
	var result []string

	model := db.Table("demo_whitelist_user").Select("code")
	for key, value := range query {
		b,err := lichv.In (SyslogFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	model =model.Order(orderBy).Find(&syslogs)
	errs = model.GetErrors()

	for _, v := range syslogs {
		result = append(result, v.Code)
	}

	return result, errs
}

func GetSyslogs( query map[string]interface{},orderBy interface{},limit int) ( []*Syslog, []error) {
	var Syslogs []*Syslog
	var errs []error
	model := db.Model(&Syslog{})
	for key, value := range query {
		b,err := lichv.In (SyslogFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&Syslogs).GetErrors()

	return Syslogs, errs
}

func GetSyslogTotal(maps interface{}) (count int,err error) {
	err = db.Model(&Syslog{}).Where("is_active = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func AddSyslog( data map[string]interface{}) (int,error ){
	syslog := Syslog{
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
	if err:= db.Create(&syslog).Error;err != nil{
		return 0,err
	}
	return syslog.Id,nil
}

func EditSyslog( code string,data map[string]interface{}) error {
	if err:= db.Model(&Syslog{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteSyslogByCode(code string) error {
	if err := db.Where("code=?",code).Delete(Syslog{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteSyslogs(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := lichv.In (SyslogFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&Syslog{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllSyslog() error {
	if err := db.Unscoped().Delete(&Syslog{}).Error; err != nil {
		return err
	}
	return nil
}
