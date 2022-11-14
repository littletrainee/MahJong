package TileType

func (tt *TileType) SetYakuAndPoint() {
	var (
		c         uint8 = tt.colorlength.Get()
		k         uint8 = tt.kindlength.Get()
		oya, koya uint8
	)

	tt.yaku.Set(ChnNumber[c] + "色  " + ChnNumber[k] + "種")
	// tt.yaku.Set(strconv.Itoa(int(c)) + "Color" + strconv.Itoa(int(k)) + "Kind")
	switch c {
	case 1:
		switch k {
		case 1:
			oya = 200
			koya = 100
		case 2:
			oya = 60
			koya = 30
		case 3:
			oya = 60
			koya = 30
		}
	case 2:
		switch k {
		case 1:
			oya = 80
			koya = 40
		case 2:
			oya = 30
			koya = 15
		case 3:
			oya = 10
			koya = 5
		}
	case 3:
		switch k {
		case 1:
			oya = 60
			koya = 30
		case 2:
			oya = 10
			koya = 5
		case 3:
			oya = 30
			koya = 15
		case 9:
			oya = 6
			koya = 3
		}

	}
	tt.score.Set([]uint8{oya, koya})
}
