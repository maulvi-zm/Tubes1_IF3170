import { Solution } from './solution.ts'; 


export function exportSolutionToFile(solution: Solution, filename:string = "solution.json") {
  try {
    const jsonData = JSON.stringify(solution, null, 2); 
    const blob = new Blob([jsonData], { type: 'application/json' }); 
    const url = URL.createObjectURL(blob); 

    const a = document.createElement('a');
    a.href = url; 
    a.download = filename; 
    document.body.appendChild(a);
    a.click(); 
    document.body.removeChild(a); 
    URL.revokeObjectURL(url); 
    console.log(`Solution exported to ${filename}`);
  } catch (error) {
    console.error("Error exporting Solution to file:", error);
  }
}

export async function loadSolutionFromFile(file: File): Promise<Solution> {
  const text = await file.text(); 
  const solution: Solution = JSON.parse(text); 
  return solution;
}
