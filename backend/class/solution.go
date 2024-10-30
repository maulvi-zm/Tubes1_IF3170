package class

import (
	"encoding/json"
)

type SolutionItem struct {
	State     []int `json:"state"`
	Iteration int   `json:"iteration"`
	Score     int   `json:"score"`
}

type Solution struct {
	Solution []SolutionItem `json:"solutions"`
}

func NewSolution() *Solution {
	return &Solution{Solution: make([]SolutionItem, 0)}
}

func (s *Solution) AddSolutionItem(iteration int, score int, state []int) {
	s.Solution = append(s.Solution, SolutionItem{Iteration: iteration, Score: score, State: state})
}

func (s *Solution) GetSolution() []SolutionItem {
	return s.Solution
}

func (s *Solution) ConvertToJson() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
