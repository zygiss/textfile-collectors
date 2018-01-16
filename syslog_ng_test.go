package main

import (
	"testing"
)

func TestParseSyslogNgMetricsReturnsValue(t *testing.T) {
	input := []byte("SourceName;SourceId;SourceInstance;State;Type;Number\n1;2;3;4;5;6\n")
	metrics := parseSyslogNgMetrics(input)
	if len(metrics) == 0 {
		t.Errorf("Returned metrics are nil")
	}

}
