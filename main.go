package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello,world")

	c, d := add(10, 11)

	fmt.Println(c, d)

	str := "abc"
	str1 := str

	str = "ABC"

	fmt.Println(str1)
	point()
	newSalesOrder()

	var lanMap = []string{"JAVA", "C#", "GO", "Python"}
	for _, item := range lanMap {
		fmt.Println(item)
	}

	testMap()

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}

func add(x, y int) (a int, b int) {
	a = x + y
	b = x - y
	return a, b
}

func point() {
	a := "abcd"
	var ip = &a

	fmt.Println(*ip)
	fmt.Println(*&a)
	fmt.Println(*ip)
	fmt.Println(&a)
	fmt.Println(&ip)
	fmt.Println(&*ip)
}

func newSalesOrder() {
	var order SalesOrder
	var item SalesOrderLine
	order.DocEntry = 1
	order.DocDate = "2021-08-04"

	// order.Lines = make([]SalsOrderLine, 2)
	// order.Lines.add(ite)

	item.DocEntry = 1
	item.LineNum = 1

	fmt.Println(order.DocEntry)

}

func testMap() {
	var city map[string]string
	city = make(map[string]string, 0)
	city["BEIJING"] = "BJ"
	city[" YUEYANG"] = "YY"

	for k, v := range city {
		fmt.Println(k, " is ", v)
		if k == "BEIJING" {
			delete(city, k)
		}
	}

	for k, v := range city {
		fmt.Println(k, " is ", v)

	}
}

type SalesOrder struct {
	DocEntry int
	DocDate  string
	Lines    []SalesOrderLine
}

type SalesOrderLine struct {
	DocEntry int
	LineNum  int
	ItemCode string
	WhsCode  string
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}
