package services

import (
	"gin-message/app/models"
	jwt2 "gin-message/utils/jwt"
	"github.com/gin-gonic/gin"
	lichv "github.com/lichv/go"
	"time"
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

func GetSyslogPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (syslogs []*Syslog, total int, errs []error) {
	total,err := models.GetSyslogTotal(query)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetSyslogPages(query,orderBy,pageNum,pageSize)
	syslogs = TransferSyslogs(us)
	return syslogs,total,nil
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
		User:u.User,
		UserAgent: u.UserAgent,
		Ip:u.Ip,
		Token: u.Token,
		Operation: u.Operation,
		Target: u.Target,
		Input: u.Input,
		Result: u.Result,
		Time: u.Time,
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

func Log(target,operation interface{},input interface{}) (int,error) {
	return models.AddSyslog(map[string]interface{}{"user":"system","user_agent":"system","token":"","IP":"127.0.0.1","target":lichv.Strval(target),"operation":lichv.Strval(operation),"input":lichv.Strval(input), "time": time.Now().Unix(), "flag": 1, "state": 1})
}
func LogFromContext(c *gin.Context,target ,operation interface{},input interface{})(int,error){
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
	return models.AddSyslog(map[string]interface{}{"user": user, "user_agent": user_agent, "token": token, "ip": ip, "target": lichv.Strval(target), "operation": lichv.Strval(operation),"input":lichv.Strval(input), "time": time.Now().Unix(), "flag": 1, "state": 1})
}