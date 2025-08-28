package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path"

	"github.com/SuNNjek/hypotd/config"
	"github.com/SuNNjek/hypotd/hyprpaper"
	"github.com/SuNNjek/hypotd/providers"
	"github.com/SuNNjek/hypotd/utils"
)

func main() {
	ctx := context.Background()

	configPath, err := getConfigPath()
	if err != nil {
		log.Fatalln(err)
	}

	conf, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	dir, err := utils.GetDownloadDir()
	if err != nil {
		log.Fatalln(err)
	}

	if err := utils.ClearOldFiles(dir, 5); err != nil {
		log.Fatalln(err)
	}

	provider, err := providers.GetConfiguredProvider(conf)
	if err != nil {
		log.Fatalln(err)
	}

	path, err := provider.DownloadPotd(ctx, dir)
	if err != nil {
		log.Fatalln(err)
	}

	if err := hyprpaper.SetWallpaper(ctx, path); err != nil {
		log.Fatalln(err)
	}
}

func getConfigPath() (string, error) {
	userConfDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	defaultConfPath := path.Join(userConfDir, "hypotd", "config.toml")

	var configPath string

	flag.StringVar(&configPath, "config", defaultConfPath, "Path to the config file")
	flag.Parse()

	return configPath, err
}
