# pkgusing
go语言中第三方包的使用和借鉴

### copier
 安装： go get github.com/jinzhu/copier  
 作用：用来给结构体，map，切片等类型之间做赋值操作，支持不同类型之间的赋值  
 参考资料：https://mp.weixin.qq.com/s/KXRqSxvZAulWp2vHNQNKfw
 
### gojsonq
 安装：go get github.com/thedevsaddam/gojsonq  
 作用：读取 JSON 中的某一些字段  
 参考资料：https://mp.weixin.qq.com/s/ZfaH6ROWiVX9IsoxJ04oXw 

### sync/singleflight
安装：go官方库  
作用：解决缓存击穿的一种思路  
参考连接：https://mp.weixin.qq.com/s/Fp9ueMRo-y8gTyxifJ84CQ 

### sjson
安装：go get github.com/tidwall/sjson  
作用：sjson用来快速设置JSON串中的值。可以和gojsonq配合使用  
参考链接：https://mp.weixin.qq.com/s/ucfJxIw_NLb6TGrhQa3UcQ

### watermill
安装：go get github.com/ThreeDotsLabs/watermill  
作用：watermill一个异步消息解决方案，支持消息重传，消息保存， 内置多种订阅-发布实现，支持kafka和rabbitmq等  
参考链接：go get github.com/ThreeDotsLabs/watermill