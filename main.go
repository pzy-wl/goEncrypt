package main

import (
	"fmt"
	"goEncrypt/utils"
)

func main() {
	pub, pri, err := utils.GetEccKey()
	if err != nil {
		fmt.Printf("生成ecc密钥对有误")
	}
	s := "这是一段原文"
	c, err := utils.EccEncrypt([]byte(s), pub)
	if err != nil {
		fmt.Printf("加密失败：%v", err.Error())
	}
	msg, err := utils.EccDecrypt(c, pri)
	if err != nil {
		fmt.Printf("解密失败：%v", err.Error())
	}
	fmt.Printf("解密后的明文是；%v\n", string(msg))
	utils.EccSign([]byte(s), pri)
}
