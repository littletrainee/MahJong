package TileType

import (
	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Wall"
)

func (k *TileType) Create(thisplayer, otherplayer *Player.Player, wall *Wall.Wall, gamestate *CV.GameState) {
	hand := thisplayer.Hand.Get()
	if len(thisplayer.Meld.Get()) != 0 {
		hand = append(hand, thisplayer.Meld.Get()[0]...)
	}
	if len(hand) != 5 {
		hand = append(hand, otherplayer.River.Get()[len(otherplayer.River.Get())-1])
	}
	k.hand.Set(hand)
	k.gameround.Set(gamestate.GameRound.Get())
	k.gameturn.Set(gamestate.GameTurn.Get())
	k.istsumo.Set(thisplayer.IsTsumo.Get())
	k.lasttile.Set(wall.LastTile.Get())
	k.istenpai.Set(thisplayer.TenPai.Get())
}
