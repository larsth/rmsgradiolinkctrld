package gps

import (
	"errors"
	"io"
	"sync"

	"github.com/larsth/rmsgradiolinkctrld/memorybuffers"
)

const (
	//12 K bytes = 3072 x 32 bits(1 rune==32bit==4 bytes)
	DefaultBufferRunesSize = 3072

	//4 K bytes = 4096 x 8 bits(1 byte=8 bits)
	DefaultBufferBytesSize = 4096
)

var (
	ErrGpsdJsonDocumentHadNotBeenParsed = errors.New(
		"ERROR: The gpsd JSON document had not been parsed")
	//'bufferSize' is not a const, because its size is set to a much lower
	//value in unit tests
	ErrIoNilReader           = errors.New("Error: the io.Reader is nil")
	ErrBufferSizeLessThanOne = errors.New("The buffer size is less than 1")
)

//Type Scanner is a LL(1) Recursive-Descent Lexical Analyzer
//that extracts gpsd JSON documents from an io.Reader stream
//of bytes.
type Scanner struct {
	mutex    sync.Mutex
	ggc      *memorybuffers.GetGiveChans
	ioReader io.Reader
	input    []rune
	c        rune //current character
}

func NewScanner(ggc *memorybuffers.GetGiveChans) *Scanner {
	s := new(Scanner)
	s.Init(ggc)
	return s
}

func (s *Scanner) Init(ggc *memorybuffers.GetGiveChans) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ggc = ggc
}
