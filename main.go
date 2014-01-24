package main

import (
	"fmt"
	"os"

	dockerclient "github.com/fsouza/go-dockerclient"
	"github.com/niemeyer/qml"
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

	images := &Images{}
	engine.Context().SetVar("images", images)
	component, err := engine.LoadFile("docker.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
	window.Show()

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
		images.Add(img)
	}

	window.Wait()
	return nil
}

type Images struct {
	list []dockerclient.APIImages
	Len  int
}

func (images *Images) Add(i dockerclient.APIImages) {
	images.list = append(images.list, i)
	images.Len = len(images.list)
	qml.Changed(images, &images.Len)
}

func (images *Images) Image(index int) string {
	img := images.list[index]
	return img.ID
}
