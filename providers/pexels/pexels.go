package pexels

import (
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/knadh/koanf/v2"
	"gopkg.in/h2non/gentleman.v2"
)

type PexelsProvider struct {
	apiKey string
}

type pageResponse[T any] struct {
	Page         int    `json:"page"`
	PerPage      int    `json:"per_page"`
	TotalResults int    `json:"total_results"`
	PrevPage     string `json:"prev_page,omitempty"`
	NextPage     string `json:"next_page,omitempty"`
}

type sourceType string

var (
	sourceTypeOriginal sourceType = "original"
)

type photo struct {
	Id      int                   `json:"id"`
	Sources map[sourceType]string `json:"src"`
}

type photoResponse struct {
	pageResponse[photo]

	Photos []*photo `json:"photos"`
}

func NewPexelsProvider(conf *koanf.Koanf) (*PexelsProvider, error) {
	apiKey := conf.String("apiKey")
	if apiKey == "" {
		return nil, errors.New("no API key configured")
	}

	return &PexelsProvider{
		apiKey,
	}, nil
}

func (p *PexelsProvider) DownloadPotd(ctx context.Context, targetDir string) (string, error) {
	client := gentleman.New().
		UseContext(ctx).
		SetHeader("Authorization", p.apiKey)

	photo, err := getCuratedPhotoUrl(client)
	if err != nil {
		return "", err
	}

	path := path.Join(targetDir, fmt.Sprintf("pexels_%d.jpg", photo.Id))

	url := photo.Sources[sourceTypeOriginal]
	req := client.Get().URL(url)
	resp, err := req.Send()
	if err != nil {
		return "", err
	}

	err = resp.SaveToFile(path)
	return path, err
}

func getCuratedPhotoUrl(client *gentleman.Client) (*photo, error) {
	req := client.Get().
		URL("https://api.pexels.com/v1/curated?per_page=1")

	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	var photoResp *photoResponse
	if err := resp.JSON(&photoResp); err != nil {
		return nil, err
	}

	return photoResp.Photos[0], nil
}
