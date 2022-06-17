package outlog

import (
	"log"
	"os"
	"time"
)

type OutLog struct {
	To_console bool
	To_file    bool
}

func (ol OutLog) Error(mess string) (rec string, err error) {
	// 現在時刻取得
	t := time.Now()

	// ログのライン
	line := "[" + t.Format(time.RFC3339) + "][Error] - " + mess

	// ファイルへの出力
	if ol.To_file {
		f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		defer f.Close()

		_, er := f.WriteString(line + "\n")
		if er != nil {
			log.Fatal(er)
			return "", er
		}
	}

	// コンソールへの出力
	if ol.To_console {
		log.Println(line)
	}

	return line, nil
}
