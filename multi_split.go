package markov

import "bufio"

// miltiSplit is a wrapper around bufio.SplitFunc that returns not each token, but combines it
// cnt - number of tokens to combine
// sep - tokens separator
func miltiSplit(data []byte, atEOF bool, cnt int, split bufio.SplitFunc, sep []byte) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	result := make([]byte, 0)
	advance := 0

	for i := 0; i < cnt; i++ {
		//if advance > len(data)-1{
		//	advance = len(data)-1
		//}
		a, token, err := split(data[advance:], i == cnt-1)
		if err != nil {
			return advance, result, err
		}
		advance += a
		result = append(result, token...)

		if advance == len(data) {
			break
		}

		if i < cnt-1 {
			result = append(result, sep...)
		}
	}
	return advance, result, nil
}
