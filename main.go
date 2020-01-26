package main

import (
	"fmt"
	"github.com/jakekeeys/hg612-exporter/internal/metrics"
	"github.com/jakekeeys/hg612-exporter/internal/rest"
	"github.com/jakekeeys/hg612-exporter/pkg/hg612"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "hg612 prometheus exporter",
		Usage: "a metrics exporter for the hg612",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "host",
				Usage:    "the fully qualified host for the hg612 modem",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "identifier",
				Usage:    "the identifier for the line and modem",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "bind",
				Usage: "the bind string for the http server ie :8080",
				Value: ":8080",
			},
			&cli.IntFlag{
				Name:  "interval",
				Usage: "the interval between collection in seconds",
				Value: 10,
			},
		},
		Action: func(c *cli.Context) error {
			client := hg612.New(fmt.Sprintf("http://%s", c.String("host")), http.DefaultClient)

			collector := metrics.New(client, c.String("host"), c.String("identifier"), c.Int("interval"))
			defer collector.Stop()
			collector.Start()

			server := rest.New(c.String("bind"))
			defer server.Stop()
			server.Start()

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
			<-sigChan

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Panic(err)
	}
}
