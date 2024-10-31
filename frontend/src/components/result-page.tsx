"use client";

import { Button } from "@/components/ui/button";
import { Solution } from "@/types/solution";
import { Link, useLocation } from "react-router-dom";
import Chart from "./chart";

export default function ResultsPage() {
  const location = useLocation();
  const data: Solution = location.state?.data;

  // Prepare data for the chart
  const chartData = data?.solutions.map((item) => ({
    iteration: item.iteration,
    score: item.score,
  }));

  return (
    <div className="flex flex-col justify-center mt-8 p-4 w-screen">
      {data && (
        <Chart
          {...{
            label: "Score",
            chartData,
            title: "Score changes over iterations",
          }}
        />
      )}
      <Link to="/">
        <Button>Back to Form</Button>
      </Link>
    </div>
  );
}
