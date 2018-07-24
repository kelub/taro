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


func TestCreateContainer(test *testing.T){
	var entrypoint = []string{"/sbin/init"}
	var cmd = []string{}
	fmt.Println("CreateContainer##############################################")
	config := docker_client.Config{
		AttachStdout: true,
		AttachStdin: true,
		Entrypoint: entrypoint,
		Cmd: cmd,
		Image: "busybox:latest",
	}
	client ,_ := NewDockerManager("")
	ccops := docker_client.CreateContainerOptions{Name: "testCreateContainer", Config: &config}
	create, err := client.CreateContainer(ccops)
	if err != nil{
		fmt.Println(err)
		//test.Errorf("error CreateContainer  %s",err)
	}else{
		fmt.Println(create.ID)
	}
}

func TestStartContainer(test *testing.T){
	id := "baec842f0da2e132a7c3f09979eb802f35409f1ab397b9895bb3631f90bb7d69"
	client ,_ := NewDockerManager("")
	err := client.StartContainer(id,&docker_client.HostConfig{})
	if err != nil{
		test.Errorf("error StartContainer  %s",err)

	}
}