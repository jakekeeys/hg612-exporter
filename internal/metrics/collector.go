package metrics

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/jakekeeys/hg612-exporter/pkg/hg612"
)

type Collector interface {
	Collect()
	Start()
	Stop()
}

type MetricsCollector struct {
	ctx                    context.Context
	collectIntervalSeconds int
	dslMetricsCollector    dslMetricsCollector
}

func New(client hg612.Client, host string, identifier string, collectIntervalSeconds int) Collector {
	return MetricsCollector{
		collectIntervalSeconds: collectIntervalSeconds,
		ctx:                    context.Background(),
		dslMetricsCollector:    newDSLMetricsCollector(client, host, identifier),
	}
}

func (c MetricsCollector) Collect() {
	err := c.dslMetricsCollector.collect()
	if err != nil {
		logrus.Error(errors.Wrap(err, "error collecting dsl metrics"))
	}
}

func (c MetricsCollector) Start() {
	go func() {
		for {
			select {
			case <-time.After(time.Second * time.Duration(c.collectIntervalSeconds)):
				c.Collect()
			case <-c.ctx.Done():
				return
			}
		}
	}()
}

func (c MetricsCollector) Stop() {
	c.ctx.Done()
}
