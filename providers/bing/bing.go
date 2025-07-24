package bing

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type image struct {
	UrlBase string `json:"urlbase"`
	Hash    string `json:"hsh"`
}

type imageArchiveResponse struct {
	Images []*image `json:"images"`
}

func DownloadPotd() (string, error) {
	image, err := getPotdImage()
	if err != nil {
		return "", err
	}

	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", nil
	}

	downloadDir := path.Join(userCacheDir, "hypotd")
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		return "", nil
	}

	path := path.Join(downloadDir, fmt.Sprintf("bing_%s.jpg", image.Hash))
	file, err := os.Create(path)
	if err != nil {
		return "", nil
	}

	defer file.Close()

	url := fmt.Sprintf("https://www.bing.com%s_UHD.jpg", image.UrlBase)
	imgResp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", nil
	}

	defer imgResp.Body.Close()

	if _, err := io.Copy(file, imgResp.Body); err != nil {
		return "", nil
	}

	return path, nil
}

func getPotdImage() (*image, error) {
	resp, err := http.DefaultClient.Get("https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := io.TeeReader(resp.Body, os.Stdout)
	decoder := json.NewDecoder(reader)

	var imgResp *imageArchiveResponse
	if err := decoder.Decode(&imgResp); err != nil {
		return nil, err
	}

	if imgResp == nil || len(imgResp.Images) < 1 {
		return nil, errors.New("no image found")
	}

	return imgResp.Images[0], nil
}
