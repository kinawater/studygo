package main

import (
	"fmt"
	"time"
)

type FilterBuilder func(next Filter) Filter

type Filter func(c *WebHandle)

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *WebHandle) {
		//打印开始时间
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 纳秒", end-start)
	}
}
