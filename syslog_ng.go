package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("/usr/local/sbin/syslog-ng-ctl", "stats").Output()
	if err != nil {
		log.Fatal(err)
	}
	strout := strings.Trim(string(out[:]), "\n")
	for _, line := range strings.Split(strout, "\n")[1:] {
		s := strings.Split(line, ";")
		sourceName, sourceId, sourceInstance, state, typeo, number := s[0], s[1], s[2], s[3], s[4], s[5]
		fmt.Printf("syslog_ng_%s_total{source_name=\"%s\",source_id=\"%s\",source_instance=\"%s\",state=\"%s\"} %s\n", typeo, sourceName, sourceId, sourceInstance, state, number)
	}
}
