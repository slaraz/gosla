package v2

import (
	"fmt"

	"gopkg.in/yaml.v3"
)


// parsujKonf zwraca pierwszą poprawnie sparsowaną konfigurację.
func parsujKonf(conf string) (konfiguracjaEx, error) {
	for i, f := range rozneParsery {
		konfEx, err := f(conf)
		if i == len(rozneParsery)-1 && err != nil {
			return konfiguracjaEx{}, fmt.Errorf("nierozpoznana konfiguracja")
		}
		if err != nil {
			continue
		}
		return konfEx, nil
	}
	// Tu nigdy nie dojdzie.
	return konfiguracjaEx{}, nil 
}

var rozneParsery = []func(string) (konfiguracjaEx, error){
	parsujYAML,
}

// YAML

type konfigYAML struct {
	Url       string `yaml:"url"` // https://www.rabbitmq.com/uri-spec.html
	Exchanger struct {
		Nazwa      string `yaml:"nazwa"`
		Kind       string `yaml:"kind"`
		Rodzaj     string `yaml:"rodzaj"`
		RoutingKey string `yaml:"routingKey,omitempty"`
	} `yaml:"exchanger"`
}

func parsujYAML(conf string) (konfiguracjaEx, error) {
	yml := konfigYAML{}
	err := yaml.Unmarshal([]byte(conf), &yml)
	if err != nil {
		return konfiguracjaEx{}, fmt.Errorf("błąd yaml.Unmarshal(): %v", err)
	}
	return konfiguracjaEx{
		url:        yml.Url,
		nazwa:      yml.Exchanger.Nazwa,
		kind:       yml.Exchanger.Kind,
		rodzaj:     yml.Exchanger.Rodzaj,
		routingKey: yml.Exchanger.RoutingKey,
	}, nil
}

// JSON

// CSV