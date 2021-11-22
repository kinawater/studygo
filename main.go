package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main(){

	//mars()
	//randNum()
	//var checkStatus = checkStr("123","4")
	//fmt.Println("result is ",checkStatus)
	var point = 1;
	var result = testSwitch(strconv.Itoa(point))
	fmt.Println("result is " , result)
}

func testSwitch(checkPoint string) bool{
	fmt.Println("checkPoint is " , checkPoint)
	switch checkPoint {
		case "1":
			return true
		case "2":
			return true
		case "3":
			return true
		default:
			return false
	}
}
func checkStr(orginString string,needle string) bool {
	return strings.Contains(orginString,needle)
}
/**
	任意数字
 */
func randNum(){
	var timeStamp = time.Now().Unix()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(timeStamp)
	var num = rand.Intn(10)
	fmt.Println(num)
	fmt.Println(timeStamp)
	num = rand.Intn(10)
	fmt.Println(num)

}

/*
	去火星
 */
func mars(){
	fmt.Printf("我的体重在火星表面是: %v,火星年龄是：%v\n",185.0*0.3783,30*365/687)
	fmt.Printf("%-15v $%4v","SpaceX",94)
	fmt.Println()
	const lightSpeed = 299792
	var distance = 56000000
	fmt.Println(distance/lightSpeed,"senconds")
	var (
		a = 1
		b = 2
		c = "123"
	)
	fmt.Println(a+b,c)
	d := a*b
	d %=3
	fmt.Println(d)
}
