package wallpaper

import (
	"context"

	"github.com/knadh/koanf/v2"
)

type Provider interface {
	SetWallpaper(ctx context.Context, path string) error
}

func GetWallpaperProvider(conf *koanf.Koanf) (Provider, error) {
	customCommand := conf.String("customCommand")
	if customCommand != "" {
		return NewCustomCommandProvider(customCommand)
	}

	return NewHyprpaperProvider(), nil
}
