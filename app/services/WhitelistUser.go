package services

import (
	"gin-message/app/models"
)

type WhitelistUser struct {
	Id string `json:"id" form:"id" gorm:"id"`
	Code string `json:"code" form:"code" gorm:"code"`
	State bool `json:"state" form:"state" gorm:"state"`
}

func ExistWhitelistUserByCode(code string) (b bool,err error) {
	b,err = models.ExistWhitelistUserByCode(code)
	return b, err
}

func GetWhitelistUserTotal(maps interface{}) (count int,err error) {
	count,err = models.GetWhitelistUserTotal(map[string]interface{}{})
	return count, err
}
func GetWhitelistUserOne( query map[string]interface{},orderBy interface{}) (*WhitelistUser, error) {
	var nu *models.WhitelistUser
	nu,err := models.GetWhitelistUserOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferWhitelistUserModel(nu),nil
}

func GetWhitelistUserPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (whitelistUsers []*WhitelistUser, total int, errs []error) {
	total,err := models.GetWhitelistUserTotal(query)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetWhitelistUserPages(query,orderBy,pageNum,pageSize)
	whitelistUsers = TransferWhitelistUsers(us)
	return whitelistUsers,total,nil
}
func GetAllWhitelistUserCode( query map[string]interface{},orderBy interface{},limit int)([]string,[]error){
	codes, errors := models.GetAllWhitelistUserCode(query, orderBy, limit)
	return codes,errors
}
func GetWhitelistUsers( query map[string]interface{},orderBy interface{},limit int) ([]*WhitelistUser,[]error) {
	users, errors := models.GetWhitelistUsers(query, orderBy, limit)
	whitelistUsers := TransferWhitelistUsers(users)
	return whitelistUsers,errors
}

func AddWhitelistUser( data map[string]interface{}) (err error ){
	err = models.AddWhitelistUser(data)
	return err
}

func EditWhitelistUser( code string,data map[string]interface{}) (err error) {
	err = models.EditWhitelistUser(code,data)
	return err
}

func DeleteWhitelistUser(maps map[string]interface{}) (err error) {
	err = models.DeleteWhitelistUsers(maps)
	return nil
}

func ClearAllWhitelistUser() (err error) {
	err = models.ClearAllWhitelistUser()
	return err
}

func TransferWhitelistUserModel(u *models.WhitelistUser)(whitelistUser *WhitelistUser){
	whitelistUser =  &WhitelistUser{
		Id:u.Id,
		Code:u.Code,
		State:u.State,
	}
	return
}
func TransferWhitelistUsers(us []*models.WhitelistUser) (whitelistUsers []*WhitelistUser) {
	for _,value := range us {
		whitelistUser := TransferWhitelistUserModel(value)
		whitelistUsers = append(whitelistUsers, whitelistUser)
	}
	return whitelistUsers
}
