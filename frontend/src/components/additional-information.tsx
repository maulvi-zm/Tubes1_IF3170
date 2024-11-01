import { ResponsiveContainer } from "recharts";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "./ui/card";
import { additionalInfo } from "@/types/solution";

interface ChartProps {
  title: string;
  additionalInfo: additionalInfo[];
}

const Chart: React.FC<ChartProps> = ({ title, additionalInfo }) => {
  return (
    <Card className="mb-8">
      <CardHeader>
        <CardTitle>{title}</CardTitle>
        <CardDescription>Below is the result information </CardDescription>
      </CardHeader>
      <CardContent>
        <ResponsiveContainer width="100%" height="100%">
          <div>
            {additionalInfo.map((item) => (
              <div key={item.itemName}>
                <p>
                  {item.itemName}: {item.itemValue}
                </p>
              </div>
            ))}
          </div>
        </ResponsiveContainer>
      </CardContent>
    </Card>
  );
};

export default Chart;
