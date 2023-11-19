package main

import (
	"fmt"
	"time"
)


type Functions struct {
	ID string `bson:"id"`
  	Endpoint string `bson:"endpoint"`
  	Dir	 string  `bson:"dir"`
  	Status bool `bson:"status"`
  	Time string `bson:"time"`
	ContainerID string `bson:"containerId"`
}


func updateFunctions(functionId string, containerId string, status bool) Functions {

	updatedFunction := updatedUsedFunctions(functionId, containerId, status)
	// Check lifecycle
	return updatedFunction
	
}


func invokerHandler(endpoint string) string{
	startTime := time.Now()

	functions := getFunctionEndpoint(endpoint)
	
	// Gets if the function is already running or not
	if functions.Status {
		usedContainer := getContainingContainer(functions.ContainerID)
		return usedContainer.ContainerAddr
	}
	
	emptyContaier := getContainerRunning(false)
	updatedFunction := updateFunctions(functions.ID, emptyContaier.ID, true)

	
	uploadCodebase(updatedFunction.Dir, emptyContaier.ID)

	timeTook := startTime.Sub(time.Now())
    fmt.Println("upload time: ", timeTook)

	return (emptyContaier.ContainerAddr)	
	// todo
	// WRITE PAPER UP UNTIL THIS POINT AND THEN DESIGN THE INVOKER HANDELER
}