package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

func makeLevel(path string) (l Level) {
	l.enemies = []Enemy{}

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
									if l.tile_map[y][x+1] != 0 {
										*t = 1
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
				l.tiles = append(l.tiles, TileF{Vec2{float64(x) * 32, float64(y) * 32}, int(l.tile_map[y][x])})
			}
		}
	}

	l.generated = true

	current_level = &l

	initPlayer()

	return l
}
