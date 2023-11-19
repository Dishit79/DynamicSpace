package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "time"
    "errors"

    "github.com/gofiber/fiber/v2"
)


func requestLoop(u url.URL) (*http.Response, error){

    for i := 0; i < 100; i++ {
        req, err := http.NewRequest("GET", u.String(), nil)
        if err == nil {
            fmt.Println("Response received")
        }
        resp, err := http.DefaultClient.Do(req)
        if err == nil {
            return resp, nil
        }

        fmt.Println("Request failed, retrying in .30 second")
        time.Sleep(30 * time.Millisecond)
    }
    return nil, errors.New("failed to call")
}


func handleForward(internalEndpoint string, path string, queryString []byte,c *fiber.Ctx) {
    // Create the request to the internal endpoint

    startTime := time.Now()
    queryParams, _ := url.ParseQuery(string(queryString))
    u := url.URL{
        Scheme: "http",
        Host:   internalEndpoint + ":8080",
        Path: path ,
        RawQuery: queryParams.Encode(),
    }


    resp, err := requestLoop(u)
    if err != nil {
        fmt.Println(err)
        c.SendStatus(fiber.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        fmt.Println(err)
        c.SendStatus(fiber.StatusBadRequest)
        return
    }

    timeTook := startTime.Sub(time.Now())
    fmt.Println(timeTook)
    // Forward the response back to the client
    responseBody, _ := ioutil.ReadAll(resp.Body)
    c.Set("content-type", resp.Header.Get("Content-Type"))
    c.Status(resp.StatusCode).Send(responseBody)
}