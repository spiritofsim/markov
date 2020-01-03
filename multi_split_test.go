package markov

import (
	"bufio"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestScanSingleRunes(t *testing.T) {
	checkScanner(t, bufio.ScanRunes, 1, "", "")
	checkScanner(t, bufio.ScanRunes, 1, "", "а", "а")
	checkScanner(t, bufio.ScanRunes, 1, "", "аб", "а", "б")
}

func TestScanDoubleRunes(t *testing.T) {
	checkScanner(t, bufio.ScanRunes, 2, "", "")
	checkScanner(t, bufio.ScanRunes, 2, "", "а", "а")
	checkScanner(t, bufio.ScanRunes, 2, "", "аб", "аб")
	checkScanner(t, bufio.ScanRunes, 2, "", "абв", "аб", "в")
	checkScanner(t, bufio.ScanRunes, 2, "", "абвг", "аб", "вг")
}

func TestScanSep(t *testing.T) {
	checkScanner(t, bufio.ScanRunes, 2, " ", "")
	checkScanner(t, bufio.ScanRunes, 2, " ", "а", "а")
	checkScanner(t, bufio.ScanRunes, 2, " ", "аб", "а б")
	checkScanner(t, bufio.ScanRunes, 2, " ", "абв", "а б", "в")
	checkScanner(t, bufio.ScanRunes, 2, " ", "абвг", "а б", "в г")
}

func TestScanSingleWord(t *testing.T) {
	checkScanner(t, bufio.ScanWords, 1, "", "")
	checkScanner(t, bufio.ScanWords, 1, "", "w1", "w1")
	checkScanner(t, bufio.ScanWords, 1, "", "w1 w2", "w1", "w2")
}

func TestScanDoubleWords(t *testing.T) {
	checkScanner(t, bufio.ScanWords, 2, "", "")
	checkScanner(t, bufio.ScanWords, 2, "", "w1", "w1")
	checkScanner(t, bufio.ScanWords, 2, "", "w1 w2", "w1w2")
	checkScanner(t, bufio.ScanWords, 2, "", "w1 w2 w3", "w1w2", "w3")
	checkScanner(t, bufio.ScanWords, 2, "", "w1 w2 w3 w4", "w1w2", "w3w4")
}

func checkScanner(t *testing.T, splitter bufio.SplitFunc, count int, sep string, str string, expectations ...string) {
	scanner := getScanner(str, splitter, count, sep)
	for _, ex := range expectations {
		require.True(t, scanner.Scan())
		require.Equal(t, ex, scanner.Text())
	}
	require.False(t, scanner.Scan())
}

func getScanner(str string, splitter bufio.SplitFunc, count int, sep string) *bufio.Scanner {
	s := bufio.NewScanner(strings.NewReader(str))
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		return miltiSplit(data, atEOF, count, splitter, []byte(sep))
	})
	return s
}
