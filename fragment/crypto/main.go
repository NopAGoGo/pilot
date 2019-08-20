/* 试试crypto */
package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/satori/go.uuid"
)

const (
	_prefix  = "st"
	_contact = "_"
)


func main(){
	// 生成 appid
	var (
		thirdDeveloperName = ""
		appNo              = 1
	)
	uAppid := uuid.NewV5(uuid.NamespaceDNS, thirdDeveloperName+_contact+strconv.Itoa(appNo))
	fmt.Printf("uAppid: %s\n",uAppid)
	hasAppid := md5.Sum(uAppid[:])
	fmt.Printf("hasAppid: %b\n",hasAppid)
	appid := _prefix + hex.EncodeToString(hasAppid[:])
	fmt.Printf("appid: %s\n",appid)
	appidTrim:=_prefix + hex.EncodeToString(hasAppid[:])[8:24]
	fmt.Printf("appidTrim: %s\n",appidTrim)


	fmt.Printf("\n")

	// 刷新 secret
	seed := uuid.NewV5(uuid.NamespaceDNS, _contact+thirdDeveloperName+_contact+strconv.Itoa(appNo))
	fmt.Printf("namespace: %s\n",seed)
	uSecret := uuid.NewV5(seed, uuid.NewV4().String())
	fmt.Printf("uSecret: %s\n",uSecret)
	hasSecret := md5.Sum(uSecret[:])
	fmt.Printf("hasSecret: %b\n",hasSecret)
	secret:=base64.RawURLEncoding.EncodeToString(append(uSecret[:8],uSecret[:]...))

	fmt.Printf("secret: %s\n",secret)
	secretTrim:=secret[:32]
	fmt.Printf("secretTrim: %s\n",secretTrim)


	fmt.Printf("\n")


	// 生成 token
	hasToken:= sha256.Sum256(uAppid[:])
	fmt.Printf("hasToken: %b\n",hasToken)
	token :=base64.RawURLEncoding.EncodeToString(append(hasSecret[:],hasToken[:]...))
	fmt.Printf("token: %s\n",token)
	tokenTrim :=token[:64]
	fmt.Printf("tokenTrim: %s\n",tokenTrim)


	fmt.Printf("\n")

	decode(token)



}

func decode(token string) {

	bytesToken,err:=base64.RawURLEncoding.DecodeString(token)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Printf("%b\n",bytesToken)

}