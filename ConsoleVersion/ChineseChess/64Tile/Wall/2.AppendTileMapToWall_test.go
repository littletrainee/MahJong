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
	if len(w.Wall.Get()) != 64 {
		t.Error("Wrong")
	} else {

		fmt.Println(w)
	}
}
