package gojsonuse

import "testing"

//goos: windows
//goarch: amd64
//pkg: pkgusing/gojsonquse
//BenchmarkUnMarshlUserInfoJson-4   	  387148	      3273 ns/op
func BenchmarkUnMarshlUserInfoJson(t *testing.B) {
	for i:=0;i<t.N;i++{
		UnMarshlUserInfo()
	}
}
//goos: windows
//goarch: amd64
//pkg: pkgusing/gojsonquse
//BenchmarkUseJsonqGetName-4   	  122455	      9857 ns/op
func BenchmarkUseJsonqGetName(t *testing.B) {
	for i:=0;i<t.N;i++{
		UseJsonqGetName()
	}
}

func TestUseJsonqGetValue(t *testing.T) {
	UseJsonqGetValue()
}
func TestGetDataFromFile(t *testing.T) {
	GetDataFromFile()
}
func TestGetDataByCondition(t *testing.T) {
	GetDataByCondition()
}