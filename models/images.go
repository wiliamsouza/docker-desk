package models

import (
	dockerclient "github.com/fsouza/go-dockerclient"

	"github.com/niemeyer/qml"
)

// It works like QtCore.QAbstractListModel
type ImageModel struct {
	images []dockerclient.APIImages
}

func (model *ImageModel) Len() int {
	return len(model.images)
}

func (model *ImageModel) Data(index int) dockerclient.APIImages {
	return model.images[index]
}

func (model *ImageModel) Add(image dockerclient.APIImages) {
	model.images = append(model.images, image)
	qml.Changed(model, &model.images)
}
