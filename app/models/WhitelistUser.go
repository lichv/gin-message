package models

import (
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type WhitelistUser struct {
	Id string `json:"id" form:"id" gorm:"id"`
	Code string `json:"code" form:"code" gorm:"code"`
	State bool `json:"state" form:"state" gorm:"state"`
}

var WhitelistUserFields = []string{"id", "code", "state"}

func ExistWhitelistUserByCode(code string) (b bool,err error) {
	var whitelistUser WhitelistUser
	err = db.Model(&WhitelistUser{}).Select("code").Where("code = ? ",code).First(&whitelistUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetWhitelistUserTotal(maps interface{}) (count int,err error) {
	err = db.Model(&WhitelistUser{}).Where("is_active = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func FindWhitelistUserByCode( code string) ( whitelistUser *WhitelistUser, err error) {
	err = db.Model(&WhitelistUser{}).Where("code = ? ",code).First(&whitelistUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &WhitelistUser{},err
	}
	return whitelistUser, err
}

func GetWhitelistUserOne( query map[string]interface{},orderBy interface{}) ( *WhitelistUser,error) {
	var whitelistUser WhitelistUser
	model := db.Model(&WhitelistUser{})
	for key, value := range query {
		b,err := lichv.In (WhitelistUserFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&whitelistUser).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &WhitelistUser{},nil
	}
	return &whitelistUser, nil
}

func GetWhitelistUserPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*WhitelistUser, []error) {
	var whitelistUsers []*WhitelistUser
	var errs []error
	model := db.Model(&WhitelistUser{})
	for key, value := range query {
		b,err := lichv.In (WhitelistUserFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&whitelistUsers)
	errs = model.GetErrors()

	return whitelistUsers, errs
}

func GetAllWhitelistUserCode( query map[string]interface{},orderBy interface{},limit int) ( []string, []error) {
	var whitelistUsers []WhitelistUser
	var errs []error
	var result []string

	model := db.Table("demo_whitelist_user").Select("code")
	for key, value := range query {
		b,err := lichv.In (WhitelistUserFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	model =model.Order(orderBy).Find(&whitelistUsers)
	errs = model.GetErrors()

	for _, v := range whitelistUsers {
		result = append(result, v.Code)
	}

	return result, errs
}

func GetWhitelistUsers( query map[string]interface{},orderBy interface{},limit int) ( []*WhitelistUser, []error) {
	var WhitelistUsers []*WhitelistUser
	var errs []error
	model := db.Model(&WhitelistUser{})
	for key, value := range query {
		b,err := lichv.In (WhitelistUserFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&WhitelistUsers).GetErrors()

	return WhitelistUsers, errs
}

func AddWhitelistUser( data map[string]interface{}) error {
	WhitelistUser := WhitelistUser{
		Id:data["id"].(string),
		Code:data["code"].(string),
		State:data["state"].(bool),
	}
	if err:= db.Create(&WhitelistUser).Error;err != nil{
		return err
	}
	return nil
}

func EditWhitelistUser( code string,data map[string]interface{}) error {
	if err:= db.Model(&WhitelistUser{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteWhitelistUserByCode(code string) error {
	if err := db.Where("code=?",code).Delete(WhitelistUser{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteWhitelistUsers(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := lichv.In (WhitelistUserFields,key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&WhitelistUser{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllWhitelistUser() error {
	if err := db.Unscoped().Delete(&WhitelistUser{}).Error; err != nil {
		return err
	}
	return nil
}
