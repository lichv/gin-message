package services

import (
	"encoding/json"
	"gin-message/app/models"
	jwt2 "gin-message/utils/jwt"
	"github.com/gin-gonic/gin"
	lichv "github.com/lichv/go"
	"time"
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

func ExistSyslogByCode(code string) (b bool,err error) {
	b,err = models.ExistSyslogByCode(code)
	return b, err
}

func GetSyslogTotal(maps interface{}) (count int,err error) {
	count,err = models.GetSyslogTotal(map[string]interface{}{})
	return count, err
}
func GetSyslogOne( query map[string]interface{},orderBy interface{}) (*Syslog, error) {
	var nu *models.Syslog
	nu,err := models.GetSyslogOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferSyslogModel(nu),nil
}

func FindSyslogValueByCode( code string) (map[string]interface{}, error) {
	var nu *models.Syslog
	var config map[string]interface{}
	nu,err := models.GetSyslogOne(map[string]interface{}{"code":code},"code asc")
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

func GetSyslogPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (syslogs []*Syslog, total int, errs []error) {
	total,err := models.GetSyslogTotal(query)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetSyslogPages(query,orderBy,pageNum,pageSize)
	syslogs = TransferSyslogs(us)
	return syslogs,total,nil
}
func GetAllSyslogCode( query map[string]interface{},orderBy interface{},limit int)([]string,[]error){
	codes, errors := models.GetAllSyslogCode(query, orderBy, limit)
	return codes,errors
}
func GetSyslogs( query map[string]interface{},orderBy interface{},limit int) ([]*Syslog,[]error) {
	users, errors := models.GetSyslogs(query, orderBy, limit)
	syslogs := TransferSyslogs(users)
	return syslogs,errors
}

func AddSyslog( data map[string]interface{}) (int, error ){
	newid,err := models.AddSyslog(data)
	return newid,err
}

func EditSyslog( code string,data map[string]interface{}) (err error) {
	err = models.EditSyslog(code,data)
	return err
}

func DeleteSyslog(maps map[string]interface{}) (err error) {
	err = models.DeleteSyslogs(maps)
	return nil
}

func ClearAllSyslog() (err error) {
	err = models.ClearAllSyslog()
	return err
}

func TransferSyslogModel(u *models.Syslog)(syslog *Syslog){
	syslog =  &Syslog{
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
func TransferSyslogs(us []*models.Syslog) (syslogs []*Syslog) {
	for _,value := range us {
		syslog := TransferSyslogModel(value)
		syslogs = append(syslogs, syslog)
	}
	return syslogs
}


func Log(c *gin.Context,target ,operation interface{})(int,error){
	token := c.DefaultQuery("token", "")
	if token == "" {
		token = c.DefaultPostForm("token", "")
	}
	if token == "" {
		token, _ = c.Cookie("token")
	}
	if token == "" {
		token = c.GetHeader("X-TOKEN")
	}
	ip := c.ClientIP()
	user_agent := c.GetHeader("User-Agent")
	claims, err := jwt2.ParseToken(token)
	if err != nil {
		return 0,err
	}
	user := claims.Code
	newid,err := models.AddSyslog(map[string]interface{}{"user": user, "user_agent": user_agent, "token": token, "ip": ip, "target": lichv.Strval(target), "operation": lichv.Strval(operation), "time": time.Now().Unix(), "flag": 1, "state": 1})
	if err != nil {
		return 0 ,err
	}
	return newid,nil
}