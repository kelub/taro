package docker

import (
	"testing"
	docker_client "github.com/fsouza/go-dockerclient"

	"fmt"
)

func TestNewDockerManager(test *testing.T){
	client ,_ := NewDockerManager("")
	ListImagesOptions := docker_client.ListImagesOptions{All: false}
	listimages, err := client.ListImages(ListImagesOptions)
	if err != nil{
		fmt.Println(err)
		test.Errorf("get listimages  %s",err)
	}
	fmt.Println(listimages)
}
