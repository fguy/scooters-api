package config

import (
	"flag"
	"fmt"
	"os"
	"path"

	"go.uber.org/config"
	"go.uber.org/zap"
)

// AppConfig stands for the application's config
type AppConfig struct {
	Addr    string     `yaml:"addr"`
	DSN     string     `yaml:"dsn"`
	Logging zap.Config `yaml:"logging"`
}

func getProvider() (config.Provider, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dir = path.Join(dir, "config")
	files := []string{
		path.Join(dir, "base.yaml"),
	}
	if flag.Lookup("test.v") != nil {
		files = append(files, path.Join(dir, fmt.Sprintf("test.yaml")))
	} else {
		switch os.Getenv("APP_ENV") {
		case "production":
		case "qa":
		case "alpha":
		case "local":
		case "integration":
			files = append(files, path.Join(dir, fmt.Sprintf("%s.yaml", os.Getenv("APP_ENV"))))
		default:
			files = append(files, path.Join(dir, "development.yaml"))
		}
	}
	opts := make([]config.YAMLOption, 0, len(files)+2)
	opts = append(opts, config.Permissive())
	for _, name := range files {
		opts = append(opts, config.File(name))
	}
	return config.NewYAML(opts...)
}

// NewAppConfig -
func NewAppConfig() *AppConfig {
	provider, err := getProvider()
	if err != nil {
		panic(err) // handle error
	}

	var c AppConfig
	if err := provider.Get("app").Populate(&c); err != nil {
		panic(err) // handle error
	}
	return &c
}
