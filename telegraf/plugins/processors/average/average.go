package average

import (
	"math"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
)

type AverageProcessor struct {
	Field string
}

func (ap *AverageProcessor) SampleConfig() string {
	return ""
}

func (ap *AverageProcessor) Description() string {
	return ""
}

func (ap *AverageProcessor) Apply(in ...telegraf.Metric) []telegraf.Metric {
	for _, metric := range in {
		if value, ok := metric.GetField(ap.Field); ok {
			if floatValue, ok := value.(float64); ok {
				// Divide the value by 2
				newValue := floatValue / 2.0

				// Round the new value to 2 decimal places
				roundedValue := math.Round(newValue*100) / 100

				// Set the new value as a new field in the metric
				metric.AddField(ap.Field+"_processed", roundedValue)
			}
		}
	}

	return in
}

func init() {
	processors.Add("average_processor", func() telegraf.Processor {
		return &AverageProcessor{}
	})
}
