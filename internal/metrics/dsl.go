package metrics

import (
	"github.com/jakekeeys/hg612-exporter/pkg/hg612"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type dslMetricsCollector struct {
	client hg612.Client
	host   string
	identifier string

	status *prometheus.GaugeVec

	upCurrRate      *prometheus.GaugeVec
	downCurrRate    *prometheus.GaugeVec
	upCurrRate2     *prometheus.GaugeVec
	downCurrRate2   *prometheus.GaugeVec
	upMaxRate       *prometheus.GaugeVec
	downMaxRate     *prometheus.GaugeVec
	upSNR           *prometheus.GaugeVec
	downSNR         *prometheus.GaugeVec
	upAttenuation   *prometheus.GaugeVec
	downAttenuation *prometheus.GaugeVec
	upPower         *prometheus.GaugeVec
	downPower       *prometheus.GaugeVec

	downHEC  *prometheus.GaugeVec
	upHEC    *prometheus.GaugeVec
	downCRC  *prometheus.GaugeVec
	upCRC    *prometheus.GaugeVec
	downFEC  *prometheus.GaugeVec
	upFEC    *prometheus.GaugeVec
	downHEC2 *prometheus.GaugeVec
	upHEC2   *prometheus.GaugeVec
	downCRC2 *prometheus.GaugeVec
	upCRC2   *prometheus.GaugeVec
	downFEC2 *prometheus.GaugeVec
	upFEC2   *prometheus.GaugeVec
}

func newDSLMetricsCollector(client hg612.Client, host string, identifier string) dslMetricsCollector {
	status := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "status",
		},
		[]string{"host", "identifier", "status", "modulation", "dataPath"},
	)

	upCurrRate := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_current_rate",
		},
		[]string{"host", "identifier"},
	)
	downCurrRate := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_current_rate",
		},
		[]string{"host", "identifier"},
	)
	upCurrRate2 := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_current_rate_2",
		},
		[]string{"host", "identifier"},
	)
	downCurrRate2 := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_current_rate_2",
		},
		[]string{"host", "identifier"},
	)
	upMaxRate := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_max_rate",
		},
		[]string{"host", "identifier"},
	)
	downMaxRate := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_max_rate",
		},
		[]string{"host", "identifier"},
	)
	upSNR := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_snr",
		},
		[]string{"host", "identifier"},
	)
	downSNR := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_snr",
		},
		[]string{"host", "identifier"},
	)
	upAttenuation := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_attenuation",
		},
		[]string{"host", "identifier"},
	)
	downAttenuation := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_attenuation",
		},
		[]string{"host", "identifier"},
	)
	upPower := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_power",
		},
		[]string{"host", "identifier"},
	)
	downPower := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_power",
		},
		[]string{"host", "identifier"},
	)

	downHEC := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_hec",
		},
		[]string{"host", "identifier"},
	)
	upHEC := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_hec",
		},
		[]string{"host", "identifier"},
	)
	downCRC := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_crc",
		},
		[]string{"host", "identifier"},
	)
	upCRC := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_crc",
		},
		[]string{"host", "identifier"},
	)
	downFEC := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_fec",
		},
		[]string{"host", "identifier"},
	)
	upFEC := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_fec",
		},
		[]string{"host", "identifier"},
	)
	downHEC2 := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_hec_2",
		},
		[]string{"host", "identifier"},
	)
	upHEC2 := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_hec_2",
		},
		[]string{"host", "identifier"},
	)
	downCRC2 := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_crc_2",
		},
		[]string{"host", "identifier"},
	)
	upCRC2 := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_crc_2",
		},
		[]string{"host", "identifier"},
	)
	downFEC2 := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "down_fec_2",
		},
		[]string{"host", "identifier"},
	)
	upFEC2 := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dsl",
			Name:      "up_fec_2",
		},
		[]string{"host", "identifier"},
	)

	return dslMetricsCollector{
		client: client,
		host:   host,
		identifier: identifier,

		status: status,

		upCurrRate:      upCurrRate,
		downCurrRate:    downCurrRate,
		upCurrRate2:     upCurrRate2,
		downCurrRate2:   downCurrRate2,
		upMaxRate:       upMaxRate,
		downMaxRate:     downMaxRate,
		upSNR:           upSNR,
		downSNR:         downSNR,
		upAttenuation:   upAttenuation,
		downAttenuation: downAttenuation,
		upPower:         upPower,
		downPower:       downPower,

		downHEC:  downHEC,
		upHEC:    upHEC,
		downCRC:  downCRC,
		upCRC:    upCRC,
		downFEC:  downFEC,
		upFEC:    upFEC,
		downHEC2: downHEC2,
		upHEC2:   upHEC2,
		downCRC2: downCRC2,
		upCRC2:   upCRC2,
		downFEC2: downFEC2,
		upFEC2:   upFEC2,
	}
}

func (c dslMetricsCollector) collect() error {
	status, err := c.client.DSLStatus()
	if err != nil {
		return errors.Wrap(err, "error getting dsl status")
	}

	c.status.Reset()
	c.status.WithLabelValues(c.host, c.identifier, status.DSLCfg.Status, status.DSLCfg.Modulation, status.DSLCfg.DataPath).Set(1)

	c.upCurrRate.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.UpCurrRate))
	c.downCurrRate.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.DownCurrRate))
	c.upCurrRate2.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.UpCurrRate2))
	c.downCurrRate2.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.DownCurrRate2))
	c.upMaxRate.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.UpMaxRate))
	c.downMaxRate.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.DownMaxRate))
	c.upSNR.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.UpSNR))
	c.downSNR.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.DownSNR))
	c.upAttenuation.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.UpAttenuation))
	c.downAttenuation.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.DownAttenuation))
	c.upPower.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.UpPower))
	c.downPower.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLCfg.DownPower))

	c.downHEC.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.DownHEC))
	c.upHEC.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.UpHEC))
	c.downCRC.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.DownCRC))
	c.upCRC.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.UpCRC))
	c.downFEC.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.DownFEC))
	c.upFEC.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.UpFEC))
	c.downHEC2.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.DownHEC2))
	c.upHEC2.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.UpHEC2))
	c.downCRC2.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.DownCRC2))
	c.upCRC2.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.UpCRC2))
	c.downFEC2.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.DownFEC2))
	c.upFEC2.WithLabelValues(c.host, c.identifier).Set(float64(status.DSLStats.UpFEC2))

	return nil
}
