package main

import (
	"playground/calculate"
	outlog "playground/outLog"
	"testing"
)

func TestCal(t *testing.T) {
	result := calculate.Calculate(2)
	if result != 2*2 {
		t.Fatal("calculateの答えが違うよ。")
	}
}

func TestConsole(t *testing.T) {
	outLog := outlog.OutLog{To_console: true, To_file: false}
	rec, err := outLog.Error("いるかはいるか？🐬")
	if rec == "" || err != nil {
		t.Fatal("エラーの出力失敗")
	}
}
