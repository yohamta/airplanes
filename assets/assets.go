package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed ships/ship_0011.png
	shipYellowSmallData []byte
	//go:embed tiles/tile_0000.png
	laserSingleData []byte

	ShipYellowSmall *ebiten.Image
	LaserSingle     *ebiten.Image
)

func LoadAssets() {
	ShipYellowSmall = mustNewEbitenImage(shipYellowSmallData)
	LaserSingle = mustNewEbitenImage(laserSingleData)
}

func mustNewEbitenImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
