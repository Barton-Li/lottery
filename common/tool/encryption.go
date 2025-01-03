package tool

import (
	"crypto/md5"
	"fmt"
	"io"
)

//通过字符串生成MD5哈希值
func Md5ByString(str string) string {
	m:=md5.New()
	_,err:=io.WriteString(m,str)
	if err!=nil{
		panic(err)
	}
	arr:=m.Sum(nil)
	return  fmt.Sprintf("%x",arr)
}

//通过字节数组生成MD5哈希值
func Md5ByBytes(bytes []byte) string {
	return fmt.Sprintf("%x", md5.Sum(bytes))
}
