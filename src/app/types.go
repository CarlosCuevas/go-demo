package main

// Team
type Team struct {
    Name string   `json:"name,omitempty"`
    Score  int   `json:"score,omitempty"`
    Result bool `json:"result,omitempty"`
    Odds float64 `json:"odds,omitempty"`
    Luck float64 `json:"luck,omitempty"`
}

type Teams []Team

func (slice Teams) Len() int {
    return len(slice)
}

func (slice Teams) Less(i, j int) bool {
    return slice[i].Score < slice[j].Score;
}

func (slice Teams) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

// Matchup
type Matchup [2]Team
