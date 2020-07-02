package logsystem

import (
	"log/syslog"
)

var (
	sysLog *syslog.Writer
	logs [2]customLog
	Gfxlog customLog
)





