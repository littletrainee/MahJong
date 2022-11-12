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
	temp := thisplayer.Hand.Get()
	if len(thisplayer.Meld.Get()) != 0 {
		for _, meld := range thisplayer.Meld.Get() {
			temp = append(temp, meld[:3]...)
		}
	}
	if len(temp) != 8 && len(temp) != 5 && len(temp) != 2 {
		temp = append(temp, otherplayer.River.Get()[len(otherplayer.River.Get())-1])
	}
	tt.hand.Set(temp)
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
