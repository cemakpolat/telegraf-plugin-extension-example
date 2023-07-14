//go:build !custom || aggregators || aggregators.minmax

package all

import _ "github.com/influxdata/telegraf/plugins/aggregators/sum" // register plugin
