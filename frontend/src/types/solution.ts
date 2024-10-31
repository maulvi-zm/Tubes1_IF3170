export type Solution = {
  solutions: SolutionItem[];
};

export type SolutionItem = {
  iteration: number;
  state: number[];
  score: number;
};
