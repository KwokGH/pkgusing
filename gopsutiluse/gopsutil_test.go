// 安装： go get github.com/shirou/gopsutil
// 作用： gopsutil 是python工具库psutil的golang移植版，
// 		 用来获取各种操作系统和硬件信息，并且屏蔽了各个系统之间的差异
// 参考链接：https://mp.weixin.qq.com/s/LCmFOHQ6Pb2UgeIjqG3cSA

package gopsutiluse

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"testing"
	"time"
)

func TestCPU(t *testing.T){
	//物理核心数
	physicalCnt,_:=cpu.Counts(false)
	//逻辑核心数
	logicalCnt,_:=cpu.Counts(true)
	fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)

	//3秒内，cpu 总的使用率和每个cpu使用率
	totalPercent, _ := cpu.Percent(3*time.Second, false)
	perPercents, _ := cpu.Percent(3*time.Second, true)
	fmt.Printf("total percent:%v per percents:%v \n", totalPercent, perPercents)

	//获取cpu信息
	infos,_:=cpu.Info()
	for _,info:=range infos{
		data,_:=json.Marshal(info)

		//格式化输出的json字符串
		//data,_:=json.MarshalIndent(info,""," ")
		fmt.Println(string(data))
	}
}

func TestDisk(t *testing.T){
	//获取每个磁盘的统计信息
	mapStat, _ := disk.IOCounters()
	for name, stat := range mapStat {
		fmt.Println(name)
		data, _ := json.MarshalIndent(stat, "", "  ")
		fmt.Println(string(data))
	}

	//获取分区信息
	infos, _ := disk.Partitions(false)
	for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(data))
	}
}

func TestMemory(t *testing.T){
	swapMemory, _ := mem.SwapMemory()
	data, _ := json.MarshalIndent(swapMemory, "", " ")
	fmt.Println(string(data))
}


