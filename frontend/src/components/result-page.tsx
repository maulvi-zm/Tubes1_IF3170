import { useSearchParams } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";
export default function ResultsPage() {
  const [searchParams] = useSearchParams();
  const data = searchParams.get("data");

  return (
    <div className="max-w-2xl mx-auto mt-8 p-4">
      <h1 className="text-2xl font-bold mb-4">Results</h1>
      <pre className="bg-gray-100 p-4 rounded-md overflow-auto">
        {data ? JSON.stringify(JSON.parse(data), null, 2) : "No data available"}
      </pre>
      <Link to="/">
        <Button className="mt-4">Back to Form</Button>
      </Link>
    </div>
  );
}
