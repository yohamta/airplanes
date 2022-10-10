package archetypes

import (
	"math"

	"github.com/yohamta/donburi"

	"github.com/m110/airplanes/assets"
	"github.com/m110/airplanes/component"
	"github.com/m110/airplanes/engine"
)

func NewBullet(w donburi.World, player *component.PlayerData, position engine.Vector) {
	width := float64(assets.LaserSingle.Bounds().Dy())

	if player.WeaponLevel == component.WeaponLevelSingle ||
		player.WeaponLevel == component.WeaponLevelSingleFast {
		bullet := newBullet(w)
		donburi.SetValue(bullet, component.Transform, component.TransformData{
			Position: engine.Vector{
				X: position.X,
				Y: position.Y - width,
			},
		})
	}

	if player.WeaponLevel == component.WeaponLevelDouble ||
		player.WeaponLevel == component.WeaponLevelDoubleFast ||
		player.WeaponLevel == component.WeaponLevelDiagonal ||
		player.WeaponLevel == component.WeaponLevelDoubleDiagonal {
		bullet1 := newBullet(w)
		bullet2 := newBullet(w)
		donburi.SetValue(bullet1, component.Transform, component.TransformData{
			Position: engine.Vector{
				X: position.X - width/2,
				Y: position.Y - width/2,
			},
		})
		donburi.SetValue(bullet2, component.Transform, component.TransformData{
			Position: engine.Vector{
				X: position.X + width/2,
				Y: position.Y - width/2,
			},
		})
	}

	if player.WeaponLevel == component.WeaponLevelDiagonal ||
		player.WeaponLevel == component.WeaponLevelDoubleDiagonal {
		bulletNW := newBullet(w)
		bulletNE := newBullet(w)
		donburi.SetValue(bulletNW, component.Transform, component.TransformData{
			Position: engine.Vector{
				X: position.X - width,
				Y: position.Y - width,
			},
		})
		component.GetTransform(bulletNW).Rotation = -30
		radians := engine.ToRadians(-30 - 90)
		component.GetVelocity(bulletNW).X = 10 * math.Cos(radians)
		component.GetVelocity(bulletNW).Y = 10 * math.Sin(radians)
		donburi.SetValue(bulletNE, component.Transform, component.TransformData{
			Position: engine.Vector{
				X: position.X + width,
				Y: position.Y - width,
			},
		})
		radians = engine.ToRadians(30 - 90)
		component.GetVelocity(bulletNE).X = 10 * math.Cos(radians)
		component.GetVelocity(bulletNE).Y = 10 * math.Sin(radians)
		component.GetTransform(bulletNE).Rotation = 30
	}

	if player.WeaponLevel == component.WeaponLevelDoubleDiagonal {
		bulletNW := newBullet(w)
		bulletNE := newBullet(w)
		donburi.SetValue(bulletNW, component.Transform, component.TransformData{
			Position: engine.Vector{
				X: position.X - width*1.1,
				Y: position.Y,
			},
		})
		component.GetTransform(bulletNW).Rotation = -30
		radians := engine.ToRadians(-30 - 90)
		component.GetVelocity(bulletNW).X = 10 * math.Cos(radians)
		component.GetVelocity(bulletNW).Y = 10 * math.Sin(radians)
		donburi.SetValue(bulletNE, component.Transform, component.TransformData{
			Position: engine.Vector{
				X: position.X + width*1.1,
				Y: position.Y,
			},
		})
		radians = engine.ToRadians(30 - 90)
		component.GetVelocity(bulletNE).X = 10 * math.Cos(radians)
		component.GetVelocity(bulletNE).Y = 10 * math.Sin(radians)
		component.GetTransform(bulletNE).Rotation = 30
	}
}

func newBullet(w donburi.World) *donburi.Entry {
	bullet := w.Entry(
		w.Create(
			component.Velocity,
			component.Transform,
			component.Sprite,
			component.Despawnable,
			component.Collider,
		),
	)

	image := assets.LaserSingle

	component.GetVelocity(bullet).Y = -10

	donburi.SetValue(bullet, component.Sprite, component.SpriteData{
		Image:            image,
		Layer:            component.SpriteLayerAirUnits,
		Pivot:            component.SpritePivotCenter,
		OriginalRotation: -90,
	})

	width, height := image.Size()

	donburi.SetValue(bullet, component.Collider, component.ColliderData{
		Width:  float64(width),
		Height: float64(height),
		Layer:  component.CollisionLayerBullets,
	})

	return bullet
}
