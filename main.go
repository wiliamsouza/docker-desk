package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/niemeyer/qml"
	"os"
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
	client, err := docker.NewClient(endpoint)
	imgs, _ := client.ListImages(true)
	fmt.Println(imgs)
	images.Add(imgs[0])
	images.Add(imgs[1])

	window.Wait()
	return nil
}

type Images struct {
	list []docker.APIImages
	Len  int
}

func (images *Images) Add(i docker.APIImages) {
	images.list = append(images.list, i)
	images.Len = len(images.list)
	qml.Changed(images, &images.Len)
}

func (images *Images) Image(index int) string {
	//img := images.list[index]
	return "ubuntu" //img.Repository
}
