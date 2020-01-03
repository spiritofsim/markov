package markov

import (
	"bufio"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestNewChainWithOneNode(t *testing.T) {
	c := NewChain(getStringWordsScanner("hello"))
	require.Len(t, c.graph, 1)
	require.Equal(t, "hello", c.graph["hello"].Token)
	require.Empty(t, c.graph["hello"].next)
}

func TestNewChainWithLoopNodes(t *testing.T) {
	c := NewChain(getStringWordsScanner("hello world hello"))
	require.Len(t, c.graph, 2)

	helloNode := c.graph["hello"]
	worldNode := c.graph["world"]

	require.Equal(t, helloNode.next[worldNode], 1)
	require.Equal(t, worldNode.next[helloNode], 1)
}

func TestNewChainWeights(t *testing.T) {
	c := NewChain(getStringWordsScanner("hello world hello world"))
	require.Len(t, c.graph, 2)

	helloNode := c.graph["hello"]
	worldNode := c.graph["world"]

	require.Equal(t, helloNode.next[worldNode], 2)
	require.Equal(t, worldNode.next[helloNode], 1)
}

func TestNextOnAcyclicGraph(t *testing.T) {
	c := NewChain(getStringWordsScanner("hello big world"))

	token, err := c.Next("hello")
	require.NoError(t, err)
	require.Equal(t, "big", token)

	token, err = c.Next(token)
	require.NoError(t, err)
	require.Equal(t, "world", token)

	token, err = c.Next(token)
	require.EqualError(t, err, io.EOF.Error())
}

func TestRandNextNode(t *testing.T) {
	c := NewChain(getStringWordsScanner(""))
	i := 0
	c.randInt = func(n int) int {
		defer func() {
			i++
		}()
		switch i {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 2
		default:
			return 0
		}
	}

	nodes := map[*node]int{
		newNode("hello"): 1,
		newNode("world"): 2,
	}

	n := c.randNext(nodes)
	require.Equal(t, "hello", n.Token)
	n = c.randNext(nodes)
	require.Equal(t, "world", n.Token)
	n = c.randNext(nodes)
	require.Equal(t, "world", n.Token)
}

func TestNextReturnsFirstNodeOnEmptyToken(t *testing.T) {
	c := NewChain(getStringWordsScanner("hello big world"))

	_, err := c.Next("")
	require.NoError(t, err)
}

func TestNextReturnsNotFound(t *testing.T) {
	c := NewChain(getStringWordsScanner("1"))

	_, err := c.Next("2")
	require.EqualError(t, err, NotFound.Error())
}

func getStringWordsScanner(str string) *bufio.Scanner {
	s := bufio.NewScanner(strings.NewReader(str))
	s.Split(bufio.ScanWords)
	return s
}
