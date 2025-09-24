package potd

import (
	"context"
	"fmt"
	"strings"

	"github.com/SuNNjek/hypotd/potd/apod"
	"github.com/SuNNjek/hypotd/potd/bing"
	"github.com/SuNNjek/hypotd/potd/pexels"
	"github.com/knadh/koanf/v2"
)

// PotdProvider is a picture-of-the-day provider
type PotdProvider interface {
	// DownloadPotd downloads the picture of the day and returns the path to it (or an error)
	DownloadPotd(ctx context.Context, targetDir string) (string, error)
}

func GetPotdProvider(config *koanf.Koanf) (PotdProvider, error) {
	providerName := config.String("provider")

	switch strings.ToLower(providerName) {
	case "bing":
		return bing.NewBingProvider(), nil

	case "pexels":
		return pexels.NewPexelsProvider(config.Cut("pexels"))

	case "apod":
		return apod.NewApodProvider(config.Cut("apod")), nil

	default:
		return nil, fmt.Errorf("invalid provider \"%s\"", providerName)
	}
}
