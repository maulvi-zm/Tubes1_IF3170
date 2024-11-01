export type Solution = {
  solutions: SolutionItem[];
  type: string;
  additionalInfo: additionalInfo[];
};

export type additionalInfo = {
  itemName: string;
  itemValue: number;
};

export type SolutionItem = {
  iteration: number;
  state: number[];
  score: number;
  probability: number;
};
