package helpers

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func SetDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error creating Docker client: %w", err)
	}

	return cli, nil
}

func GetContainerByName(containerName string) (*types.Container, error) {
	cli, err := SetDockerClient()

	if err != nil {
		return nil, fmt.Errorf("error creating Docker client: %w", err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("error listing containers: %w", err)
	}

	for _, container := range containers {
		if container.Names[0] == "/"+containerName {
			return &container, nil
		}
	}

	return nil, fmt.Errorf("container not found: %s", containerName)
}

func ContainerExists(containerName string) bool {
	_, err := GetContainerByName(containerName)
	return err == nil
}

func ExecCommandInContainer(containerID string, command []string) error {
	cli, err := SetDockerClient()
	if err != nil {
		return fmt.Errorf("error creating Docker client: %w", err)
	}

	// Create the exec options
	execConfig := container.ExecOptions{
		Cmd:          command,
		Tty:          false,
		AttachStdout: true,
		AttachStderr: true,
	}

	// Create an exec instance
	execIDResp, err := cli.ContainerExecCreate(context.Background(), containerID, execConfig)
	if err != nil {
		return fmt.Errorf("error creating exec instance: %w", err)
	}

	// Create exec attach options
	attachOptions := container.ExecAttachOptions{
		Detach: false,
		Tty:    false,
	}

	// Attach to the exec instance
	execResp, err := cli.ContainerExecAttach(context.Background(), execIDResp.ID, attachOptions)
	if err != nil {
		return fmt.Errorf("error attaching to exec instance: %w", err)
	}
	defer execResp.Close()

	// Create a multi-writer to capture both stdout and stderr
	multiWriter := io.MultiWriter(os.Stdout, os.Stderr)

	// Read the output in real-time and print it
	_, err = io.Copy(multiWriter, execResp.Reader) // Copy the output in real-time
	if err != nil {
		return fmt.Errorf("error reading output: %w", err)
	}

	// Create exec start options
	startOptions := container.ExecStartOptions{
		Detach: false,
		Tty:    false,
	}

	// Start the exec instance
	if err := cli.ContainerExecStart(context.Background(), execIDResp.ID, startOptions); err != nil {
		return fmt.Errorf("error starting exec instance: %w", err)
	}

	return nil
}
