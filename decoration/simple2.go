package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(int64, int64) int64

func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("---(%s) Time Elapsed: %v ---\n", getFuncName(f), time.Since(t))
		}(time.Now())
		return f(start, end)
	}
}
func sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (start + end) * (end - start + 1) / 2
}

func main() {
	sum1 := timedSumFunc(sum1)
	sum2 := timedSumFunc(sum2)

	fmt.Printf("%d,%d\n", sum1(1, 10), sum2(0, 10))
}
