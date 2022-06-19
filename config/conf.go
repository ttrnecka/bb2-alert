package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/disgoorg/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Discord DiscordConfig `yaml:"discord"`
	BB2     BB2Config     `yaml:"bb2"`
}

type DiscordConfig struct {
	ServerID uint64 `yaml:"server_id"`
	Token    string `yaml:"token"`
	UserID   uint64 `yaml:"user_id"`
}

type BB2Config struct {
	MatchFoundAlarm bool `yaml:"match_found_alarm"`
}

func (c *Config) InitConf(name string) {

	yamlFile, err := ioutil.ReadFile(name)
	if err != nil {
		log.Errorf(fmt.Sprintf("error reading config file:/ %v ", err))
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Errorf(fmt.Sprintf("error parsing config file: %v", err))
		os.Exit(1)
	}
}
