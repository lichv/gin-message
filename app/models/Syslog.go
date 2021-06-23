package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type Syslog struct {
	Id int `json:"id" form:"id" gorm:"id"`
	User string `json:"user" form:"user" gorm:"user"`
	UserAgent string `json:"user_agent" form:"user_agent" gorm:"user_agent"`
	Ip string `json:"ip" form:"ip" gorm:"ip"`
	Token string `json:"token" form:"token" gorm:"token"`
	Operation string `json:"operation" form:"opration" gorm:"operation"`
	Target string `json:"target" form:"target" gorm:"target"`
	Input string `json:"input" form:"input" gorm:"input"`
	Result string `json:"result" form:"options" gorm:"options"`
	Time int `json:"time" form:"time" gorm:"time"`
	Flag int `json:"flag" form:"flag" gorm:"flag"`
	State int `json:"state" form:"state" gorm:"state"`
}

var SyslogFields = []string{"id", "user", "user_agent", "ip", "token", "operation", "target", "input", "result", "time", "flag", "state"}

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
	syslog := Syslog{}
	_,ok := data["id"]
	if ok {
		syslog.Id = data["id"].(int)
	}
	_,ok = data["user"]
	if ok{
		syslog.User = lichv.Strval(data["user"])
	}
	_,ok = data["user_agent"]
	if ok {
		syslog.UserAgent = lichv.Strval(data["user_agent"])
	}
	_,ok = data["ip"]
	if ok {
		syslog.Ip = lichv.Strval(data["ip"])
	}
	_,ok = data["token"]
	if ok {
		syslog.Token = lichv.Strval(data["token"])
	}
	_,ok = data["operation"]
	if ok {
		syslog.Operation = lichv.Strval(data["operation"])
	}
	_,ok = data["target"]
	if ok {
		syslog.Target = lichv.Strval(data["target"])
	}
	_,ok = data["input"]
	if ok {
		syslog.Input = lichv.Strval(data["input"])
	}
	_,ok = data["result"]
	if ok {
		syslog.Result = lichv.Strval(data["result"])
	}
	_,ok = data["time"]
	if ok {
		syslog.Time = lichv.IntVal(data["time"])
	}
	_,ok = data["flag"]
	if ok {
		syslog.Flag = lichv.IntVal(data["flag"])
	}
	_,ok = data["state"]
	if ok {
		syslog.State = lichv.IntVal(data["state"])
	}
	if syslog.Operation== "" && syslog.Input=="" {
		return 0,errors.New("记录数据为空")
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
