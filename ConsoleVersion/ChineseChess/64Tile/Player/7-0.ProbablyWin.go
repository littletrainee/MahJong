package Player

func probablywin(temp []string, meld [][]string) bool {
	var tempplayer Player
	tempplayer.Hand.Set(temp)
	tempplayer.Meld.Set(meld)
	tempplayer.TsumoCheck()
	return tempplayer.Iswin.Get()
}
