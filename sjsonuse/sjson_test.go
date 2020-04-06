// 安装：go get github.com/tidwall/sjson
// 作用：sjson用来快速设置JSON串中的值。可以和gojsonq配合使用
// 参考链接：https://mp.weixin.qq.com/s/ucfJxIw_NLb6TGrhQa3UcQ
package sjsonuse

import (
	"github.com/tidwall/sjson"
	"testing"
)

const jsonStr = `{"name":{"first":"li","last":"dj"},"age":18,"isStudent":false,"deposit":1000.5,"courses":["yuwen","shuxue","yingyu"]}`

// 支持的类型有 nil/bool/int/float/string
// 如果传入的是不支持的类型，则会json.Marshal为字符串后加入到对应的键路径
func TestSJSON(t *testing.T){
	//测试用例
	var tests = []struct{ path string; value interface{}; result string }{
		//测试string类型
		{"name.last", "zangsan",
			`{"name":{"first":"li","last":"zangsan"},"age":18,"isStudent":false,"deposit":1000.5,"courses":["yuwen","shuxue","yingyu"]}`},

		//测试int类型
		{"age",20,
			`{"name":{"first":"li","last":"dj"},"age":20,"isStudent":false,"deposit":1000.5,"courses":["yuwen","shuxue","yingyu"]}`},

		//测试bool类型
		{"isStudent",true,
			`{"name":{"first":"li","last":"dj"},"age":18,"isStudent":true,"deposit":1000.5,"courses":["yuwen","shuxue","yingyu"]}`},

		//测试不支持的类型map
		{"name",map[string]interface{}{"first":"guo","last":"he"},
			`{"name":{"first":"guo","last":"he"},"age":18,"isStudent":false,"deposit":1000.5,"courses":["yuwen","shuxue","yingyu"]}`},

		//测试nil
		{"name",nil,
			`{"name":null,"age":18,"isStudent":false,"deposit":1000.5,"courses":["yuwen","shuxue","yingyu"]}`},

		//测试数组,-1或数组的长度，标识在数组后添加一个新元素
		//使用的索引超出数组长度，会中间添加许多null值
		{"courses.1","lishi",
			`{"name":{"first":"li","last":"dj"},"age":18,"isStudent":false,"deposit":1000.5,"courses":["yuwen","lishi","yingyu"]}`},

		//测试删除
		{"name",-1,
			`{"age":18,"isStudent":false,"deposit":1000.5,"courses":["yuwen","shuxue","yingyu"]}`},
	}
	for _,tt:=range tests{
		if tt.value==-1 {
			if val,_:=sjson.Delete(jsonStr,tt.path);val!=tt.result{
				t.Errorf("sjson.Set %s,%v; got %s; expect %s; \n",tt.path,tt.value,val, tt.result)
			}
			continue
		}
		if val,_:=sjson.Set(jsonStr,tt.path,tt.value);val!=tt.result{
			t.Errorf("sjson.Set %s,%v; got %s; expect %s; \n",tt.path,tt.value,val, tt.result)
		}
	}
}
