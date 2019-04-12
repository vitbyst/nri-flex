package outputs

import (
	"github.com/newrelic/nri-flex/internal/load"
	"github.com/newrelic/nri-flex/internal/logger"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
)

// CreateStatusSample creates flexStatusSample
func CreateStatusSample() {
	flexStatusSample := load.Entity.NewMetricSet("flexStatusSample")
	for counter, value := range load.FlexStatusCounter.M {
		logger.Flex("debug", flexStatusSample.SetMetric("flex."+counter, value, metric.GAUGE), "", false)
	}
}
