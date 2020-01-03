package markov

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

// Write like L.Tolstoy :)

// TestRunesSplitter demonstrates text generation based on blocks of runes
func TestRunesSplitter(t *testing.T) {
	printChainFromFile(t, "testdata/warandpeace.txt", bufio.ScanRunes, 3, "")
	printChainFromFile(t, "testdata/warandpeace.txt", bufio.ScanRunes, 4, "")
	printChainFromFile(t, "testdata/warandpeace.txt", bufio.ScanRunes, 5, "")
}

// TestWordsSplitter demonstrates text generation based on blocks of words
func TestWordsSplitter(t *testing.T) {
	printChainFromFile(t, "testdata/warandpeace.txt", bufio.ScanWords, 1, " ")
	printChainFromFile(t, "testdata/warandpeace.txt", bufio.ScanWords, 2, " ")
	printChainFromFile(t, "testdata/warandpeace.txt", bufio.ScanWords, 3, " ")
}

func printChainFromFile(t *testing.T, name string, split bufio.SplitFunc, cnt int, sep string) {
	chain := chainFromFile(t, name, split, cnt, sep)
	prev := ""
	for i := 0; i < 100; i++ {
		token, err := chain.Next(prev)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(token + sep)
		prev = token
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", 20))
}

func chainFromFile(t *testing.T, name string, split bufio.SplitFunc, cnt int, sep string) Chain {
	f, err := os.Open(name)
	require.NoError(t, err)
	defer func() {
		require.NoError(t, f.Close())
	}()

	s := bufio.NewScanner(f)
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		return miltiSplit(data, atEOF, cnt, split, []byte(sep))
	})

	return NewChain(s)
}
