package main

import (
	"fmt"
	"playground/calculate"
	outlog "playground/outLog"
)

func main() {
	ol := outlog.OutLog{To_console: true, To_file: false}
	ol.Error("エラーが出力できた🎉")
	res := calculate.Calculate(3)
	fmt.Printf("計算結果：%d", res)
}
