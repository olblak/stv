package stv

import (
	"errors"
	//"fmt"
	"sort"
)

// rankMapStringInt takes a map of [string]int  as parameters then
// returns a revert sorted array of keys
// sorted based on their integer values
// Remark: an array is needed here as a map datastructure is always unordered
func rankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}

// contains takes as parameter a list of string and a value string
// then returns true if that value string is located in the list
func contains(list []string, value string) (found bool) {
	found = false
	for _, a := range list {
		if a == value {
			found = true
		}
	}
	return found
}

// Process calcules vote results
func Process(votes *[]map[string]int, choices int) (result []map[string]int, winner []string, rejected []string, err error) {

	result = []map[string]int{{}}
	rejected = []string{}
	winner = []string{}

	// Max round still need to be calculated but
	// we need somehow the number of possible choices
	// but for now we only specify a big value
	maxRound := 1000000
	for round := 0; round < maxRound; round++ {
		//fmt.Printf("\n\nROUND %d \n\n", round)
		if round == maxRound {
			err = errors.New("round limit reached, results can't be trusted")
			//fmt.Println(err)
			return nil, nil, nil, err
		}
		for _, vote := range *votes {

			// Sorts vote by key position
			rankedChoice := rankMapStringInt(vote)

			for _, r := range rankedChoice {
				// If a choice has already been validated or rejected then
				// then look for the next one in the row
				if contains(rejected, r) || contains(winner, r) {
					continue
				}
				result[round][r] = 1
				if _, ok := result[round][r]; !ok {
					break

				} else if ok {
					result[round][r] = result[round][r] + 1
					break
				}
			}

		}

		rankedResult := rankMapStringInt(result[round])

		limit := choices - len(winner)
		if len(rankedResult) <= limit {
			for counter := 1; counter <= limit && counter <= len(rankedResult); counter++ {
				winner = append(winner, rankedResult[len(rankedResult)-counter])
			}
		}

		// If we get more results than what we need then
		// we reject the last one and start another round
		if len(rankedResult) > limit {
			if !contains(winner, rankedResult[0]) {
				rejected = append(rejected, rankedResult[0])
			}
		}

		//fmt.Printf("Needed choice: %v\n", choices)
		//fmt.Printf("Results: %v\n", result[round])
		//fmt.Printf("Winner: %v\n", winner)
		//fmt.Printf("Rejected: %v\n", rejected)

		if choices <= len(winner) {
			//fmt.Println("Done, all results found")
			break
		} else {
			result = append(result, map[string]int{})
		}
	}

	return result, winner, rejected, err
}
