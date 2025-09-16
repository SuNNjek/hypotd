package apod

import (
	"context"
	"fmt"
	"log"
	"path"

	"github.com/knadh/koanf/v2"
	"gopkg.in/h2non/gentleman.v2"
)

type ApodProvider struct {
	apiKey string
}

type apiResponse struct {
	Date  string `json:"date"`
	Url   string `json:"url"`
	HdUrl string `json:"hdurl,omitempty"`
}

func NewApodProvider(conf *koanf.Koanf) *ApodProvider {
	apiKey := conf.String("apiKey")
	if apiKey == "" {
		log.Println("No API key set for APOD provider, falling back to demo key...")
		apiKey = "DEMO_KEY"
	}

	return &ApodProvider{
		apiKey,
	}
}

func (p *ApodProvider) DownloadPotd(ctx context.Context, targetDir string) (string, error) {
	client := gentleman.New().UseContext(ctx)

	apiResp, err := p.getApiResponse(client)
	if err != nil {
		return "", err
	}

	path := path.Join(targetDir, fmt.Sprintf("apod_%s.jpg", apiResp.Date))
	url := getPictureUrl(apiResp)

	req := client.Get().URL(url)
	resp, err := req.Send()
	if err != nil {
		return "", err
	}

	err = resp.SaveToFile(path)
	return path, err
}

func (p *ApodProvider) getApiResponse(client *gentleman.Client) (*apiResponse, error) {
	req := client.Get().
		URL("https://api.nasa.gov/planetary/apod").
		SetQuery("api_key", p.apiKey)

	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	var result *apiResponse
	if err := resp.JSON(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func getPictureUrl(response *apiResponse) string {
	if response.HdUrl != "" {
		return response.HdUrl
	}

	return response.Url
}
