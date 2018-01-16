/*TODO:
 * sort the metrics output
 * output HELP information for each metric
 * output TYPE information for each metric
 */
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// parseSyslogNgMetrics takes output from `syslog-ng-ctl stats` command
// and converts it to Prometheus text metrics format
func parseSyslogNgMetrics(stats []byte) (metrics string) {
	var metricsSlice []string
	strStats := strings.Trim(string(stats[:]), "\n")
	// Remove the first line of output, which contains field
	// headers.
	for _, line := range strings.Split(strStats, "\n")[1:] {
		s := strings.Split(line, ";")
		// since `type` is a keyword in Go, we rename the `type`
		// field from SyslogNG to `typeo`.  No reason behind
		// picking this name.
		sourceName, sourceID, sourceInstance, state, typeo, number := s[0], s[1], s[2], s[3], s[4], s[5]
		metricsSlice = append(metricsSlice, fmt.Sprintf("syslog_ng_%s_total{source_name=\"%s\",source_id=\"%s\",source_instance=\"%s\",state=\"%s\"} %s\n", typeo, sourceName, sourceID, sourceInstance, state, number))
	}
	metrics = strings.Trim(strings.Join(metricsSlice, ""), "\n")
	return
}

func main() {
	stdout, err := exec.Command("/usr/local/sbin/syslog-ng-ctl", "stats").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", stdout)
	fmt.Println(parseSyslogNgMetrics(stdout))
}
