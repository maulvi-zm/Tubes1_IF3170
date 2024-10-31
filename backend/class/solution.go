package class

import (
	"encoding/json"
)

type SolutionItem struct {
	State     []int `json:"state"`
	Iteration int   `json:"iteration"`
	Score     int   `json:"score"`
	DeltaE    int   `json:"deltaE"`
}

type addtitionalInfor struct {
	ItemName  string `json:"itemName"`
	ItemValue string `json:"itemValue"`
}

type Solution struct {
	Type           string             `json:"type"`
	Solution       []SolutionItem     `json:"solutions"`
	AdditionalInfo []addtitionalInfor `json:"additionalInfo"`
}

func NewSolution() *Solution {
	return &Solution{Solution: make([]SolutionItem, 0), Type: "default", AdditionalInfo: make([]addtitionalInfor, 0)}
}

func addAdditionalInfo(s *Solution, itemName string, itemValue string) {
	s.AdditionalInfo = append(s.AdditionalInfo, addtitionalInfor{ItemName: itemName, ItemValue: itemValue})
}

func setType(s *Solution, solutionType string) {
	s.Type = solutionType
}

func (s *Solution) AddSolutionItem(iteration int, score int, state []int, deltaE ...int) {
	defaultDeltaE := 0
	if len(deltaE) > 0 {
		defaultDeltaE = deltaE[0]
	}
	s.Solution = append(s.Solution, SolutionItem{Iteration: iteration, Score: score, State: state, DeltaE: defaultDeltaE})
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
