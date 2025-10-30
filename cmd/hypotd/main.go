package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path"

	"github.com/SuNNjek/hypotd/internal/config"
	"github.com/SuNNjek/hypotd/internal/potd"
	"github.com/SuNNjek/hypotd/internal/utils"
	"github.com/SuNNjek/hypotd/internal/wallpaper"
	"github.com/knadh/koanf/v2"
)

var applicationConf *koanf.Koanf

func main() {
	ctx := context.Background()

	if err := loadConfig(); err != nil {
		log.Fatalln(err)
	}

	path, err := downloadWallpaper(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	if err := setWallpaper(ctx, path); err != nil {
		log.Fatalln(err)
	}
}

func loadConfig() error {
	userConfDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	defaultConfPath := path.Join(userConfDir, "hypotd", "config.toml")

	var configPath string

	flag.StringVar(&configPath, "config", defaultConfPath, "Path to the config file")
	flag.Parse()

	config, err := config.LoadConfig(configPath)
	if err != nil {
		return err
	}

	applicationConf = config
	return nil
}

func downloadWallpaper(ctx context.Context) (string, error) {
	dir, err := utils.GetDownloadDir()
	if err != nil {
		return "", err
	}

	if err := utils.ClearOldFiles(dir, 5); err != nil {
		return "", err
	}

	provider, err := potd.GetPotdProvider(applicationConf)
	if err != nil {
		return "", err
	}

	path, err := provider.DownloadPotd(ctx, dir)
	if err != nil {
		return "", err
	}

	return path, nil
}

func setWallpaper(ctx context.Context, path string) error {
	provider, err := wallpaper.GetWallpaperProvider(applicationConf)
	if err != nil {
		return err
	}

	return provider.SetWallpaper(ctx, path)
}
