package TileType

import "testing"

func TestWinonlasttile(t *testing.T) {
	var k TileType
	k.lasttile.Set(false)
	k.winonlasttile()
}
