package main

import (
	"encoding/json"
	"fmt"
	"log"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/gofiber/fiber/v2"
)

type Request struct {
    ContainerID string `json:"containerId"`
    Message  string    `json:"msg"`
}

func handleComms(c *fiber.Ctx) string{
    // Read message from client
    body := c.Body()

    var req Request
    if err := json.Unmarshal(body, &req); err != nil {
        fmt.Println("Error: " + err.Error())
        return "err"
    }

    // Print message from client
	fmt.Println(req.ContainerID)

    //GET FULL CONTAINER ID && function
    container := getDockerInfo(req.ContainerID)
    function := getFunctionContainer(container.ID)

    fmt.Println(function.ID)
    //update function
    updateFunctions(function.ID, "null", false)

    //delete used contianer
    deleteContainer(container.ID)

    client, err := docker.NewClientFromEnv()
    if err != nil {
        log.Fatal(err)
    }

    buildContainer(client)

    // delete container 
	return "ok"
}