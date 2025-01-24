package main

type Vec2 struct {
	x, y float64
}

func collide(pos1, size1, pos2, size2 Vec2) bool {
	if pos1.x < pos2.x+size2.x && pos1.x+size1.x > pos2.x {
		if pos1.y < pos2.y+size2.y && pos1.y+size1.y > pos2.y {
			return true
		}
	}
	return false
}

func removeEnemy(index_to_remove int, slice []Enemy) []Enemy {
	return append(slice[:index_to_remove], slice[index_to_remove+1:]...)
}

func removePlayerEntity(index_to_remove int, slice []PlayerEntity) []PlayerEntity {
	return append(slice[:index_to_remove], slice[index_to_remove+1:]...)
}

func removeProjectile(index_to_remove int, slice []Projectile) []Projectile {
	return append(slice[:index_to_remove], slice[index_to_remove+1:]...)
}
