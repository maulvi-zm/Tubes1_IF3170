import { Solution } from "@/types/solution";
import Chart from "./chart";
import AdditionalInformation from "./additional-information";
import { useEffect, useState } from "react";

export default function ResultsPage() {
  const [data, setData] = useState<Solution>();

  useEffect(() => {
    const data = localStorage.getItem("solution");
    if (data) {
      const parsedData = JSON.parse(data);
      setData(parsedData);
      document.title = `[Results] ${parsedData.type}`;
    }
  }, []);

  // Prepare data for the chart
  const chartData =
    data?.solutions.map((item) => ({
      iteration: item.iteration,
      score: item.score,
    })) || [];

  return (
    <div className="flex flex-col justify-center mt-8 p-4 w-screen">
      {data && (
        <>
          {data.additionalInfo.length > 0 && (
            <AdditionalInformation
              {...{
                title: "Additional Information",
                additionalInfo: data.additionalInfo,
              }}
            />
          )}
          <Chart
            label="Score"
            chartData={chartData}
            title="Score changes over iterations"
          />
          {data.type === "Simulated Annealing" && (
            <Chart
              {...{
                label: "Probability",
                chartData: data.solutions.map((item) => ({
                  iteration: item.iteration,
                  score: item.probability,
                })),
                title: "Probability changes over iterations",
              }}
            />
          )}
        </>
      )}
    </div>
  );
}
