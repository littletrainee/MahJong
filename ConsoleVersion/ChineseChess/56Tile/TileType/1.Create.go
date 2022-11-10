package TileType

import (
	"github.com/littletrainee/GetSet"
	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Wall"
)

func sortHand(hand *GetSet.Type[[]string]) {
	var p Player.Player
	p.Hand.Set(hand.Get())
	p.SortHand()
	hand.Set(p.Hand.Get())
}

func (tt *TileType) Create(thisplayer, otherplayer *Player.Player, wall *Wall.Wall, gamestate *CV.GameState) {
	if thisplayer.IsConcealed.Get() {
		tt.hand.Set(thisplayer.Hand.Get())
	} else {
		tt.meld.Set(thisplayer.Meld.Get())
	}
	if len(thisplayer.Hand.Get()) != 8 || len(thisplayer.Hand.Get()) != 5 || len(thisplayer.Hand.Get()) != 2 {
		tt.hand.Set(append(tt.hand.Get(), otherplayer.River.Get()[len(otherplayer.River.Get())-1]))
	}
	sortHand(&tt.hand)
	tt.gameround.Set(gamestate.GameRound.Get())
	tt.gameturn.Set(gamestate.GameTurn.Get())
	tt.istsumo.Set(thisplayer.IsTsumo.Get())
	tt.lasttile.Set(wall.LastTile.Get())
	tt.istenpai.Set(thisplayer.TenPai.Get())
	tt.eye.Set(thisplayer.Eye.Get())
	tt.isconcealed.Set(thisplayer.IsConcealed.Get())
	tt.twokang.Set(thisplayer.TwoKang.Get())
	tt.winonthewalltail.Set(thisplayer.WinOnTheWallTail.Get())
	tt.bookmaker.Set(thisplayer.Bookmaker.Get())
	tt.continuetobookmaker.Set(thisplayer.ContinueToBookmaker.Get())
}
