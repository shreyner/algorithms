package check_win_even

var (
	StatusWIN  = "WIN"
	StatusFAIL = "FAIL"
)

func CheckWinEvent(a, b, c int) string {
	aE, bE, cE := a%2 == 0, b%2 == 0, c%2 == 0

	if aE == bE && bE == cE && cE == aE {
		return StatusWIN
	}

	return StatusFAIL
}
