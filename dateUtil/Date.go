// 日期生成函数
// 作者:nemowen
package main

import (
	"fmt"
)

func main() {
	fmt.Println(increaseData(201401))
	fmt.Println("------")
	fmt.Println(decreaseData(201412))
}

// 参数:传入年月 如:201305 类型:int
// 返回:现在到未来12个月的日期
func IncreaseData(yearmonth int) (result []int) {
	// 取得日期中的月
	var month = yearmonth % 100
	for i := 0; i < 12; i++ {
		yearmonth += 1
		month++
		if month%12 == 1 {
			yearmonth += 88
		}
		result = append(result, yearmonth)
	}
	return
}

// 参数:传入年月 如:201305 类型:int
// 返回:从现在到过去12个月的日期
func DecreaseData(yearmonth int) (result []int) {
	// 取得日期中的月
	var month = yearmonth % 100
	for i := 0; i < 12; i++ {
		yearmonth -= 1
		month--
		if month%12 == 0 {
			yearmonth -= 88
		}
		result = append(result, yearmonth)
	}
	return result
}
