package lko

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type App struct {
}

func InitApp() *App {
	return &App{}
}

func (app *App) Run() error {
	docker, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Tags:   []string{"lko"},
		Labels: map[string]string{},
	}

	build, err := archive.TarWithOptions("../../test/server", &archive.TarOptions{})
	if err != nil {
		return err
	}

	ctx := context.Background()

	response, err := docker.ImageBuild(
		ctx,
		build,
		opts,
	)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	defer response.Body.Close()
	println("running")
	return nil
}
