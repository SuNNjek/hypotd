package bing

import (
	"context"
	"errors"
	"fmt"
	"path"

	"gopkg.in/h2non/gentleman.v2"
)

type BingProvider struct{}

func NewBingProvider() *BingProvider {
	return &BingProvider{}
}

type image struct {
	UrlBase string `json:"urlbase"`
	Hash    string `json:"hsh"`
}

type imageArchiveResponse struct {
	Images []*image `json:"images"`
}

func (b *BingProvider) DownloadPotd(ctx context.Context, targetDir string) (string, error) {
	client := gentleman.New().UseContext(ctx)

	image, err := getPotdImage(client)
	if err != nil {
		return "", err
	}

	path := path.Join(targetDir, fmt.Sprintf("bing_%s.jpg", image.Hash))

	url := fmt.Sprintf("https://www.bing.com%s_UHD.jpg", image.UrlBase)
	req := client.Get().URL(url)
	resp, err := req.Send()
	if err != nil {
		return "", err
	}

	err = resp.SaveToFile(path)
	return path, err
}

func getPotdImage(client *gentleman.Client) (*image, error) {
	req := client.Get().
		URL("https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")

	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	var imgResp *imageArchiveResponse
	if err := resp.JSON(&imgResp); err != nil {
		return nil, err
	}

	if imgResp == nil || len(imgResp.Images) < 1 {
		return nil, errors.New("no image found")
	}

	return imgResp.Images[0], nil
}
