package hackerrank

// towerBreakers determines the winner of the Tower Breakers game.
func towerBreakers(n int32, m int32) int32 {
	if m == 1 {
		return 2
	} else {
		if n%2 == 0 {
			return 2
		} else {
			return 1
		}
	}
}
