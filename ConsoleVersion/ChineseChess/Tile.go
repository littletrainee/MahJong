package ChineseChess

import "github.com/littletrainee/GetSet"

var Tile_key []string = []string{
	"1b", "2b", "3b", "4b", "5b", "6b", "7b",
	"1r", "2r", "3r", "4r", "5r", "6r", "7r",
}

var Tile_value []string = []string{
	"將", "士", "象", "車", "馬", "砲", "卒",
	"帥", "仕", "相", "俥", "傌", "炮", "兵",
}

var Tilemap GetSet.Type[map[string]string]
