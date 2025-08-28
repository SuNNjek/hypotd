package providers

import (
	"context"
	"fmt"
	"strings"

	"github.com/SuNNjek/hypotd/providers/bing"
	"github.com/knadh/koanf/v2"
)

// PotdProvider is a picture-of-the-day provider
type PotdProvider interface {
	// DownloadPotd downloads the picture of the day and returns the path to it (or an error)
	DownloadPotd(ctx context.Context, targetDir string) (string, error)
}

func GetConfiguredProvider(config *koanf.Koanf) (PotdProvider, error) {
	providerName := config.String("provider")

	switch strings.ToLower(providerName) {
	case "bing":
		return bing.NewBingProvider(), nil

	default:
		return nil, fmt.Errorf("invalid provider \"%s\"", providerName)
	}
}
