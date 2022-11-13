package Wall

import CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"

// append tile map to wall
func (w *Wall) AppendTileMapToWall() {
	var (
		temp       []string
		repeattime int
	)
	for i := range CC.Tilemap.Get() {
		switch i[0] {
		case '1':
			repeattime = 2
		case '7':
			repeattime = 10
		default:
			repeattime = 4
		}
		for j := 0; j < repeattime; j++ {
			temp = append(temp, i)
		}
	}
	w.Wall.Set(temp)
}
