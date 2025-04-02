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
	Update       func(e *PlayerEntity, level_hitbox []utils.HitBox)
	ID           int
}

func (p *Player) NewEntity(pos utils.Vec2, starting_vel utils.Vec2, cooldown float64, lifespan float64, img textures.RenderableTexture, Update func(e *PlayerEntity, level_hitbox []utils.HitBox)) (e *PlayerEntity) {
	entity := PlayerEntity{}

	entity.Img = img

	entity.Pos = pos
	entity.Rotation = 0
	entity.Vel = starting_vel

	entity.Cooldown = cooldown
	entity.Max_Cooldown = cooldown

	entity.Lifespan = lifespan

	entity.Update = Update

	p.Entities = append(p.Entities, entity)

	return &p.Entities[len(p.Entities)-1]
}

func (e *PlayerEntity) SetID(num int) {
	e.ID = num
}
