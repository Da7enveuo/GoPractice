package main

import (
	"os"
	"fmt"
)


func main(){
	env := os.Environ()
 
  for name, envs := range env {
 	if name == "TWITTER_API_KEY"{
    	fmt.Println(envs)
    }
}