package lko

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/go-connections/nat"
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

	config := container.Config{
		Image: "lko",
		ExposedPorts: nat.PortSet{
			nat.Port("8000/tcp"): {},
		},
	}
	hostConfig := container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("8000/tcp"): []nat.PortBinding{},
		},
	}
	networkingConfig := network.NetworkingConfig{}

	containerCreateResponse, err := docker.ContainerCreate(
		ctx,
		&config,
		&hostConfig,
		&networkingConfig,
		nil,
		"lko-container",
	)
	if err != nil {
		return err
	}

	containerId := containerCreateResponse.ID
	err = docker.ContainerStart(
		ctx,
		containerId,
		container.StartOptions{},
	)
	if err != nil {
		return err
	}

	containerJson, err := docker.ContainerInspect(
		ctx,
		containerId,
	)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", containerJson.NetworkSettings.Ports)

	println("running")
	return nil
}
