package main

// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TwoPlayer/Handler"

import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/ThreePlayer/Handler"

// import "github.com/littletrainee/MahJong/ConsoleVersion/Dice/TwoPlayer/Handler"

// import "github.com/littletrainee/MahJong/ConsoleVersion/Dice/ThreePlayer/Handler"

// import "github.com/littletrainee/MahJong/ConsoleVersion/Dice/FourPlayer/Handler"

func main() {
	var handler Handler.Handler
	handler.Start()
	handler.Update()
}
