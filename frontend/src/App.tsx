import "./App.css";
import NumberedCube from "./components/numbered-cube";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import AlgorithmForm from "./components/algorithm-form";
import ResultsPage from "./components/result-page";
import VideoPlayerPage from "./components/videoplayer";
function App() {
  return (
    <Router>
      <div className="w-full">
        <div className="re top-0 left-0"></div>
        <Routes>
          <Route
            path="/"
            element={
              <>
                <NumberedCube />
                <AlgorithmForm />
              </>
            }
          />
          <Route path="/results" element={<ResultsPage />} />
          <Route path="/video" element={<VideoPlayerPage />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
