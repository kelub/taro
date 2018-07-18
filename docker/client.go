package docker

import (
	docker_client "github.com/fsouza/go-dockerclient"
	"fmt"
	"strings"
)

const (
	defaultEndpoint = "unix:///var/run/docker.sock"
)

type DockerClienter interface {
	ListImages(docker_client.ListImagesOptions)([]docker_client.APIImages, error)
	CreateContainer(opts docker_client.CreateContainerOptions) (*docker_client.Container, error)
	//CreateContainer()()
}

type DockerManager struct {
	Client DockerClienter
	EndPoint   string
}

func NewDockerManager(endpoint string)(*DockerManager,error){
	if len(strings.TrimSpace(endpoint)) == 0 {
		endpoint = defaultEndpoint
	}
	client, err := docker_client.NewClient(endpoint)
	if err != nil {
		return nil, fmt.Errorf("new Docker client with error %s", err.Error())
	}
	return &DockerManager{
		Client: client,
		EndPoint: endpoint,
	},nil
}

func (dm *DockerManager)ListImages(opt docker_client.ListImagesOptions)([]docker_client.APIImages, error){
	docker_lists, err:= dm.Client.ListImages(opt)
	if err != nil{
		return nil, fmt.Errorf("get Docker images list with error %s", err.Error())
	}
	return docker_lists, nil
}
//func (dm *DockerManager)CreateContainer()(){
//
//}
//func client()(){
//	endpoint := "unix:///var/run/docker.sock"
//	client, err := docker_client.NewClient(endpoint)
//	imgs, err := client.ListImages(docker_client.ListImagesOptions{All: false})
//}