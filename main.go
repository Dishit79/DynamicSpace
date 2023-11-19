package main

import (
	"fmt"
    "github.com/gofiber/fiber/v2"
	docker "github.com/fsouza/go-dockerclient"
)

type initedContainer struct {
	ID string `bson:"id"`
	ContainerAddr string `bson:"containerAddr"`
	ContainerDir string  `bson:"containerDir"`
	Running bool `bson:"running"`
} 

//starts the containers with given amount
func startContainers(client *docker.Client, number int){
	//Creates fresh containers 
	for i := 0; i < number; i++ {
		buildContainer(client)
	}	
}

//initlizies containers 
func initContainters(){
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	startContainers(client, 1)

}

func main() {
    app := fiber.New()
	initContainters()

	endpoints := []string{"test2", "lyrics", "lyrics1", "basin"}

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Post("/app/container/kill", func (c *fiber.Ctx) error {
		t := handleComms(c)
        return c.SendString("killing container" + t)
    })

	app.Get("/api/:host/*", func(c *fiber.Ctx) error {
		calledEndpoint := c.Params("host")
		calledPath := c.Params("*")
		queryParams := c.Request().URI().QueryString()
		    
		endpointCalled := false
		for _, x := range endpoints {
			if x == calledEndpoint {
				endpointCalled = true
				break
			}
   		} 

		fmt.Println(calledPath)

		if endpointCalled {
			add := invokerHandler(calledEndpoint)
			handleForward(add, calledPath, queryParams, c)
			return(nil)
		} else {
			msg := ("Endpoint isn't real")
			return c.SendString(msg) 
		}
    })


    // log.Fatal(app.Listen(":3000"))
	app.Listen(":3000")
	fmt.Println("webserver started")
}