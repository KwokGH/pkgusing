// 安装：go get github.com/Knetic/govaluate
// 作用：用于计算任意表达式的值
// 参考链接：https://mp.weixin.qq.com/s/X6yMzAoylNbuXj4CfudQLw

package govaluateuse

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"testing"
)

func TestGovaluate_Simple(t *testing.T){
	expr,_:=govaluate.NewEvaluableExpression("a + b")
	params:=make(map[string]interface{})
	params["a"] = 10
	params["b"] = 5
	result,_:=expr.Evaluate(params)
	fmt.Println(result)
}

// 测试函数调用
func TestGovaluate_Function(t *testing.T){
	//构造函数处理逻辑 govaluate.ExpressionFunction
	functions:=map[string]govaluate.ExpressionFunction{
		"strlen":func(args ...interface{}) (interface{},error){
			if len(args) == 1 {
				if str,ok:=args[0].(string);ok{
					return len(str),nil
				}
			}
			return nil,errors.New("params is error")
		},
	}
	exprStr:="strlen('guohe')"
	//创建表达式
	expr,err:=govaluate.NewEvaluableExpressionWithFunctions(exprStr,functions)
	if err!=nil{
		t.Error(err)
		return
	}
	//指向表达式
	result,err:=expr.Eval(nil)
	if err!=nil{
		t.Error(err)
		return
	}
	fmt.Println(result)
}
