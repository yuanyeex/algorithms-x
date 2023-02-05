package prb0605

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			checkLeft := true
			if i > 0 {
				checkLeft = flowerbed[i-1] == 0
			}
			checkRight := true
			if i < len(flowerbed)-1 {
				checkRight = flowerbed[i+1] == 0
			}

			if checkLeft && checkRight {
				n--
				flowerbed[i] = 1
				if n <= 0 {
					return true
				}
			}
		}
	}

	return n <= 0
}
