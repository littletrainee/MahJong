package main

// import "github.com/littletrainee/MahJong/ConsoleVersion/Dice/TwoPlayer/Handler"
// import "github.com/littletrainee/MahJong/ConsoleVersion/Dice/ThreePlayer/Handler"
// import "github.com/littletrainee/MahJong/ConsoleVersion/Dice/FourPlayer/Handler"

// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TwoPlayer/Handler"
// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/ThreePlayer/Handler"
// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/FourPlayer/Handler"

// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/TwoPlayer/Handler"
// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/ThreePlayer/Handler"
// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/FourPlayer/Handler"

// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/TwoPlayer/Handler"
// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/ThreePlayer/Handler"
// import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/FourPlayer/Handler"

import "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/TwoPlayer/Handler"

func main() {
	var handler Handler.Handler
	handler.Start()
	handler.Update()
}
