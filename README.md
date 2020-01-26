# HG612 Prometheus Exporter

[![Docker Repository on Quay](https://quay.io/repository/jakekeeys/hg612-exporter/status "Docker Repository on Quay")](https://quay.io/repository/jakekeeys/hg612-exporter)

A prometheus exporter for the HG612 modem

![](https://i0.wp.com/codeblog.dotsandbrackets.com/wp-content/uploads/2017/01/prometheus-logo.jpg?resize=231%2C231) ![](https://kitz.co.uk/routers/images/huawei_echolife_hg612.jpg)

*Requires unlocked firmware with exposed metrics*

see https://kitz.co.uk/routers/hg612unlock.htm for flashing instructions and firmware

see also https://support.aa.net.uk/Router_-_EchoLife_HG612


## Usage

```
NAME:
   hg612 prometheus exporter - a metrics exporter for the hg612

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value        the fully qualified host for the hg612 modem
   --identifier value  the identifier for the line and modem
   --bind value        the bind string for the http server ie :8080 (default: ":8080")
   --interval value    the interval between collection in seconds (default: 10)
   --help, -h          show help (default: false)
```

## Build
go
```
go build .
```

docker 
```
docker build .
```

## Run
go 
```
go run main.go --host <host> --identifier <identifier>
```

docker 
```
docker run docker run quay.io/jakekeeys/hg612-exporter --host <host> --identifier <identifier>
```

## Metrics

| Domain | Implemented | Status
| --- | --- | --- |
| dsl | ✓ | alpha |
| device | - | - |
| wan | - | - |
| lan | - | - |
| atm | - | - |

| Domain | Metric Name | Type | Labels |
| --- | --- | --- | --- |
| dsl | dsl_down_attenuation | gauge | `host` `identifier` |
| dsl | dsl_down_crc | gauge | `host` `identifier` |
| dsl | dsl_down_crc_2 | gauge | `host` `identifier` |
| dsl | dsl_down_current_rate | gauge | `host` `identifier` |
| dsl | dsl_down_current_rate_2 | gauge | `host` `identifier` |
| dsl | dsl_down_fec | gauge | `host` `identifier` |
| dsl | dsl_down_fec_2 | gauge | `host` `identifier` |
| dsl | dsl_down_hec | gauge | `host` `identifier` |
| dsl | dsl_down_hec_2 | gauge | `host` `identifier` |
| dsl | dsl_down_max_rate | gauge | `host` `identifier` |
| dsl | dsl_down_power | gauge | `host` `identifier` |
| dsl | dsl_down_snr | gauge | `host` `identifier` |
| dsl | dsl_status | gauge | `host` `identifier` `status` `modulation` `dataPath` |
| dsl | dsl_up_attenuation | gauge | `host` `identifier` |
| dsl | dsl_up_crc | gauge | `host` `identifier` |
| dsl | dsl_up_crc_2 | gauge | `host` `identifier` |
| dsl | dsl_up_current_rate | gauge | `host` `identifier` |
| dsl | dsl_up_current_rate_2 | gauge | `host` `identifier` |
| dsl | dsl_up_fec | gauge | `host` `identifier` |
| dsl | dsl_up_fec_2 | gauge | `host` `identifier` |
| dsl | dsl_up_hec | gauge | `host` `identifier` |
| dsl | dsl_up_hec_2 | gauge | `host` `identifier` |
| dsl | dsl_up_max_rate | gauge | `host` `identifier` |
| dsl | dsl_up_power | gauge | `host` `identifier` |
| dsl | dsl_up_snr | gauge | `host` `identifier` |

### Grafana Dashboard Preview

![](https://raw.githubusercontent.com/jakekeeys/hg612-exporter/master/resources/dashboard.png)
