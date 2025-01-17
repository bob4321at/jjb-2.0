package main

import (
	"encoding/json"
	"os"
)

func loadLevel(path string) Level {
	l := makeLevel(path+"level.png", path+"tileset.png", path+"bg.png")

	f, err := os.ReadFile(path + "waves.json")
	if err != nil {
		panic(err)
	}

	temp_data := Waves{}

	json.Unmarshal(f, &temp_data)

	l.waves = temp_data

	return l
}
