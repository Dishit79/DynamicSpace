package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createClient() mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	// Set up a connection to the MongoDB server
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return *client

}

func instertFunction() {
	
	client := createClient()
	
	user := Functions{
	Endpoint: "test2",
	Dir: "/home/nawaf/Documents/GitHub/runnr/DynamicSpace/test2.tar.xz",
	}

	// Insert a document into the "users" collection
	collection := client.Database("dynamicspace").Collection("functions")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a document into the users collection")
}

func insertContainer(data initedContainer) string {

	client := createClient()

	// Insert a document into the "users" collection
	collection := client.Database("dynamicspace").Collection("containers")
	_, err := collection.InsertOne(context.TODO(), data)

	if err != nil {
		log.Fatal(err)
	}

	return ("Inserted a container into the users collection")

}

func getContainingContainer(id string) initedContainer{

	client := createClient()

	// Retrieve a single document from the "users" collection
	collection := client.Database("dynamicspace").Collection("containers")
	var result initedContainer
	err := collection.FindOne(context.TODO(), bson.M{ "id": id }).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func updatedUsedFunctions(functionId string, containerid string, status bool) Functions{

	client := createClient()

	// Retrieve a single document from the "users" collection
	collection := client.Database("dynamicspace").Collection("functions")
	var result Functions
	err := collection.FindOne(context.TODO(), bson.M{ "id": functionId }).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	update := bson.M{"$set": bson.M{"status": status, "containerId": containerid}}
	filter := bson.M{"id" : result.ID}
	test, err := collection.UpdateOne(context.TODO(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(test)
	return result
}

func getContainerRunning(running bool) initedContainer{

	client := createClient()

	// Retrieve a single document from the "users" collection
	collection := client.Database("dynamicspace").Collection("containers")
	var result initedContainer
	err := collection.FindOne(context.TODO(), bson.M{ "running": running }).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	update := bson.M{"$set": bson.M{"running": true}}
	filter := bson.M{"id" : result.ID}
	test, err := collection.UpdateOne(context.TODO(),filter,update)
	fmt.Println(test)
	return result
}

func getFunctionEndpoint(endpoint string) Functions{

	client := createClient()

	fmt.Println(endpoint)
	// Retrieve a single document from the "users" collection
	collection := client.Database("dynamicspace").Collection("functions")
	var result Functions
	err := collection.FindOne(context.TODO(), bson.M{"endpoint": endpoint}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	
	return result
}

func deleteContainer(id string) {

	client := createClient()

	collection := client.Database("dynamicspace").Collection("containers")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		panic(err)
	}

	deleteDocker(id)
}
 

func getFunctionContainer(id string) Functions{

	client := createClient()

	// Retrieve a single document from the "users" collection
	collection := client.Database("dynamicspace").Collection("functions")
	var result Functions
	err := collection.FindOne(context.TODO(), bson.M{"containerId": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	
	return result
}
