package application

import (
	"flag"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents configuration for the application
type Config struct {
	Logger   Logger `yaml:"logger"`
	RestPort string `yaml:"rest-port"`
	GrpcPorg string `yaml:"grpc-port"`
	JWT      JWT    `yamll:"jwt"`
}

// Logger has values for the logger
type Logger struct {
	Level string `yaml:"level"`
}

type JWT struct {
	Secret   string `yaml:"secret"`
	Duration int    `yaml:"duration"`
}

func readConfig() *Config {
	path := flag.String("c", "./configs/config.yaml", "set path to config yaml-file")
	flag.Parse()

	log.Printf("config file, %s", *path)

	f, err := os.Open(*path)
	if err != nil {
		log.Fatalf("cannot open %s config file: %v", *path, err)
	}
	defer f.Close()

	return readConfigFile(f)
}

// read parses yaml file to get application Config
func readConfigFile(r io.Reader) *Config {

	cfg := &Config{}
	d := yaml.NewDecoder(r)
	if err := d.Decode(cfg); err != nil {
		log.Fatalf("cannot parse config %v", err)
	}
	return cfg
}
