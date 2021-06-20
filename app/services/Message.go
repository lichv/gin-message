package services

import (
	"encoding/json"
	"errors"
	"gin-message/app/models"
	"gin-message/utils"
	"gin-message/utils/setting"
	lichv "github.com/lichv/go"
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
	Flag     bool   `json:"flag" form:"flag" gorm:"flag"`
	State    bool   `json:"state" form:"state" gorm:"state"`
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
func GetAllMessageCode(query map[string]interface{}, orderBy interface{}, limit int) ([]int, []error) {
	codes, errors := models.GetAllMessageCode(query, orderBy, limit)
	return codes, errors
}
func GetMessages(query map[string]interface{}, orderBy interface{}, limit int) ([]*Message, []error) {
	users, errors := models.GetMessages(query, orderBy, limit)
	messages := TransferMessages(users)
	return messages, errors
}

func AddMessage(data map[string]interface{}) (err error) {
	err = models.AddMessage(data)
	return err
}

func EditMessage(code string, data map[string]interface{}) (err error) {
	err = models.EditMessage(code, data)
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
