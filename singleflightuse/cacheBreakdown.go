// 安装：go官方库
// 作用：解决缓存击穿的一种思路
// 参考连接：https://mp.weixin.qq.com/s/Fp9ueMRo-y8gTyxifJ84CQ

package main

import (
	"errors"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
)

var errorNotExist = errors.New("not exist")
var cacheMap = make(map[string]string)
var g singleflight.Group
func main() {
	simCache(true)
}


func simCache(gooduse bool){
	var wg sync.WaitGroup
	wg.Add(10)

	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			if gooduse{
				//采用singleflight优化
				_,err:=getData2("key")
				if err!=nil{
					log.Println(err)
					return
				}
			}else{
				//无特殊处理，模拟当缓存过期时产生的效果
				//启动时，没有缓存，会发现10个goroutine都从
				_,err:=getData("key")
				if err!=nil{
					log.Println(err)
					return
				}
			}
			//log.Println(data)
		}()
	}

	wg.Wait()
}

//采用singleflight方式解决缓存击穿的问题
func getData2(key string) (string,error){
	data,err:=getDataFromCache(key)
	if err==errorNotExist{
		v,err,_:=g.Do(key, func() (i interface{}, e error) {
			data,err = getDataFromDB(key)
			cacheMap[key] = data
			return data,err
		})
		if err!=nil{
			log.Println(err)
			return "",err
		}
		data = v.(string)
	}else if err!=nil{
		return "",err
	}
	return data,nil
}

//获取数据
func getData(key string) (string,error){
	//先从缓存中获取数据
	data,err:=getDataFromCache(key)
	if err==errorNotExist{
		//缓存数据不存在，从数据库中获取
		data,_ = getDataFromDB(key)
		//设置缓存
		cacheMap[key] = data
	}else if err!=nil{
		return "",err
	}
	return data,nil
}

//模拟从缓存中获取数据
func getDataFromCache(key string)(string,error){
	if data,ok:= cacheMap[key];ok{
		log.Printf("get %s from cache!",key)
		return data,nil
	}
	return "",errorNotExist
}
//模拟从数据库中获取数据
func getDataFromDB(key string)(string,error){
	log.Printf("get %s from database!",key)
	return "data",nil
}