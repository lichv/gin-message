package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-message/app/models"
	"gin-message/utils"
	"gin-message/utils/setting"
	lichv "github.com/lichv/go"
	"time"
)

type Message struct {
	Id       int    `json:"id" form:"id" gorm:"id"`
	Provider string `json:"provider" form:"provider" gorm:"provider"`
	Target   string `json:"target" form:"target" gorm:"target"`
	Template string `json:"template" form:"template" gorm:"template"`
	Params   string `json:"params" form:"params" gorm:"params"`
	Content  string `json:"content" form:"content" gorm:"content"`
	Result   string `json:"result" form:"result" gorm:"result"`
	Ptime    int    `json:"ptime" form:"ptime" gorm:"ptime"`
	Dtime    int    `json:"dtime" form:"dtime" gorm:"dtime"`
	Flag     int   `json:"flag" form:"flag" gorm:"flag"`
	State    int   `json:"state" form:"state" gorm:"state"`
}

func ExistMessageByCode(code string) (b bool, err error) {
	b, err = models.ExistMessageByCode(code)
	return b, err
}

func GetMessageTotal(maps interface{}) (count int, err error) {
	count, err = models.GetMessageTotal(map[string]interface{}{})
	return count, err
}
func GetMessageOne(query map[string]interface{}, orderBy interface{}) (*Message, error) {
	var nu *models.Message
	nu, err := models.GetMessageOne(query, orderBy)
	if err != nil {
		return nil, err
	}
	return TransferMessageModel(nu), nil
}

func GetMessagePages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) (messages []*Message, total int, errs []error) {
	total, err := models.GetMessageTotal(query)
	if err != nil {
		return nil, 0, errs
	}
	us, errs := models.GetMessagePages(query, orderBy, pageNum, pageSize)
	messages = TransferMessages(us)
	return messages, total, nil
}

func GetMessages(query map[string]interface{}, orderBy interface{}, limit int) ([]*Message, []error) {
	users, errors := models.GetMessages(query, orderBy, limit)
	messages := TransferMessages(users)
	return messages, errors
}

func AddMessage(data map[string]interface{}) (int, error) {
	newid,err := models.AddMessage(data)
	return newid,err
}

func ModifyMessage(id int, data map[string]interface{}) (err error) {
	err = models.ModifyMessage(id, data)
	return err
}

func DeleteMessage(maps map[string]interface{}) (err error) {
	err = models.DeleteMessages(maps)
	return nil
}

func ClearAllMessage() (err error) {
	err = models.ClearAllMessage()
	return err
}

func TransferMessageModel(u *models.Message) (message *Message) {
	message = &Message{
		Id:       u.Id,
		Provider: u.Provider,
		Target:   u.Target,
		Template: u.Template,
		Params:   u.Provider,
		Content:  u.Content,
		Result:   u.Result,
		Ptime:    u.Ptime,
		Dtime:    u.Dtime,
		Flag:     u.Flag,
		State:    u.State,
	}
	return
}
func TransferMessages(us []*models.Message) (messages []*Message) {
	for _, value := range us {
		message := TransferMessageModel(value)
		messages = append(messages, message)
	}
	return messages
}

func AddChuanglanMessage(target string,content string, ptime string) (int,error) {
	now := time.Now().Unix()
	if ptime != "" {
		now = int64(lichv.IntVal(ptime))
	}

	data := map[string]interface{}{"type":"sms","provider":"chuanglan","target":target,"template":"","params":"","content":content,"ptime":now,"flag":1,"state":1}
	return AddMessage(data)
}

func AddChuanglanMessageWithTemplate(target string,template_code string, ptime string, input map[string]interface{}) (int,error) {
	params, _ := json.Marshal(input)
	now := time.Now().Unix()
	if ptime != "" {
		now = int64(lichv.IntVal(ptime))
	}
	template,_ := models.FindMessageTemplateByCode(template_code)
	content, _ := utils.ParseTemplateWithParams(template.Content, input)
	data := map[string]interface{}{"type":"sms","provider":"chuanglan","target":target,"template":template_code,"params":params,"content":content,"ptime":now,"flag":1,"state":1}
	return AddMessage(data)
}

func SendSms( handle string, target string, template_code string, input map[string]interface{}) (string, error) {
	if handle == "chuanglan" {
		clConfig, ok := FindMessageConfigValueByCode("chuanglan")
		if ok != nil {
			return "", ok
		}
		appid, o := clConfig["appid"]
		if !o || appid == ""{
			return "", errors.New("appid错误")
		}
		appsecret, o := clConfig["appsecret"]
		if !o || appsecret == ""{
			return "", errors.New("密钥错误")
		}
		sign, o := clConfig["sign"]
		if !o || sign==""{
			return "", errors.New("签名错误")
		}
		template,_ := models.FindMessageTemplateByCode(template_code)
		content, _ := utils.ParseTemplateWithParams(template.Content, input)
		postdata := map[string]interface{}{"appid":lichv.Strval(appid),"appsecret":lichv.Strval(appsecret),"sign":lichv.Strval(sign),"target":target,"content":content}
		if setting.ServerSetting.RunMode=="production" {
			astr, k := lichv.CreateChuanglan(lichv.Strval(appid), lichv.Strval(appsecret), lichv.Strval(sign)).Send(target, content)
			if k != nil {
				return "", k
			}
			return *astr, nil
		}
		strs, _ := json.Marshal(postdata)
		return string(strs), nil
	}

	return "", nil
}

func HandleMesssage() (string,error) {
	var flag string
	msg, err := models.GetMessageOne(map[string]interface{}{"dtime":0,"state":1}, "id asc")
	if err != nil{
		fmt.Println(err.Error())
		return "",err
	}
	if msg.Type=="sms" && msg.Provider=="chuanglan" {
		clConfig, ok := FindMessageConfigValueByCode("chuanglan")
		if ok != nil {
			return "", ok
		}
		appid, o := clConfig["appid"]
		if !o || appid == ""{
			return "", errors.New("appid错误")
		}
		appsecret, o := clConfig["appsecret"]
		if !o || appsecret == ""{
			return "", errors.New("密钥错误")
		}
		sign, o := clConfig["sign"]
		if !o || sign==""{
			return "", errors.New("签名错误")
		}
		flag,err= SendChuanglan(lichv.Strval(appid),lichv.Strval(appsecret),lichv.Strval(sign),msg.Target,msg.Content)
		fmt.Println(flag)
		fmt.Println(err)
		if err != nil {
			return "",err
		}

	}
	if flag=="success" {
		err := models.ModifyMessage(msg.Id, map[string]interface{}{"state": -1,"dtime":time.Now().Unix()})
		if err != nil {
			return "",err
		}
	}

	return "do nothing",nil
}

func SendChuanglan(appid, appsecret, sign, target, content string) (string, error) {
	debug,e := FindSysparamValueByCode("debug")
	if e != nil {
		return "",e
	}
	fmt.Println(debug)

	isIn,err := lichv.In([]string{"false",""},debug)
	if err != nil {
		return "",err
	}
	if isIn{
		fmt.Println("是False，不用发信息")
	}else{
		fmt.Println("是True，需要发信息")
		astr, k := lichv.CreateChuanglan(lichv.Strval(appid), lichv.Strval(appsecret), lichv.Strval(sign)).Send(target, content)
		if k != nil {
			return "", k
		}
		return *astr, nil
	}

	_, err = Log( "target", "sms_chuanglan", map[string]interface{}{"appid": appid, "appsecret": appsecret, "sign": sign, "target": target, "content": content})
	if err != nil{
		return "",err
	}
	return "success",nil
}
