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

	if err := json.Unmarshal(f, &temp_data); err != nil {
		panic(err)
	}

	l.waves = temp_data
	l.current_wave = 0
	l.spawn_timer = 50
	l.origonal_spawn_timer = 50

	return l
}

func loadAllLevels(path string) (levels []Level) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for folder_index := 0; folder_index < len(dir); folder_index++ {
		folder := dir[folder_index]
		levels = append(levels, loadLevel(path+folder.Name()+"/"))
	}

	return levels
}
