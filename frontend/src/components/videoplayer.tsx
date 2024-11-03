import React, { useState, useEffect } from 'react';
import Component from './parameteried-cube';
import { Solution } from "@/types/solution";
import Draggable from "react-draggable";
import { exportSolutionToFile,loadSolutionFromFile } from '@/types/loader';
import { Button } from "@/components/ui/button";


interface VideoPlayerProps {
  states?: Solution; 
}

const VideoPlayer: React.FC<VideoPlayerProps> = ({ states }) => {
  const [isPlaying, setIsPlaying] = useState(false);
  const [currentFrame, setCurrentFrame] = useState(0);
  const [playbackRate, setPlaybackRate] = useState(1);
  const [filename, setFilename] = useState<string>('filename.json');
  const [inputFileName, setInputFileName] = useState<string>('')
  const [data, setData] = useState<Solution | null>(states? states:null);
  const [isLoadVisible, setIsLoadVisible] = useState(false);
  const [isExportVisible, setIsExportVisible] = useState(false);
  
  useEffect(() => {
    document.title = `[Video] ${data?.type ?? "Video"}`;
    let interval: NodeJS.Timeout | undefined;

    if (isPlaying && data?.solutions.length) {
      interval = setInterval(() => {
        setCurrentFrame((prevFrame) =>
          (prevFrame + 1) % data.solutions.length
        );
      }, 1000 / playbackRate);
    }

    return () => {
      if (interval) clearInterval(interval);
    };
  }, [isPlaying, playbackRate, data?.solutions.length]);

  
  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setInputFileName(file.name); 
      const solution = await loadSolutionFromFile(file); 
      setData(solution);
      setCurrentFrame(0);
    }
    console.log("Loaded Solution from: ", inputFileName);

  };

  const handleExportSolution = () => {
    if (data) {
      exportSolutionToFile(data, filename); 
    } else {
      console.error("No solution loaded to export.");
    }
  };

  const togglePlayPause = () => {
    setIsPlaying((prev) => !prev);
  };

  const handleProgressChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const frame = parseInt(e.target.value, 10);
    setCurrentFrame(frame);
    if (isPlaying) {
      setIsPlaying(false); 
    }
  };
  const toggleLoad = ()=>{
    setIsLoadVisible((prev)=>!prev);
  }
  const toggleExport = ()=>{
    setIsExportVisible((prev)=>!prev);
  }
  return (
    <div>

      <Component numbers={data?.solutions[currentFrame].state ?? [0]} />
      <Draggable>
        <div className=" absolute bottom-0 left-1/2 flex flex-col items-center bg-white p-4 rounded shadow-md space-y-4 min-w-[600px] max-w-[800px]">
          <button
            onClick={togglePlayPause}
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          >
            {isPlaying ? 'Pause' : 'Play'}
          </button>
          
          <input
            type="range"
            min="0"
            max={(data?.solutions.length ?? 1) - 1}
            value={currentFrame}
            onChange={handleProgressChange}
            className="w-full bg-white"
          />
          
          <label className="flex flex-row items-center space-y-2">
            <span className="text-gray-700 mr-[20px]">Playback Speed:</span>
            <select
              value={playbackRate}
              onChange={(e) => setPlaybackRate(parseFloat(e.target.value))}
              className="border border-gray-300 rounded p-2 text-black bg-white"
            >
              <option value="0.5">0.5x</option>
              <option value="1">1x</option>
              <option value="2">2x</option>
              <option value="4">4x</option>
              <option value="8">8x</option>
              <option value="16">16x</option>
              <option value="32">32x</option>
            </select>
          </label>
          <button onClick={toggleExport} className='bg-gray-100'><h4>Export</h4></button> 
          <div className={` ${isExportVisible ? 'block' : 'hidden'}`}>
            <label htmlFor="filename" className='mr-[10px]'>Enter filename:</label>
            <input className='bg-white border-gray-400 border-[2px]'
              type="text"
              id="filename"
              value={filename}
              onChange={(e) => setFilename(e.target.value)
              }
            />
            <button className='bg-gray-50' type="submit" onClick={handleExportSolution}>Export</button>
          </div>
          <button onClick={toggleLoad} className='bg-gray-100'><h4>Load</h4></button> 
        <div className={` ${isLoadVisible ? 'block' : 'hidden'}`}>
          <input
          type="file"
          accept=".json" 
          onChange={handleFileChange} 
        />
        </div>
        
        </div>
      </Draggable>
      <Draggable>
        <div className="absolute left-4 top-1/4 p-4 flex flex-col flex-wrap gap-4 max-w-[400px] bg-white rounded-lg cursor-grab">
            <div className="flex gap-2 m-auto items-center">
              <h2 className="text-2xl font-bold">Initial/Final State</h2>
            </div>
            <Button 
              onClick={()=>setCurrentFrame(0)}
              className={currentFrame === 0 ? "bg-blue-500" : ""}
            >
              Initial State
            </Button>
            <Button 
              onClick={()=>setCurrentFrame(data ? data.solutions.length - 1 : 0)}
              className={currentFrame === (data ? data.solutions.length - 1 : 0) ? "bg-blue-500" : ""}
            >
              Final State
            </Button>
        </div> 
      </Draggable>
    </div>
  );
};

export default function VideoPlayerPage() {
  const [data, setData] = useState<Solution | null>(null);

  useEffect(() => {
    const storedData = localStorage.getItem("solution");
    if (storedData) {
      setData(JSON.parse(storedData));
    }
  }, []);
  
  return data ? <VideoPlayer states={data} /> : <div>Loading...</div>;
};
