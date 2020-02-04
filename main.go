package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//wall_path := "/opt/wallpapers"
	wall_path := os.Args[1]
	var wallpapers = make([]string, 0, 1164)
	//var wallpapers []string
	err := filepath.Walk(wall_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("something wrong at &q: &v\n", path, err)
			return err
		}
		if info.IsDir() == false {
			wallpapers = append(wallpapers, "\""+path+"\"")
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	rnd_src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(rnd_src)
	fmt.Println(wallpapers[rnd.Intn(len(wallpapers))])
}
