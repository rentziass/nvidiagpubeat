package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/deepujain/nvidiagpubeat/config"
	"github.com/deepujain/nvidiagpubeat/nvidia"
)

type Nvidiagpubeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

//New Creates the Beat object
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Nvidiagpubeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

//Run Contains the main application loop that captures data and sends it to the defined output using the publisher
func (bt *Nvidiagpubeat) Run(b *beat.Beat) error {
	logp.Info("nvidiagpubeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		metrics := nvidia.NewMetrics()
		events := metrics.Get(bt.config.Env, bt.config.Query)

		for _, event := range events {
			bt.client.PublishEvent(event)
		}
		logp.Info("Event sent")
		counter++
	}
}

//Stop Contains logic that is called when the Beat is signaled to stop
func (bt *Nvidiagpubeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
