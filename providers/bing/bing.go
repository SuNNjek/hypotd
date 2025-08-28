package bing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/SuNNjek/hypotd/utils"
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
	image, err := getPotdImage(ctx)
	if err != nil {
		return "", err
	}

	path := path.Join(targetDir, fmt.Sprintf("bing_%s.jpg", image.Hash))
	file, err := os.Create(path)
	if err != nil {
		return "", nil
	}

	defer file.Close()

	url := fmt.Sprintf("https://www.bing.com%s_UHD.jpg", image.UrlBase)
	imgBody, err := utils.GetContext(ctx, url)
	if err != nil {
		return "", nil
	}

	defer imgBody.Close()

	if _, err := io.Copy(file, imgBody); err != nil {
		return "", nil
	}

	return path, nil
}

func getPotdImage(ctx context.Context) (*image, error) {
	body, err := utils.GetContext(ctx, "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
	if err != nil {
		return nil, err
	}

	defer body.Close()

	decoder := json.NewDecoder(body)

	var imgResp *imageArchiveResponse
	if err := decoder.Decode(&imgResp); err != nil {
		return nil, err
	}

	if imgResp == nil || len(imgResp.Images) < 1 {
		return nil, errors.New("no image found")
	}

	return imgResp.Images[0], nil
}
