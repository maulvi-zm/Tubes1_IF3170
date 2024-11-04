package class

import (
	"encoding/json"
)

type SolutionItem struct {
	State       []int   `json:"state"`
	Iteration   int     `json:"iteration"`
	Score       int     `json:"score"`
	Probability float64 `json:"probability"`
}

type additionalInfor struct {
	// Additional info for genetic algorithm
	ItemName  string  `json:"itemName"`
	ItemValue float64 `json:"itemValue"`
}

type Solution struct {
	Type           string            `json:"type"`
	Solution       []SolutionItem    `json:"solutions"`
	AdditionalInfo []additionalInfor `json:"additionalInfo"`
}

func NewSolution() *Solution {
	return &Solution{Solution: make([]SolutionItem, 0), Type: "default", AdditionalInfo: make([]additionalInfor, 0)}
}

func (s *Solution) AddAdditionalInfo(itemName string, itemValue float64) {
	s.AdditionalInfo = append(s.AdditionalInfo, additionalInfor{itemName, itemValue})
}

func (s *Solution) SetType(solutionType string) {
	s.Type = solutionType
}

func (s *Solution) AddSolutionItem(iteration int, score int, state []int, probability ...float64) {
	defaultProbability := 0.0
	if len(probability) > 0 {
		defaultProbability = probability[0]
	}
	s.Solution = append(s.Solution, SolutionItem{Iteration: iteration, Score: score, State: state, Probability: defaultProbability})
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

func (s *Solution) AddElapsedTime(totalTime float64) {
	s.AdditionalInfo = append(s.AdditionalInfo, additionalInfor{"Search Duration", float64(totalTime)})
}

func (s *Solution) AddLastScore(lastScore int) {
	s.AdditionalInfo = append(s.AdditionalInfo, additionalInfor{"Last Objective Function", float64(lastScore)})
}
