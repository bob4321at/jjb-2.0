package level

import (
	"encoding/json"
	"os"
)

func LoadLevel(path string) Level {
	level := makeLevel(path+"level.png", path+"tileset.png", path+"bg.png")

	file, err := os.ReadFile(path + "waves.json")
	if err != nil {
		panic(err)
	}

	temp_data := Waves{}

	if err := json.Unmarshal(file, &temp_data); err != nil {
		panic(err)
	}

	level.Waves = temp_data
	level.Current_Wave = 0
	level.Spawn_Timer = 50
	level.Origonal_Spawn_Timer = 50

	return level
}

func LoadAllLevels(path string, levels *[]Level) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for folder_index := 1; folder_index < len(dir); folder_index++ {
		folder := dir[folder_index]
		*levels = append(*levels, LoadLevel(path+folder.Name()+"/"))
	}
}
