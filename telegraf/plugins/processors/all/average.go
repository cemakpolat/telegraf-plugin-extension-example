//go:build !custom || processors || processors.date

package all

import _ "github.com/influxdata/telegraf/plugins/processors/average" // register plugin
