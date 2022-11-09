package Interface

type IAppendToWall interface {
	AppendTileMapToWall()
}

type IPlayer interface {
	DrawCard()
	IAppendToWall
	IWin
}

type IWin interface {
	RonCheck()
	TsumoCheck()
}

type IWall interface {
	IAppendToWall
}
