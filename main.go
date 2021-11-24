package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main(){

}
/**
猜数字
 */
func guessNum(insertNum int)  {
	temp := 0
	n := 0
	for {
		temp = randNum(int64(n))
		n++
		if (temp != insertNum) {
			println("Not equal,Now is ",temp)
		} else {
			println("Equal! Now is",temp)
			break
		}
	}
}
/**
任意数字
*/
func randNum(seedOther int64) int{
	var timeStamp = time.Now().Unix()
	r := rand.New(rand.NewSource(timeStamp+seedOther))
	num := r.Intn(100)
	return num
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
