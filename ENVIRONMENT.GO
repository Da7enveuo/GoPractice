package main

import (
	"os"
	"fmt"
)

// very basic iteration of enviornment variables to find the twitter api key
func main(){
	env := os.Environ()
 
  for name, envs := range env {
 	if name == "TWITTER_API_KEY"{
    	fmt.Println(envs)
    }
}
