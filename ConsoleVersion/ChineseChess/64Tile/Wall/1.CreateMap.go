package Wall

import CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"

func CreateMap() {
	var temptilemap map[string]string = make(map[string]string)
	if len(CC.Tile_key) == len(CC.Tile_value) {
		for i := range CC.Tile_key {
			temptilemap[CC.Tile_key[i]] = CC.Tile_value[i]
		}
	}
	CC.Tilemap.Set(temptilemap)
}
