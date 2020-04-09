// 安装：go get github.com/go-playground/validator/v10
// 作用：对数据进行校验，如web开发中用户提交的某些数据
// 参考链接：https://mp.weixin.qq.com/s/1DcEJvQpQr-xX3LFyUf4Nw
package validatoruse

import (
	"github.com/go-playground/validator/v10"
	"testing"
	"time"
)

type TestStruct struct{
	Name string `validate:"min=6,max=10"`
	Age int `validate:"min=1,max=100"`
	Sex string `validate:"oneof=male female"`
	Email string `validate:"email"`
	// 爱好 限制成员唯一，大于等于1，小于等于3
	Hobby []string `validate:"unique,gte=1,lte=3"`
	// 成绩 unique限制值不能重复，限制大于1，小于5
	Grade map[string]int `validate:"gt=1,lt=5"`
	// 创建时间 注册时间必须小于当前的 UTC 时间
	CreateTime time.Time `validate:"lte"`
	//密码
	Password string `validate:"min=6,max=20"`
	//确认密码 跨字段约束
	SurePwd string `validate:"eqfield=Password"`
}

func TestValidator(t *testing.T){
	test:=TestStruct{
		Name:       "zhangsan",
		Age:        10,
		Sex:        "male",
		Email:      "1528933588@qq.com",
		Hobby:      []string{"daqiu","paobu"},
		Grade: 		map[string]int{"语文":60,"数学":90},
		CreateTime: time.Now().Add(-10),
		Password:   "123456",
		SurePwd:    "123456",
	}
	validate:=validator.New()
	err:=validate.Struct(test)
	if err!=nil{
		t.Error(err)
	}
}


