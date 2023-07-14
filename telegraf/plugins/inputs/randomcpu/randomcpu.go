package randomcpu

import (
	"math/rand"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type RandomCPUPlugin struct {
	// Add any configuration fields for your plugin here, if needed
}

func (r *RandomCPUPlugin) Description() string {
	return "Generate random CPU usage metrics"
}

func (r *RandomCPUPlugin) SampleConfig() string {
	return ``
}

func (r *RandomCPUPlugin) Gather(acc telegraf.Accumulator) error {
	tags := map[string]string{
		"cpu": "cpu-total",
	}

	fields := map[string]interface{}{
		"usage": rand.Float64() * 100, // Generate random CPU usage percentage
	}

	// Create a new metric and add it to the accumulator
	acc.AddFields("cpu_usage", fields, tags, time.Now())

	return nil
}

func init() {
	inputs.Add("randomcpu", func() telegraf.Input { return &RandomCPUPlugin{} })
}
