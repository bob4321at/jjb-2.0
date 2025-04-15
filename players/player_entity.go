package players

import (
	"jjb/textures"
	"jjb/utils"
)

type PlayerEntity struct {
	Pos          utils.Vec2
	Rotation     float64
	Vel          utils.Vec2
	Cooldown     float64
	Max_Cooldown float64
	Lifespan     float64
	Img          textures.RenderableTexture
	Dir          bool
	Update       func(entity *PlayerEntity, level_hitbox []utils.HitBox)
	ID           int
}

func (player *Player) NewEntity(pos utils.Vec2, starting_vel utils.Vec2, cooldown float64, lifespan float64, img textures.RenderableTexture, Update func(entity *PlayerEntity, level_hitbox []utils.HitBox)) *PlayerEntity {
	entity := PlayerEntity{}

	entity.Img = img

	entity.Pos = pos
	entity.Rotation = 0
	entity.Vel = starting_vel

	entity.Cooldown = cooldown
	entity.Max_Cooldown = cooldown

	entity.Lifespan = lifespan

	entity.Update = Update

	player.Entities = append(player.Entities, entity)

	return &player.Entities[len(player.Entities)-1]
}

func (entity *PlayerEntity) SetID(number int) {
	entity.ID = number
}
