package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf(time.Now().Format("2006-01-02"))
	timestamp := time.Now().Unix()

	// 19 位时间戳（纳秒）
	timestampNano := time.Now().UnixNano()
	// 13 位时间戳（毫秒）
	timestampMillisecond := timestampNano / 1e6

	fmt.Printf("时间戳（秒）：%v\n", timestamp)
	fmt.Printf("时间戳（纳秒）：%v\n", timestampNano)
	fmt.Printf("时间戳（毫秒）：%v\n", timestampMillisecond)
	fmt.Printf("时间戳（纳秒转换为秒）：%v\n", timestampNano/1e9)

	datetime := time.Unix(timestampMillisecond/1e3, 0).Format("2006-01-02 15:04:05")
	fmt.Println(datetime)
}
