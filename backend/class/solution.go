package class

type SolutionItem struct {
	state     []int
	iteration int
	score     int
}

type Solution struct {
	solution []SolutionItem
}

func NewSolution() *Solution {
	return &Solution{solution: make([]SolutionItem, 0)}
}

func (s *Solution) AddSolutionItem(iteration int, score int, state []int) {
	s.solution = append(s.solution, SolutionItem{iteration: iteration, score: score, state: state})
}

func (s *Solution) GetSolution() []SolutionItem {
	return s.solution
}
