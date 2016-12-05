package main

import (
    "sort"
    "fmt"
    )

func QuantifyLuck (matchups []Matchup) Teams {
    teams := Teams{}

    for _, m := range matchups {
        m[0].Result = m[0].Score > m[1].Score
        m[1].Result = !m[0].Result
        teams = append(teams, m[0])
        teams = append(teams, m[1])
    }

    sort.Sort(teams)

    for i, t := range teams {
        teams[i].Odds = float64(i) / float64(len(teams) - 1)
        fmt.Println("i: ", float64(i))
        fmt.Println("num teams - 1: ", float64(len(teams) - 1))
        fmt.Println("result: ", float64(i) / float64(len(teams) - 1))
        teams[i].Luck = float64(btoi(t.Result)) - teams[i].Odds
    }

    return teams
}

func btoi (b bool) int {
    if b {
        return 1
    }
    return 0
}
