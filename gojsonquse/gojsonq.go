// go get github.com/thedevsaddam/gojsonq
// 参考连接：https://mp.weixin.qq.com/s/ZfaH6ROWiVX9IsoxJ04oXw
// github:https://github.com/thedevsaddam/gojsonq
package gojsonuse

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"log"
)

const content = `{
"name":"zhangsan",
"age":10,
"address":{
	"provice":"shanghai",
	"district":"minhang"
	}
}`
const content2 = `{
  "user": {
    "name": "dj",
    "age": 18,
    "address": {
      "provice": "shanghai",
      "district": "xuhui"
    },
    "hobbies":["chess", "programming", "game"]
  }
}`
type UserInfo struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Address Address `json:"address"`
}

type Address struct{
	Provice string `json:"provice"`
	District string `json:"district"`
}

func MarshalUserInfo(){
	obj:=UserInfo{
		Name:    "zhangsan",
		Age:     10,
		Address: Address{
			Provice:"shanghai",
			District:"minhang",
		},
	}
	b,err:=json.Marshal(obj)
	if err!=nil{
		log.Fatalln(err)
		return
	}
	log.Println(string(b))
}

func UnMarshlUserInfo(){
	var userInfo UserInfo
	json.Unmarshal([]byte(content),&userInfo)
}
func UseJsonqGetName(){
	gq:=gojsonq.New().FromString(content2)
	gq.Find("user.name")
}

//获取某个属性
func UseJsonqGetValue(){
	gq:=gojsonq.New().FromString(content2)
	provice:=gq.Find("user.address.provice")
	log.Println(provice)
	// JSONQ对象在调用Find方法时，
	// 内部会记录当前的节点，下一个查询会从上次查找的节点开始。
	// 所以需要重置下
	gq.Reset()
	hobby := gq.Find("user.hobbies.[0]")
	log.Println(hobby)
}

//数据源为json文件
func GetDataFromFile(){
	gq:=gojsonq.New().File("./data.json")
	fmt.Println(gq.Find("items.[1].price"))
}

//使用select,where等语法过滤数据
func GetDataByCondition(){
	gq:=gojsonq.New().File("./data.json")
	data:=gq.From("items").Select("id","name").Get()
	fmt.Println(data)
	gq.Reset()
	r := gq.From("items").Select("id", "name").
		Where("id", "=", 1).OrWhere("id", "=", 2).Get()
	fmt.Println(r)
}