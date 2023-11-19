package main

import (
	"fmt"
	"os"
	docker "github.com/fsouza/go-dockerclient"
)	

func buildContainer(client *docker.Client){

		containerConfig := &docker.Config{
			Image: "runnermain",
		}
	
		// Define the host configuration
		hostConfig := &docker.HostConfig{}
	
		// Create the container
		container, err := client.CreateContainer(docker.CreateContainerOptions{
			Config:     containerConfig,
			HostConfig: hostConfig,
		})
	
		if err != nil {
			panic(err)
		}

		// Start the container 
		err = client.StartContainer(container.ID, nil)
		if err != nil {
			panic(err)
		}

		container, err = client.InspectContainerWithOptions(docker.InspectContainerOptions{
			ID: container.ID,
		})

		if err != nil {
			panic(err)
		}

		info := initedContainer{ID: container.ID, ContainerAddr: container.NetworkSettings.IPAddress, ContainerDir: container.GraphDriver.Data["MergedDir"], Running: false}

		t := insertContainer(info)
		fmt.Println(t)
	
}


func uploadCodebase(srcDir string, id string){

	client, err := docker.NewClientFromEnv()
    if err != nil {
        fmt.Printf("Error creating Docker client: %s\n", err)
    }

	// print file name
	fmt.Println("srcDir")

	compDir, err := os.Open("/home/nawaf/Documents/GitHub/runnr/DynamicSpace/test/test.txt")

	if err != nil {
        fmt.Printf("Error compressing: %s\n", err)
    }
	
	// Copy the file from the host file system to the container
	err = client.UploadToContainer(id, docker.UploadToContainerOptions{
		Path:    "/app",
		InputStream:  compDir,
	})
	if err != nil {
		fmt.Printf("Error copying file: %s\n", err)
	}
}


func getDockerInfo(id string) *docker.Container{


	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	container, err := client.InspectContainerWithOptions(docker.InspectContainerOptions{
		ID: id,
	})

	if err != nil {
		panic(err)
	}

	println(container)
	return container
}

func deleteDocker(id string) {

	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	opts := docker.RemoveContainerOptions{
		ID:    id,
		Force: true,
	}
	
	err = client.RemoveContainer(opts)
	if err != nil {
		panic(err)
	}
}