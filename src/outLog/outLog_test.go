package outlog

import "testing"

func TestConsole(t *testing.T) {
	outLog := OutLog{true, false}
	rec, err := outLog.Error("いるかはいるか？🐬")
	if rec == "" || err != nil {
		t.Fatal("エラーの出力失敗")
	}
}
