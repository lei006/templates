package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Setup struct {
	Name   string `orm:"column(name);pk" bson:"name"  json:"name"`
	Title  string `bson:"title"  json:"title"`
	Data   string `orm:"type(text)" bson:"data" json:"data"` //域类型
	Desc   string `orm:"size(1024)" bson:"desc" json:"desc"` //说明...
	IsShow bool   `orm:"index" bson:"is_show" json:"is_show"`

	CreatedAt time.Time `orm:"index;auto_now_add;type(datetime);default('CURRENT_TIMESTAMP')" bson:"created_at" json:"created_at"` //创建时间
	UpdatedAt time.Time `orm:"index;auto_now;type(datetime);default('CURRENT_TIMESTAMP')" bson:"updated_at" json:"updated_at"`     //修改时间
}

func init() {
	orm.RegisterModel(new(Setup))
}

type SetupModel struct {
}

var ModSetup SetupModel

func (model *SetupModel) GetAll() ([]Setup, error) {

	var objects []Setup
	_, err := orm.NewOrm().QueryTable(Setup{}).Filter("is_show", true).All(&objects)

	return objects, err
}

func (model *SetupModel) GetTotal() (int64, error) {
	count, err := orm.NewOrm().QueryTable(Setup{}).Count()
	return count, err
}

func (model *SetupModel) AddOne(name, title, data, desc string, is_show bool) (int64, error) {
	tmp_orm := orm.NewOrm()
	u := Setup{Name: name, Title: title, Data: data, Desc: desc, IsShow: is_show}
	return tmp_orm.Insert(&u)
}

func (model *SetupModel) DeleteOne(name string) error {
	tmp := Setup{Name: name}
	_, err := orm.NewOrm().Delete(&tmp)
	return err
}

func (model *SetupModel) UpdateOne(name, data string) error {

	object := Setup{}
	object.Name = name
	object.Data = data
	_, err := orm.NewOrm().Update(&object, "data")
	if err != nil {
		return err
	}

	return nil
}

func (model *SetupModel) UpdateInt64(name string, data int64) error {
	return model.UpdateOne(name, fmt.Sprintf("%d", data))
}

func (model *SetupModel) GetOne(name string) (*Setup, error) {
	tmp := Setup{Name: name}
	err := orm.NewOrm().Read(&tmp)
	if err != nil {
		return nil, err
	}
	return &tmp, nil
}

func (model *SetupModel) GetInt64(name string) (int64, error) {

	tmp, err := model.GetString(name)
	if err != nil {
		return -1, err
	}

	ret_data, err := strconv.ParseInt(tmp, 10, 64)
	if err != nil {
		return -1, err
	}

	return ret_data, nil
}

func (model *SetupModel) GetString(name string) (string, error) {

	obj, err := model.GetOne(name)
	if err != nil {
		return "", err
	}

	return obj.Data, nil
}

func (model *SetupModel) IsExist(name string) (bool, error) {

	count, err := orm.NewOrm().QueryTable(Setup{}).Filter("name", name).Count()
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

func (model *SetupModel) ReadOrCreate(name string, title string, def_data string, desc string, is_show bool) (*Setup, error) {

	// 不存在name --> 创建
	is_exist, err := model.IsExist(name)
	if err != nil {
		return nil, err
	}

	if is_exist == true {
		//处理已经存在的情况
		tmp_setup, err := model.GetOne(name)
		if err != nil {
			return nil, err
		}
		return tmp_setup, nil
	}

	_, err = model.AddOne(name, title, def_data, desc, is_show)
	if err != nil {
		return nil, err
	}
	tmp, err := model.GetOne(name)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func (model *SetupModel) DefaultString(name string, def_data string) string {

	obj, err := model.ReadOrCreate(name, name, def_data, "", false)
	if err != nil {
		return def_data
	}
	return obj.Data
}

func (model *SetupModel) DefaultInt64(name string, def_data int64) int64 {

	// 通过 DefaultString

	def_str := fmt.Sprintf("%d", def_data)
	ret_str := model.DefaultString(name, def_str)

	ret_data, err := strconv.ParseInt(ret_str, 10, 64)
	if err != nil {
		return def_data
	}

	return ret_data
}

func (model *SetupModel) SetString(name string, data string, desc string) error {

	// 不存在name --> 创建
	is_exist, err := model.IsExist(name)
	if err != nil {
		return err
	}

	if is_exist == true {
		//处理已经存在的情况
		err := model.UpdateOne(name, data)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = model.AddOne(name, name, data, desc, false)

	return err
}

func (model *SetupModel) SetInt64(name string, data int64, desc string) error {
	def_str := fmt.Sprintf("%d", data)
	return model.SetString(name, def_str, desc)
}
