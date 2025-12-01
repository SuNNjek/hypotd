package e621

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/knadh/koanf/v2"
	"gopkg.in/h2non/gentleman.v2"
)

type E621Provider struct {
	username string
	apiKey   string
	sfw      bool
	tags     []string
}

func NewE621Provider(conf *koanf.Koanf) (*E621Provider, error) {
	username := conf.String("username")
	if username == "" {
		return nil, errors.New("no username configured")
	}

	apiKey := conf.String("apiKey")
	if apiKey == "" {
		return nil, errors.New("no API key configured")
	}

	sfw := conf.Bool("sfw")
	tags := conf.Strings("tags")

	return &E621Provider{
		username,
		apiKey,
		sfw,
		tags,
	}, nil
}

func (p *E621Provider) DownloadPotd(ctx context.Context, targetDir string) (string, error) {
	client, err := p.createClient(ctx)
	if err != nil {
		return "", err
	}

	post, err := p.getTopPost(client)
	if err != nil {
		return "", err
	}

	return p.downloadPost(client, post, targetDir)
}

func (p *E621Provider) createClient(ctx context.Context) (*gentleman.Client, error) {
	var baseUrl string
	if p.sfw {
		baseUrl = "https://e926.net/"
	} else {
		baseUrl = "https://e621.net/"
	}

	authHeader := fmt.Sprintf("%s:%s", p.username, p.apiKey)
	encAuthHeader := base64.StdEncoding.EncodeToString([]byte(authHeader))

	client := gentleman.New().
		UseContext(ctx).
		URL(baseUrl).
		AddHeader("Authorization", fmt.Sprintf("Basic %s", encAuthHeader)).
		AddHeader("User-Agent", "hypotd/1.0 (by Sunner on e621)")

	return client, nil
}

func (p *E621Provider) getTopPost(client *gentleman.Client) (*Post, error) {
	req := client.Get().
		Path("/posts.json").
		SetQuery("tags", strings.Join(p.tags, " "))

	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	var result *PostsResponse
	if err := resp.JSON(&result); err != nil {
		return nil, err
	}

	allowedExtensions := []string{
		"png",
		"jpg",
		"jpeg",
		"webp",
	}

	for _, post := range result.Posts {
		if slices.ContainsFunc(allowedExtensions, func(ext string) bool {
			return strings.EqualFold(ext, post.File.Extention)
		}) {
			return post, nil
		}
	}

	return nil, errors.New("no compatible posts found")
}

func (p *E621Provider) downloadPost(client *gentleman.Client, post *Post, targetDir string) (string, error) {
	filename := fmt.Sprintf("e621_%d.%s", post.Id, post.File.Extention)
	downloadPath := path.Join(targetDir, filename)

	if verifyExistingFile(post, downloadPath) {
		return downloadPath, nil
	}

	req := client.Get().URL(post.File.Url)
	resp, err := req.Send()
	if err != nil {
		return "", err
	}

	if err := resp.SaveToFile(downloadPath); err != nil {
		return "", err
	}

	return downloadPath, nil
}

func verifyExistingFile(post *Post, filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}

	algo := md5.New()
	if _, err := io.Copy(algo, file); err != nil {
		return false
	}

	hash := hex.EncodeToString(algo.Sum(nil))
	return strings.EqualFold(hash, post.File.Md5)
}
