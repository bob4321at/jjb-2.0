package level

import (
	"image"
	"image/color"
	"jjb/enemyai"
	"jjb/players"
	"jjb/textures"
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func makeLevel(path string, tileset_path string, background_path string) (level Level) {
	level.Enemies = []enemyai.Enemy{}
	level.Background = newBackground(utils.Vec2{X: 0, Y: 1024 * 1.8}, 100, textures.NewTexture(background_path, ""))

	temporary_image, _, err := ebitenutil.NewImageFromFile(tileset_path)
	if err != nil {
		panic(err)
	}

	level.Tileset = map[int]*ebiten.Image{
		0:  ebiten.NewImage(32, 32),
		1:  ebiten.NewImage(32, 32),
		2:  ebiten.NewImage(32, 32),
		3:  ebiten.NewImage(32, 32),
		4:  ebiten.NewImage(32, 32),
		5:  ebiten.NewImage(32, 32),
		6:  ebiten.NewImage(32, 32),
		7:  ebiten.NewImage(32, 32),
		8:  ebiten.NewImage(32, 32),
		9:  ebiten.NewImage(32, 32),
		10: ebiten.NewImage(32, 32),
		11: ebiten.NewImage(32, 32),
		12: ebiten.NewImage(32, 32),
		13: ebiten.NewImage(32, 32),
		14: ebiten.NewImage(32, 32),
		15: ebiten.NewImage(32, 32),
		16: ebiten.NewImage(32, 32),
	}

	level.Tileset[1] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(0, 0, 32, 32)))
	level.Tileset[2] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(32, 0, 64, 32)))
	level.Tileset[3] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(64, 0, 96, 32)))
	level.Tileset[4] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(0, 32, 32, 64)))
	level.Tileset[5] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(32, 32, 64, 64)))
	level.Tileset[6] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(64, 32, 96, 64)))
	level.Tileset[7] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(0, 64, 32, 96)))
	level.Tileset[8] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(32, 64, 64, 96)))
	level.Tileset[9] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(64, 64, 96, 96)))
	level.Tileset[10] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(96, 0, 128, 32)))
	level.Tileset[11] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(96, 32, 128, 64)))
	level.Tileset[12] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(96, 64, 128, 96)))
	level.Tileset[13] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(0, 96, 32, 128)))
	level.Tileset[14] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(32, 96, 64, 128)))
	level.Tileset[15] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(64, 96, 96, 128)))
	level.Tileset[16] = ebiten.NewImageFromImage(temporary_image.SubImage(image.Rect(96, 96, 128, 128)))

	map_img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}

	row := []uint8{}
	for y := 0; y < map_img.Bounds().Dy(); y++ {
		for x := 0; x < map_img.Bounds().Dx(); x++ {
			white := color.RGBA{255, 255, 255, 255}
			yellow := color.RGBA{255, 255, 0, 255}
			blue := color.RGBA{0, 0, 255, 255}
			if map_img.At(x, y) == white {
				row = append(row, 5)
			} else if map_img.At(x, y) == yellow {
				level.Player_Spawn.X = float64(x * 32)
				level.Player_Spawn.Y = float64(y * 32)
				level.Background.start.Y = float64(y*32 - 550)
				row = append(row, 0)
			} else if map_img.At(x, y) == blue {
				level.Spawn_Points = append(level.Spawn_Points, utils.Vec2{X: float64(x)*32 - 16, Y: float64(y)*32 - 16})
				row = append(row, 0)
			} else {
				row = append(row, 0)
			}
		}
		level.Tile_map = append(level.Tile_map, row)
		row = []uint8{}
	}

	new_map := level.Tile_map

	for y := 0; y < len(level.Tile_map); y++ {
		for x := 0; x < len(level.Tile_map[y]); x++ {
			t := &new_map[y][x]
			if level.Tile_map[y][x] == 5 {
				// top left
				if x-1 > 0 {
					if level.Tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] == 0 {
								if level.Tile_map[y+1][x] != 0 {
									if x+1 < len(level.Tile_map[y]) {
										if level.Tile_map[y][x+1] != 0 {
											*t = 1
										}
									}
								}
							}
						}
					}
				}
				//top middle
				if x-1 > 0 {
					if level.Tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] == 0 {
								if level.Tile_map[y+1][x] != 0 {
									if x+1 < len(level.Tile_map[y]) {
										if level.Tile_map[y][x+1] != 0 {
											*t = 2
										}
									}
								}
							}
						}
					}
				}
				//top right
				if x-1 > 0 {
					if level.Tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] == 0 {
								if level.Tile_map[y+1][x] != 0 {
									if x+1 < len(level.Tile_map[y]) {
										if level.Tile_map[y][x+1] == 0 {
											*t = 3
										}
									}
								}
							}
						}
					}
				}
				//middle left
				if x-1 > 0 {
					if level.Tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] != 0 {
								if level.Tile_map[y+1][x] != 0 {
									if x+1 < len(level.Tile_map[y]) {
										if level.Tile_map[y][x+1] != 0 {
											*t = 4
										}
									}
								}
							}
						}
					}
				}
				//middle middle
				if x-1 > 0 {
					if level.Tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] != 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] != 0 {
										if x+1 < len(level.Tile_map[y]) {
											if level.Tile_map[y][x+1] != 0 {
												*t = 5
											}
										}
									}
								}
							}
						}
					}
				}
				//middle right
				if x-1 > 0 {
					if level.Tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] != 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] != 0 {
										if x+1 < len(level.Tile_map[y]) {
											if level.Tile_map[y][x+1] == 0 {
												*t = 6
											}
										}
									}
								}
							}
						}
					}
				}
				//bottom left
				if x-1 > 0 {
					if level.Tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] != 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if level.Tile_map[y][x+1] != 0 {
												*t = 7
											}
										}
									}
								}
							}
						}
					}
				}
				//bottom middle
				if x-1 > 0 {
					if level.Tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] != 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] == 0 {
										if x+1 < len(level.Tile_map) {
											if level.Tile_map[y][x+1] != 0 {
												*t = 8
											}
										}
									}
								}
							}
						}
					}
				}
				//bottom right
				if x-1 > 0 {
					if level.Tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] != 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] == 0 {
										if x+1 < len(level.Tile_map) {
											if level.Tile_map[y][x+1] == 0 {
												*t = 9
											}
										}
									}
								}
							}
						}
					}
				}
				//pole top
				if x-1 > 0 {
					if level.Tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] == 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] != 0 {
										if x-1 > 0 {
											if x+1 < len(level.Tile_map[y]) {
												if level.Tile_map[y][x+1] == 0 {
													*t = 10
												}
											}
										}
									}
								}
							}
						}
					}
				}
				//pole middle
				if x-1 > 0 {
					if level.Tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] != 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] != 0 {
										if x+1 < len(level.Tile_map) {
											if level.Tile_map[y][x+1] == 0 {
												*t = 11
											}
										}
									}
								}
							}
						}
					}
				}
				//pole bottom
				if x-1 > 0 {
					if level.Tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] != 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if level.Tile_map[y][x+1] == 0 {
												*t = 12
											}
										}
									}
								}
							}
						}
					}
				}
				//pipe left
				if x-1 > 0 {
					if level.Tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] == 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if level.Tile_map[y][x+1] != 0 {
												*t = 13
											}
										}
									}
								}
							}
						}
					}
				}
				//pipe middle
				if x-1 > 0 {
					if level.Tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] == 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if level.Tile_map[y][x+1] != 0 {
												*t = 14
											}
										}
									}
								}
							}
						}
					}
				}
				//pipe right
				if x-1 > 0 {
					if level.Tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] == 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if level.Tile_map[y][x+1] == 0 {
												*t = 15
											}
										}
									}
								}
							}
						}
					}
				}
				//block
				if x-1 > 0 {
					if level.Tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if level.Tile_map[y-1][x] == 0 {
								if y+1 < len(level.Tile_map) {
									if level.Tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if level.Tile_map[y][x+1] == 0 {
												*t = 16
											}
										}
									}
								}
							}
						}
					}
				}
				//
			}
		}
	}

	level.Tile_map = new_map

	for y := 0; y < len(level.Tile_map); y++ {
		for x := 0; x < len(level.Tile_map); x++ {
			if level.Tile_map[y][x] != 0 {
				level.Tiles = append(level.Tiles, Tile{utils.Vec2{X: float64(x) * 32, Y: float64(y) * 32}, int(level.Tile_map[y][x])})
			}
		}
	}

	for ti := 0; ti < len(level.Tiles); ti++ {
		t := &level.Tiles[ti]
		level.HitBox = append(level.HitBox, utils.HitBox{X: t.Pos.X, Y: t.Pos.Y, W: 32, H: 32})
	}
	level.Gnerated = true

	Current_Level = &level

	players.InitPlayer(Current_Level.Player_Spawn)

	return level
}
