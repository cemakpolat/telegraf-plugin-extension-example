package sum

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/aggregators"
)

type ExampleAggregator struct {
	Field string
	Sum   float64
}

func (ea *ExampleAggregator) SampleConfig() string {
	return ""
}

func (ea *ExampleAggregator) Description() string {
	return ""
}

func (ea *ExampleAggregator) Add(in telegraf.Metric) {
	if value, ok := in.GetField(ea.Field); ok {
		if floatValue, ok := value.(float64); ok {
			ea.Sum += floatValue
		}
	}
}

func (ea *ExampleAggregator) Push(acc telegraf.Accumulator) {
	fields := map[string]interface{}{
		"sum": ea.Sum,
	}

	tags := map[string]string{} // Add any necessary tags

	// Add the aggregated data to the accumulator
	acc.AddFields("aggregated_data", fields, tags)
}

func (ea *ExampleAggregator) Reset() {
	ea.Sum = 0
}

func init() {
	aggregators.Add("example_aggregator", func() telegraf.Aggregator {
		return &ExampleAggregator{}
	})
}
