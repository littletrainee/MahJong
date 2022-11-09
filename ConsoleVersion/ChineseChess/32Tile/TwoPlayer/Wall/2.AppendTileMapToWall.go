package Wall

import CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"

// append tile map to wall
func (w *Wall) AppendTileMapToWall() {
	var repeattime int
	var temp []string
	for i := range CC.Tilemap.Get() {
		switch i[0] {
		case '1':
			repeattime = 1
		case '7':
			repeattime = 5
		default:
			repeattime = 2
		}
		for j := 0; j < repeattime; j++ {
			temp = append(temp, i)
		}
	}
	w.Wall.Set(temp)
}
