// 安装： go get github.com/jinzhu/copier
// 作用：用来给结构体，map，切片等类型之间做赋值操作，支持不同类型之间的赋值
// 参考资料：https://mp.weixin.qq.com/s/KXRqSxvZAulWp2vHNQNKfw
package copieruse

import "github.com/jinzhu/copier"

type User struct{
	Name string
	Age int
}

type Employee struct{
	Name string
	Age int
	Role string
}

func (e *Employee) CopyUser(u *User){
	copier.Copy(e,u)
}
func (e *Employee) CopyManual(u *User){
	e.Name = u.Name
	e.Age = u.Age
}


