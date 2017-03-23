package nvidia

import "github.com/elastic/beats/libbeat/common"

//GPUMetrics provides slice of metrics passed as argument for a given environment
type GPUMetrics interface {
	GetMetrics(env string, query string) []common.MapStr
}

//Metrics implements one flavour of GPUMetrics interface.
type Metrics struct {
}

//NewMetrics returns instance of Metrics
func NewMetrics() Metrics {
	return Metrics{}
}

//Get return a slice of GPU metrics
func (m Metrics) Get(env string, query string) []common.MapStr {
	gpuCount := NewCount()
	gpuCountCmd := gpuCount.command()
	count := gpuCount.run(gpuCountCmd, env)

	gpuUtilization := NewUtilization()
	gpuUtilizationCmd := gpuUtilization.command(env, query)
	events := gpuUtilization.run(gpuUtilizationCmd, count, query, NewLocal())
	return events
}
