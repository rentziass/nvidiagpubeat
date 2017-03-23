package nvidia

import (
	"encoding/csv"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/common"
)

//GPUUtilization provides interface to utilization metrics and state of GPU.
type GPUUtilization interface {
	command(env string) *exec.Cmd
	run(cmd *exec.Cmd, gpuCount int, query string, action Action) int
}

//Utilization implements one flavour of GPUCount interface.
type Utilization struct {
}

//NewUtilization returns instance of Utilization
func NewUtilization() Utilization {
	return Utilization{}
}

func (g Utilization) command(env string, query string) *exec.Cmd {
	if env == "test" {
		return exec.Command("localnvidiasmi")
	}
	return exec.Command("nvidia-smi", "--query-gpu="+query, "--format=csv")
}

func (g Utilization) run(cmd *exec.Cmd, gpuCount int, query string, action Action) []common.MapStr {
	reader := action.start(cmd)
	gpuIndex := 0
	events := make([]common.MapStr, gpuCount, 2*gpuCount)

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
