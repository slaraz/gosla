package v2

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type konfiguracjaEx struct {
	url, nazwa, kind, rodzaj, routingKey string
}

func MusiNowyExchanger(conf string) *Ex {
	konfEx, err := parsujYAML(conf)
	if err != nil {
		log.Fatalf("błąd parsujYAML(): %v", err)
	}
	log.Println(konfEx)

	ex, err := nowyEx(konfEx)
	if err != nil {
		log.Fatalf("błąd NowyExchanger(): %v", err)
	}

	return ex
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
