package producer

import "github.com/kelseyhightower/envconfig"

const envPrefix = "PRODUCER"

type Authentication struct {
	Type     string `json:"type" envconfig:"AUTH_TYPE"`
	Username string `json:"username" envconfig:"AUTH_USERNAME"`
	Password string `json:"password" envconfig:"AUTH_PASSWORD"`
}

type Configuration struct {
	ProduceRate    uint            `json:"produceRate" envconfig:"PRODUCE_RATE"`
	Broker         string          `json:"broker" envconfig:"BROKER_ADDRESS"`
	Topic          string          `json:"topic" envconfig:"TOPIC_NAME"`
	Authentication *Authentication `json:"authentication"`
}

func GetConfig() (*Configuration, error) {
	c := Configuration{}

	err := envconfig.Process(envPrefix, &c)
	if err != nil {
		return nil, err
	}

	if err = validate(&c); err != nil {
		return nil, err
	}

	if c.Authentication.Type == "" {
		c.Authentication = nil
	}

	return &c, nil
}

func validate(_ *Configuration) error {
	return nil
}
