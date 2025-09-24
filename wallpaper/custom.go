package wallpaper

import (
	"bytes"
	"context"
	"os/exec"
	"text/template"
)

func executeTemplateString(tmpl *template.Template, input any) (string, error) {
	var b bytes.Buffer
	if err := tmpl.Execute(&b, input); err != nil {
		return "", err
	}

	return b.String(), nil
}

type CustomCommandProvider struct {
	template *template.Template
}

func NewCustomCommandProvider(templateString string) (*CustomCommandProvider, error) {
	template, err := template.New("customCommand").Parse(templateString)
	if err != nil {
		return nil, err
	}

	return &CustomCommandProvider{
		template,
	}, nil
}

type customCommandData struct {
	Path string
}

func (p *CustomCommandProvider) SetWallpaper(ctx context.Context, path string) error {
	data := customCommandData{
		Path: path,
	}

	command, err := executeTemplateString(p.template, data)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", command)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
