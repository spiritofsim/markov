package markov

// node represents chain graph
type node struct {
	Token string
	next  map[*node]int
}

func newNode(token string) *node {
	return &node{
		Token: token,
		next:  make(map[*node]int),
	}
}

func (n node) String() string {
	return n.Token
}
