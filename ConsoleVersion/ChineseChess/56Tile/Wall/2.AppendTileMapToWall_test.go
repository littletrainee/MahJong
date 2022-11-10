package Wall

import (
	"fmt"
	"testing"
)

func TestAppendTileMapToWall(t *testing.T) {
	var w Wall
	Init()
	CreateMap()
	w.AppendTileMapToWall()
	fmt.Println(w)
}
