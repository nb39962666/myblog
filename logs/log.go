package logs

import (
	"os"
	"time"
)





//每天更新日志
func todayFilename() string {
	today := time.Now().Format("20060102")
	return "./logs/output"+ today + ".log"
}

//打开文件
func NewLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}
