package hg612

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const DSLPath = "html/status/xdslStatus.asp"

type VDSLStatus struct {
	DSLCfg    DSLCfg
	DSLStats  DSLStats
	DSLUpTime int
	Time      int
}

type DSLCfg struct {
	Domain          string
	Status          string
	Modulation      string
	DataPath        string
	UpCurrRate      int
	DownCurrRate    int
	UpCurrRate2     int
	DownCurrRate2   int
	UpMaxRate       int
	DownMaxRate     int
	UpSNR           int
	DownSNR         int
	UpAttenuation   int
	DownAttenuation int
	UpPower         int
	DownPower       int
	TrafficType     string
}

type DSLStats struct {
	Domain   string
	DownHEC  int
	UpHEC    int
	DownCRC  int
	UpCRC    int
	DownFEC  int
	UpFEC    int
	DownHEC2 int
	UpHEC2   int
	DownCRC2 int
	UpCRC2   int
	DownFEC2 int
	UpFEC2   int
}

func (c HG612Client) DSLStatus() (*VDSLStatus, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.basePath, DSLPath), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating dsl status request")
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "error executing dsl status request")
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading dsl status response")
	}

	lines := strings.Split(string(all), "\n")
	if len(lines) != 4 {
		return nil, errors.New("unexpected line length in dsl status response")
	}

	dslCfgRaw := strings.Split(strings.TrimSuffix(strings.TrimPrefix(lines[0], "var DSLCfg = new Array(new stDsl(\""), "\"),null);"), "\",\"")

	if len(dslCfgRaw) != 17 {
		return nil, errors.New("unexpected dslcfg length")
	}

	var dslCfgRawInts []int
	for i := 4; i < 16; i++ {
		atoi, err := strconv.Atoi(dslCfgRaw[i])
		if err != nil {
			return nil, errors.Wrap(err, "error converting numeric dslcfg value")
		}

		dslCfgRawInts = append(dslCfgRawInts, atoi)
	}

	cfg := DSLCfg{
		Domain:          dslCfgRaw[0],
		Status:          dslCfgRaw[1],
		Modulation:      dslCfgRaw[2],
		DataPath:        dslCfgRaw[3],
		UpCurrRate:      dslCfgRawInts[0],
		DownCurrRate:    dslCfgRawInts[1],
		UpCurrRate2:     dslCfgRawInts[2],
		DownCurrRate2:   dslCfgRawInts[3],
		UpMaxRate:       dslCfgRawInts[4],
		DownMaxRate:     dslCfgRawInts[5],
		UpSNR:           dslCfgRawInts[6],
		DownSNR:         dslCfgRawInts[7],
		UpAttenuation:   dslCfgRawInts[8],
		DownAttenuation: dslCfgRawInts[9],
		UpPower:         dslCfgRawInts[10],
		DownPower:       dslCfgRawInts[11],
		TrafficType:     dslCfgRaw[16],
	}

	dslStatsRaw := strings.Split(strings.TrimSuffix(strings.TrimPrefix(lines[1], "var DSLStats = new Array(new stStats(\""), "\"),null);"), "\",\"")
	if len(dslStatsRaw) != 13 {
		return nil, errors.New("unexpected dslstats length")
	}

	var dslStatsRawInts []int
	for i := 1; i < 13; i++ {
		atoi, err := strconv.Atoi(dslStatsRaw[i])
		if err != nil {
			return nil, errors.Wrap(err, "error converting numeric dsl stats value")
		}

		dslStatsRawInts = append(dslStatsRawInts, atoi)
	}

	stats := DSLStats{
		Domain:   dslStatsRaw[0],
		DownHEC:  dslStatsRawInts[0],
		UpHEC:    dslStatsRawInts[1],
		DownCRC:  dslStatsRawInts[2],
		UpCRC:    dslStatsRawInts[3],
		DownFEC:  dslStatsRawInts[4],
		UpFEC:    dslStatsRawInts[5],
		DownHEC2: dslStatsRawInts[6],
		UpHEC2:   dslStatsRawInts[7],
		DownCRC2: dslStatsRawInts[8],
		UpCRC2:   dslStatsRawInts[9],
		DownFEC2: dslStatsRawInts[10],
		UpFEC2:   dslStatsRawInts[11],
	}

	dslUpTime, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(lines[2], "var DslUpTime = \""), "\";"))
	if err != nil {
		return nil, errors.Wrap(err, "error converting dsl uptime")
	}

	time, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(lines[3], "var time = "), ";"))
	if err != nil {
		return nil, errors.Wrap(err, "error converting dsl time")
	}

	return &VDSLStatus{
		DSLCfg:    cfg,
		DSLStats:  stats,
		DSLUpTime: dslUpTime,
		Time:      time,
	}, nil
}
