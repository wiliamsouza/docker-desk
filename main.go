package main

import (
	"fmt"
	"os"

	dockerclient "github.com/fsouza/go-dockerclient"
	"github.com/niemeyer/qml"

	"github.com/wiliamsouza/docker-desk/models"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	qml.Init(nil)
	engine := qml.NewEngine()
	imageModel := &models.ImageModel{}
	endpoint := "unix:///var/run/docker.sock"
	client, err := dockerclient.NewClient(endpoint)
	imgs, _ := client.ListImages(true)
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentId)
		fmt.Println("Repository: ", img.Repository)
		imageModel.Add(img)
	}
	engine.Context().SetVar("imageModel", imageModel)
	component, err := engine.LoadFile("views/docker.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
	window.Show()
	window.Wait()
	return nil
}
