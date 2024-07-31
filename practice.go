package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script_name.go number")
		return
	}

	s := os.Args[1]

	datas := strings.Split(s, "/")
	id := datas[len(datas)-2]
	number := datas[len(datas)-1]
	_, err := os.Stat("practice/" + id)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println(id + " does not exist, creating directory...")
		err = os.Mkdir("practice/"+id, 0777)
		if err != nil {
			fmt.Println("error creating directory:" + err.Error())
			os.Exit(1)
		}
	}

	err = os.MkdirAll("practice/"+id+"/"+number, 0777)
	if err != nil {
		fmt.Println("error creating directory:" + err.Error())
		os.Exit(1)
	}

	_, err = os.Create(strings.Join([]string{"practice", id, number, number + ".go"}, "/"))
	if err != nil {
		fmt.Println("error creating file:" + err.Error())
		os.Exit(1)
	}

	file, err := os.ReadFile("template.go")
	if err != nil {
		fmt.Println("error reading file:" + err.Error())
		os.Exit(1)
	}

	err = os.WriteFile(strings.Join([]string{"practice", id, number, number + ".go"}, "/"), file, 0666)
	if err != nil {
		fmt.Println("error creating file:" + err.Error())
		os.Exit(1)
	}

	fmt.Println("complete!")
}
