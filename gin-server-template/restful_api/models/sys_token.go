package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Token struct {
	Id        int64     `pk:"auto" bson:"id" json:"id"`
	Token     string    `orm:"unique;index" bson:"token" json:"token"`
	UserId    int64     `orm:"index" bson:"user_id" json:"user_id"`                                   //用户名
	CreatedAt time.Time `orm:"index;auto_now_add;type(datetime)" bson:"created_at" json:"created_at"` //创建时间
	UpdatedAt time.Time `orm:"index;auto_now;type(datetime)" bson:"updated_at" json:"updated_at"`     //修改时间
	ExpiredAt time.Time `orm:"index;type(datetime)" bson:"expired_at" json:"expired_at"`              //过期时间
}

func init() {
	orm.RegisterModel(new(Token))
}

func (token *Token) ToString() string {
	return fmt.Sprintf("%+v", token)
}

type TokenModel struct {
}

var ModToken TokenModel

func (tokenMod *TokenModel) AddOne(new_token string, user_id int64, keep_second int) error {

	tmp := &Token{
		Token:     new_token,
		UserId:    user_id,
		ExpiredAt: time.Now().Add(time.Duration(keep_second) * time.Second),
	}
	_, err := orm.NewOrm().Insert(tmp)
	if err != nil {
		return err
	}

	return nil
}

func (tokenMod *TokenModel) DeleteOne(token string) error {
	_, err := orm.NewOrm().QueryTable(Token{}).Filter("token", token).Delete()
	return err
}

//删除多久前的数据....
func (tokenMod *TokenModel) DeleteOverTime(over_time time.Duration) error {
	overTime := time.Now().Add(-over_time)
	_, err := orm.NewOrm().QueryTable(Token{}).Filter("updated_at__lte", overTime).Delete()
	return err
}

func (tokenMod *TokenModel) RenewalOne(token string, keep_second int) error {

	//o.Update(&user, "Name")
	// 指定多个字段
	// o.Update(&user, "Field1", "Field2", ...)
	fmt.Println("renewal token: ", token, keep_second)
	params := orm.Params{}
	params["expired_at"] = time.Now().Add(time.Duration(keep_second) * time.Second)

	_, err := orm.NewOrm().QueryTable(Token{}).Filter("token", token).Update(params)
	if err != nil {
		return err
	}

	return nil
}

func (tokenMod *TokenModel) GetOne(token string) (*Token, error) {

	tmp := Token{}
	err := orm.NewOrm().QueryTable(Token{}).Filter("token", token).One(&tmp)
	if err != nil {
		return nil, err
	}

	return &tmp, nil
}

func (tokenMod *TokenModel) CheckToken(token string) error {

	tmp := Token{}
	err := orm.NewOrm().QueryTable(Token{}).Filter("token", token).One(&tmp)
	if err != nil {
		return err
	}
	return nil
}

func (tokenMod *TokenModel) IsExpired(token string) (*Token, error) {

	tokenInfo, err := tokenMod.GetOne(token)
	if err != nil {
		return nil, err
	}
	if tokenInfo.ExpiredAt.Unix() <= time.Now().Unix() {
		//已经过期
		return nil, errors.New("token is expired!")
	}

	fmt.Println("token time: ", tokenInfo.ExpiredAt.Unix(), time.Now().Unix())
	fmt.Println("token time: ", tokenInfo.ExpiredAt, time.Now())

	//没有过期
	return tokenInfo, nil
}
