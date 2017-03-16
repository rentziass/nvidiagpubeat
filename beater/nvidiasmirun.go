package beater

import (
	"bufio"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/common"
)

func Run(c NVIDIACommands, query string) []common.MapStr {
	smiCmd := c.GetSMI()
	stdout, _ := smiCmd.StdoutPipe()
	smiCmd.Start()
	reader := bufio.NewReader(stdout)
	gpuIndex := 0
	events := make([]common.MapStr, c.getNumGpus(), 2*c.getNumGpus())

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// Ignore header
		if strings.Contains(line, "utilization") {
			continue
		}
		// Remove units put by nvidia-smi
		line = strings.Replace(line, " %", "", -1)
		line = strings.Replace(line, " MiB", "", -1)
		line = strings.Replace(line, " P", "", -1)
		line = strings.Replace(line, " ", "", -1)

		r := csv.NewReader(strings.NewReader(line))
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		headers := strings.Split(query, ",")
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"gpuIndex":   gpuIndex,
		}
		for i := 0; i < len(record); i++ {
			value, _ := strconv.Atoi(record[i])
			event.Put(headers[i], value)
		}
		events[gpuIndex] = event
		gpuIndex++
	}
	return events
}
