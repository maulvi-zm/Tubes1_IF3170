import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import Draggable from "react-draggable";
import { Solution } from "@/types/solution";

export default function AlgorithmForm() {
  const [algorithm, setAlgorithm] = useState("");
  const [params, setParams] = useState({
    maxStochasticMove: "",
    maxSidewayMove: "",
    maxRandomRestart: "",
    populationNum: "",
    iteration: "",
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    let endpoint = "";
    let body: Record<string, unknown> = {};

    switch (algorithm) {
      case "steepest":
        endpoint = "/hill-climbing/steepest";
        break;
      case "stochastic":
        endpoint = "/hill-climbing/stochastic";
        body = {maxStochasticMove: parseInt(params.maxStochasticMove) };
        break;
      case "sidewaymove":
        endpoint = "/hill-climbing/sidewaymove";
        body = { maxSidewayMove: parseInt(params.maxSidewayMove) };
        break;
      case "randomrestart":
        endpoint = "/hill-climbing/randomrestart";
        body = { maxRandomRestart: parseInt(params.maxRandomRestart) };
        break;
      case "genetic":
        endpoint = "/genetic-algorithm";
        body = {
          populationNum: parseInt(params.populationNum),
          iteration: parseInt(params.iteration),
        };
        break;
      case "simulated":
        endpoint = "/simulated-annealing";
        break;
    }

    try {
      console.log(JSON.stringify(body));
      const response = await fetch("http://localhost:8080" + endpoint, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });
      const data: Solution = await response.json();
      localStorage.setItem("solution", JSON.stringify(data));
      window.open(
        "http://localhost:3000/results",
        "_blank",
        "noopener,noreferrer",
      );
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <Draggable>
      <form
        onSubmit={handleSubmit}
        className="space-y-4 max-w-md absolute left-[50px] top-1/2 cursor-grab bg-white p-4 rounded-lg"
      >
        <div>
          <Label htmlFor="algorithm">Algorithm</Label>
          <Select onValueChange={(value) => setAlgorithm(value)}>
            <SelectTrigger id="algorithm">
              <SelectValue placeholder="Select algorithm" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="steepest">Hill Climbing Steepest</SelectItem>
              <SelectItem value="sidewaymove">Hill Climbing Sideway</SelectItem>
              <SelectItem value="stochastic">Hill Climbing Stochastic</SelectItem>
              <SelectItem value="randomrestart">Hill Climbing Random Restart</SelectItem>
              <SelectItem value="genetic">Genetic Algorithm</SelectItem>
              <SelectItem value="simulated">Simulated Annealing</SelectItem>
            </SelectContent>
          </Select>
        </div>

        {algorithm === "stochastic" && (
          <div>
            <Label htmlFor="maxStochasticMove">Max Stochastic Move</Label>
            <Input
              id="maxStochasticMove"
              type="number"
              value={params.maxStochasticMove}
              onChange={(e) =>
                setParams({ ...params, maxStochasticMove: e.target.value })
              }
            />
          </div>
        )}

        {algorithm === "sidewaymove" && (
          <div>
            <Label htmlFor="maxSidewayMove">Max Sideway Move</Label>
            <Input
              id="maxSidewayMove"
              type="number"
              value={params.maxSidewayMove}
              onChange={(e) =>
                setParams({ ...params, maxSidewayMove: e.target.value })
              }
            />
          </div>
        )}

        {algorithm === "randomrestart" && (
          <div>
            <Label htmlFor="maxRandomRestart">Max Random Restart</Label>
            <Input
              id="maxRandomRestart"
              type="number"
              value={params.maxRandomRestart}
              onChange={(e) =>
                setParams({ ...params, maxRandomRestart: e.target.value })
              }
            />
          </div>
        )}

        {algorithm === "genetic" && (
          <>
            <div>
              <Label htmlFor="populationNum">Population Number</Label>
              <Input
                id="populationNum"
                type="number"
                value={params.populationNum}
                onChange={(e) =>
                  setParams({ ...params, populationNum: e.target.value })
                }
              />
            </div>
            <div>
              <Label htmlFor="iteration">Iteration</Label>
              <Input
                id="iteration"
                type="number"
                value={params.iteration}
                onChange={(e) =>
                  setParams({ ...params, iteration: e.target.value })
                }
              />
            </div>
          </>
        )}

        <Button type="submit" className="z-40">
          Submit
        </Button>
      </form>
    </Draggable>
  );
}
