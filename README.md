# HG612 Prometheus Exporter

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

## Dashbaord

![](https://raw.githubusercontent.com/jakekeeys/hg612-exporter/master/resources/dashboard.png)

## Metrics
| Metric | Type | Lables |
| --- | --- | --- |
| dsl_down_attenuation | gauge | `host` `identifier` |
| dsl_down_crc | gauge | `host` `identifier` |
| dsl_down_crc_2 | gauge | `host` `identifier` |
| dsl_down_current_rate | gauge | `host` `identifier` |
| dsl_down_current_rate_2 | gauge | `host` `identifier` |
| dsl_down_fec | gauge | `host` `identifier` |
| dsl_down_fec_2 | gauge | `host` `identifier` |
| dsl_down_hec | gauge | `host` `identifier` |
| dsl_down_hec_2 | gauge | `host` `identifier` |
| dsl_down_max_rate | gauge | `host` `identifier` |
| dsl_down_power | gauge | `host` `identifier` |
| dsl_down_snr | gauge | `host` `identifier` |
| dsl_status | gauge | `host` `identifier` `status` `modulation` `dataPath` |
| dsl_up_attenuation | gauge | `host` `identifier` |
| dsl_up_crc | gauge | `host` `identifier` |
| dsl_up_crc_2 | gauge | `host` `identifier` |
| dsl_up_current_rate | gauge | `host` `identifier` |
| dsl_up_current_rate_2 | gauge | `host` `identifier` |
| dsl_up_fec | gauge | `host` `identifier` |
| dsl_up_fec_2 | gauge | `host` `identifier` |
| dsl_up_hec | gauge | `host` `identifier` |
| dsl_up_hec_2 | gauge | `host` `identifier` |
| dsl_up_max_rate | gauge | `host` `identifier` |
| dsl_up_power | gauge | `host` `identifier` |
| dsl_up_snr | gauge | `host` `identifier` |
