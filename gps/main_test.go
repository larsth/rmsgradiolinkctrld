package gps

import (
	"os"
	"testing"

	"github.com/larsth/rmsgradiolinkctrld/memorybuffers"
)

var gst_testing memorybuffers.GetGiveChans
var byteSliceRecycler_testing func()
var runeSliceRecycler_testing func()

func TestMain(m *testing.M) {
	byteSliceRecycler_testing, gst_testing.GetBytes, gst_testing.GiveBytes =
		memorybuffers.MakeByteSliceRecycler()
	go byteSliceRecycler_testing()
	runeSliceRecycler_testing, gst_testing.GetRunes, gst_testing.GiveRunes =
		memorybuffers.MakeRuneSliceRecycler()
	go runeSliceRecycler_testing()
	os.Exit(m.Run())
}
