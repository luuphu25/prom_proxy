package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main(){
	var services []string
	files, err := ioutil.ReadDir("/etc/systemd/system/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "exporter.service"){
			services = append(services, file.Name())
		}
	}
	fmt.Println(services)
}
