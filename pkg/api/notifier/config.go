package notifier

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Type string `split_words:"true" default:"sms"`
}

func loadConfig() (*config, error) {
	var conf config
	if err := envconfig.Process("NOTIFY", &conf); err != nil {
		return nil, fmt.Errorf("failed to read fcm client config, err: %s", err.Error())
	}

	return &conf, nil
}
