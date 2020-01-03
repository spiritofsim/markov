package markov

import (
	"bufio"
	"io"
	"math/rand"
	"time"
)

// Chain is a Markov chain
// Holds nodes graph with weight probability
// It also responsible for selecting next nodes based on its weight
type Chain struct {
	graph   map[string]*node
	randInt func(n int) int // for testing
}

// NewChain creates Markov chain from any scanner
// It reads from scanner until EOF and compose nodes graph
func NewChain(scanner *bufio.Scanner) Chain {
	graph := make(map[string]*node)
	var prev *node = nil
	for scanner.Scan() {
		token := scanner.Text()
		if prev == nil {
			graph[token] = newNode(token)
			prev = graph[token]
			continue
		}

		node, found := graph[token]
		if !found {
			node = newNode(token)
			graph[token] = node
		}
		prev.next[node]++
		prev = node
	}

	return Chain{
		graph:   graph,
		randInt: rand.New(rand.NewSource(time.Now().UnixNano())).Intn,
	}
}

// Next takes passed token and return next token based on its weight
// If token is empty, random token will be selected
// If token not empty and not found, NotFound err will be returned
// If no next tokens, io.EOF will be returned
func (c Chain) Next(token string) (string, error) {
	if token == "" {
		for k, _ := range c.graph {
			return k, nil
		}
	}

	node, found := c.graph[token]
	if !found {
		return "", NotFound
	}

	next := c.randNext(node.next)
	if next == nil {
		return "", io.EOF
	}

	return next.Token, nil
}

// randNext returns next weighted random node
func (c Chain) randNext(nodes map[*node]int) *node {
	sum := 0
	for _, weight := range nodes {
		sum += weight
	}
	if sum == 0 {
		return nil
	}

	r := c.randInt(sum)
	for i, weight := range nodes {
		r -= weight
		if r < 0 {
			return i
		}
	}
	return nil
}
