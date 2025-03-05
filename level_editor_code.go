package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func makeLevel(path string, tileset_path string, background_path string) (l Level) {
	l.enemies = []Enemy{}
	l.background = newBackground(Vec2{0, 1024 * 1.8}, 100, newTexture(background_path))

	timg, _, err := ebitenutil.NewImageFromFile(tileset_path)
	if err != nil {
		panic(err)
	}

	l.tileset = map[int]*ebiten.Image{
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

	l.tileset[1] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(0, 0, 32, 32)))
	l.tileset[2] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(32, 0, 64, 32)))
	l.tileset[3] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(64, 0, 96, 32)))
	l.tileset[4] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(0, 32, 32, 64)))
	l.tileset[5] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(32, 32, 64, 64)))
	l.tileset[6] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(64, 32, 96, 64)))
	l.tileset[7] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(0, 64, 32, 96)))
	l.tileset[8] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(32, 64, 64, 96)))
	l.tileset[9] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(64, 64, 96, 96)))
	l.tileset[10] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(96, 0, 128, 32)))
	l.tileset[11] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(96, 32, 128, 64)))
	l.tileset[12] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(96, 64, 128, 96)))
	l.tileset[13] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(0, 96, 32, 128)))
	l.tileset[14] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(32, 96, 64, 128)))
	l.tileset[15] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(64, 96, 96, 128)))
	l.tileset[16] = ebiten.NewImageFromImage(timg.SubImage(image.Rect(96, 96, 128, 128)))

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
				l.player_spawn.x = float64(x * 32)
				l.player_spawn.y = float64(y * 32)
				l.background.start.y = float64(y*32 - 550)
				row = append(row, 0)
			} else if map_img.At(x, y) == blue {
				l.spawn_points = append(l.spawn_points, Vec2{float64(x)*32 - 16, float64(y)*32 - 16})
				row = append(row, 0)
			} else {
				row = append(row, 0)
			}
		}
		l.tile_map = append(l.tile_map, row)
		row = []uint8{}
	}

	new_map := l.tile_map

	for y := 0; y < len(l.tile_map); y++ {
		for x := 0; x < len(l.tile_map[y]); x++ {
			t := &new_map[y][x]
			if l.tile_map[y][x] == 5 {
				// top left
				if x-1 > 0 {
					if l.tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] == 0 {
								if l.tile_map[y+1][x] != 0 {
									if x+1 < len(l.tile_map[y]) {
										if l.tile_map[y][x+1] != 0 {
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
					if l.tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] == 0 {
								if l.tile_map[y+1][x] != 0 {
									if x+1 < len(l.tile_map[y]) {
										if l.tile_map[y][x+1] != 0 {
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
					if l.tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] == 0 {
								if l.tile_map[y+1][x] != 0 {
									if x+1 < len(l.tile_map[y]) {
										if l.tile_map[y][x+1] == 0 {
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
					if l.tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] != 0 {
								if l.tile_map[y+1][x] != 0 {
									if x+1 < len(l.tile_map[y]) {
										if l.tile_map[y][x+1] != 0 {
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
					if l.tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] != 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] != 0 {
										if x+1 < len(l.tile_map[y]) {
											if l.tile_map[y][x+1] != 0 {
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
					if l.tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] != 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] != 0 {
										if x+1 < len(l.tile_map[y]) {
											if l.tile_map[y][x+1] == 0 {
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
					if l.tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] != 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if l.tile_map[y][x+1] != 0 {
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
					if l.tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] != 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] == 0 {
										if x+1 < len(l.tile_map) {
											if l.tile_map[y][x+1] != 0 {
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
					if l.tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] != 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] == 0 {
										if x+1 < len(l.tile_map) {
											if l.tile_map[y][x+1] == 0 {
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
					if l.tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] == 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] != 0 {
										if x-1 > 0 {
											if x+1 < len(l.tile_map[y]) {
												if l.tile_map[y][x+1] == 0 {
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
					if l.tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] != 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] != 0 {
										if x+1 < len(l.tile_map) {
											if l.tile_map[y][x+1] == 0 {
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
					if l.tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] != 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if l.tile_map[y][x+1] == 0 {
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
					if l.tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] == 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if l.tile_map[y][x+1] != 0 {
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
					if l.tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] == 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if l.tile_map[y][x+1] != 0 {
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
					if l.tile_map[y][x-1] != 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] == 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if l.tile_map[y][x+1] == 0 {
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
					if l.tile_map[y][x-1] == 0 {
						if y-1 > 0 {
							if l.tile_map[y-1][x] == 0 {
								if y+1 < len(l.tile_map) {
									if l.tile_map[y+1][x] == 0 {
										if x-1 > 0 {
											if l.tile_map[y][x+1] == 0 {
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

	l.tile_map = new_map

	for y := 0; y < len(l.tile_map); y++ {
		for x := 0; x < len(l.tile_map); x++ {
			if l.tile_map[y][x] != 0 {
				l.tiles = append(l.tiles, Tile{Vec2{float64(x) * 32, float64(y) * 32}, int(l.tile_map[y][x])})
			}
		}
	}

	l.generated = true

	current_level = &l

	initPlayer()

	return l
}
