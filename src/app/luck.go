package main

import "sort"

func QuantifyLuck(matchups []Matchup) Teams {
	teams := Teams{}

    // extract teams from matchups
	for _, m := range matchups {
		m[0].Result = m[0].Score > m[1].Score
		m[1].Result = !m[0].Result
		teams = append(teams, m[0])
		teams = append(teams, m[1])
	}

    // sort teams by score
	sort.Sort(teams)

	for i, t := range teams {
        // index / number of other teams tells us how many matchups you would've won
		teams[i].Odds = float64(i) / float64(len(teams)-1)
        // taking the difference of your actual result (win = 1, loss = 0) and
        // your odds of winning tells us how lucky/unlucky you are (higher the
        // positive number is lucky, lower the negative number is unlucky)
		teams[i].Luck = float64(btoi(t.Result)) - teams[i].Odds
	}

	return teams
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
