package memorybuffers

type GetGiveChans struct {
	GetBytes  chan []byte
	GetRunes  chan []rune
	GiveBytes chan []byte
	GiveRunes chan []rune
}
