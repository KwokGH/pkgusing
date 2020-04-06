package copieruse

import (
	"fmt"
	"testing"
)

func TestEmployee_CopyUserName(t *testing.T) {
	e:=Employee{}
	u:=User{
		Name: "zhangsan",
		Age:  10,
	}
	e.CopyUser(&u)
	t.Log(e)
}
//goos: windows
//goarch: amd64
//pkg: pkgusing/copieruse
//BenchmarkEmployee_CopyUserName-4   	  500044	      2344 ns/op
func BenchmarkEmployee_CopyUserName(b *testing.B) {
	e:=&Employee{}
	u:=&User{
		Name: "zhangsan",
		Age:  10,
	}
	for i:=0;i<b.N;i++{
		e.CopyUser(u)
	}
	fmt.Println(e)
}
//goos: windows
//goarch: amd64
//pkg: pkgusing/copieruse
//BenchmarkEmployee_CopyManual-4   	1000000000	         0.422 ns/op
func BenchmarkEmployee_CopyManual(b *testing.B) {
	e:=&Employee{}
	u:=&User{
		Name: "zhangsan",
		Age:  10,
	}
	for i:=0;i<b.N;i++{
		e.CopyManual(u)
	}
	fmt.Println(e)
}


