package models

import (
	"errors"
	"gin-server-template/entity"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       int64  `pk:"auto" bson:"id" json:"id"`
	Username string `orm:"unique;index" bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Nickname string `bson:"nickname" json:"nickname"`
	IsEnable bool   `bson:"is_enable" json:"is_enable"`
	Avatar   string `bson:"avatar" json:"avatar"`
	Desc     string `orm:"size(8196)" bson:"desc" json:"desc"`       //
	ExData   string `orm:"type(text)" bson:"ex_data" json:"ex_data"` //扩展数据..

	CreatedAt time.Time `orm:"index;auto_now_add;type(datetime)" bson:"created_at" json:"created_at"` //创建时间
	UpdatedAt time.Time `orm:"index;auto_now;type(datetime)" bson:"updated_at" json:"updated_at"`     //修改时间
}

func init() {
	orm.RegisterModel(new(User))
}

type UserModel struct {
}

var ModUser UserModel

func (userMod *UserModel) GetAll() ([]User, error) {

	var users []User
	_, err := orm.NewOrm().QueryTable(User{}).All(&users)

	return users, err
}

func (userMod *UserModel) GetTotal() (int64, error) {
	count, err := orm.NewOrm().QueryTable(User{}).Count()
	return count, err
}

func (userMod *UserModel) AddOne(u User) (int64, error) {
	tmp_orm := orm.NewOrm()
	u.Id = 0
	return tmp_orm.Insert(&u)
}

func (userMod *UserModel) DeleteOne(user_id int64) error {
	tmp := User{Id: user_id}
	_, err := orm.NewOrm().Delete(&tmp)
	return err
}

func (userMod *UserModel) UpdateOne(user_id int64, user *User) error {

	user.Id = user_id
	_, err := orm.NewOrm().Update(user, "nickname", "password", "avatar", "desc", "is_enable", "ex_data", "desc")
	if err != nil {
		return err
	}

	return nil
}

func (model *UserModel) SetOne(id int64, datafield entity.DataField) error {

	object := User{Id: id}

	object.Id = id
	_, err := orm.NewOrm().QueryTable(User{}).Filter("id", id).Update(orm.Params{datafield.Name: datafield.Data})
	if err != nil {
		return err
	}

	return nil
}

func (userMod *UserModel) GetOne(user_id int64) (*User, error) {
	tmp := User{Id: user_id}
	err := orm.NewOrm().Read(&tmp)
	if err != nil {
		return nil, err
	}
	return &tmp, nil
}

func (userMod *UserModel) GetOneByUsername(username string) (*User, error) {
	tmp := User{}
	err := orm.NewOrm().QueryTable(User{}).Filter("username", username).One(&tmp)
	if err != nil {
		return nil, err
	}
	return &tmp, nil
}

//用户是否存在..
func (userMod *UserModel) IsExist(username string) error {

	count, err := orm.NewOrm().QueryTable(User{}).Filter("username", username).Count()
	if err != nil {
		return err
	}

	if count != 1 {
		return errors.New("user is not exist: " + username)
	}

	return nil
}
