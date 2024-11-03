"use client"
import React, { useState } from 'react';
import Navbar from '../../components/navbar.js';
import { css } from "@emotion/react";
import { BeatLoader } from "react-spinners";
import axios from 'axios';
import { Canvas } from '@react-three/fiber';
import { OrbitControls } from '@react-three/drei';
import CubeMatrix from '../../components/CubeMatrix.js';


const generateRandomMatrixData = (rows, cols, depth, minValue, maxValue) => {
  const totalNumbers = rows * cols * depth;
  if (maxValue - minValue + 1 < totalNumbers) {
    throw new Error('Range too small for unique values.');
  }

  const uniqueNumbers = new Set();
  
  while (uniqueNumbers.size < totalNumbers) {
    const randomNum = Math.floor(Math.random() * (maxValue - minValue + 1)) + minValue;
    uniqueNumbers.add(randomNum);
  }

  const result = [];
  const numbersArray = Array.from(uniqueNumbers);
  
  for (let d = 0; d < depth; d++) {
    const matrix = [];
    for (let r = 0; r < rows; r++) {
      const row = [];
      for (let c = 0; c < cols; c++) {
        row.push(numbersArray.pop());
      }
      matrix.push(row);
    }
    result.push(matrix);
  }
  
  return result;
};


const override = css`
  /* Definisikan gaya khusus di sini */
`;

export default function Wikirace() {
  function delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
  const [submitted, setSubmitted] = useState(false);
  const [results, setResults] = useState(null);
  const [loading, setLoading] = useState(false);
  const [activeAlgorithm, setActiveAlgorithm] = useState(''); // default kosong
  const [errorMessage, setErrorMessage] = useState(null)
  const [separateX, setSeparateX] = useState(1.1);
  const [separateY, setSeparateY] = useState(1.1);
  const [separateZ, setSeparateZ] = useState(1.1);
  const [matrixData, setMatrixData] = useState(generateRandomMatrixData(5, 5, 5, 1, 125));

  const handleGenerateRandom = () => {
    setMatrixData(generateRandomMatrixData(5, 5, 5, 1, 125)); // Generate new random data
    console.log("Matrix data", matrixData);
  };

  const handleAlgorithmClick = (algorithm) => {
    console.log("Algorithm", algorithm);
    setActiveAlgorithm(algorithm);
  };

  const handleSubmit = async (event) => {
    setSubmitted(false)
    event.preventDefault();

    console.log("Handle Submit triggered")
    
    if (activeAlgorithm == ''){
      setErrorMessage("Please choose the algorithm.");
      setLoading(false)
      await delay(1500);
      setErrorMessage(null);
      return;
    }
    setLoading(true);
    try {
      // TODO: frontend request to backend server endpoint
      const response = await axios.post('http://localhost:8080/search', {
        cube : matrixData,
        algorithm: activeAlgorithm
      });
      setLoading(false);
      setSubmitted(true);
      setResults(response.data);
      setErrorMessage(null)
    } catch (error) {
      console.error(error);
      setLoading(false);
    }
  };


// Button style base
const baseStyle = "mx-4 rounded border-2 px-7 pb-[8px] pt-[10px] text-sm font-medium uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0";
const dynamicStyle = (isActive) => 
  `${baseStyle} ${isActive ? 'border-neutral-300 text-gray-800 bg-neutral-100' : 'border-neutral-200 text-gray-600 hover:border-neutral-300 hover:bg-neutral-200 hover:text-gray-800'}`;

return (
  
  <div className="bg-gray-100 text-gray-800 font-sans pb-12 min-h-screen"> {/* Updated background */}
  <Navbar />    
  <div className="flex flex-col h-[60vh] w-[70vw] items-center mt-[20px] mx-auto mb-12 relative"> {/* Make the parent div relative */}
  {/* Canvas Container with Border */}
  <div className="border-2 border-gray-300 rounded-lg p-4 mb-[30px] w-full h-full relative"> {/* Border and padding for canvas container */}
    <Canvas>
      <ambientLight intensity={0.5} />
      <directionalLight position={[5, 5, 5]} intensity={1} />
      <CubeMatrix data={matrixData} separateX={separateX} separateY={separateY} separateZ={separateZ} rotationSpeed={0} />
      <OrbitControls enableRotate={true} enablePan={true} enableZoom={true} />
    </Canvas>
    
    {/* Instruction Box */}
    <div className="absolute bottom-4 left-4 bg-white bg-opacity-75 border border-gray-300 rounded-lg p-2 shadow-md">
      <p className="text-gray-800">Drag to pan. Scroll to zoom.</p>
    </div>

        {/* Slider Controls */}
        <div className="absolute bottom-4 right-4 bg-white bg-opacity-75 border border-gray-300 rounded-lg p-4 shadow-md">
      <div>
        <label className="block">Separate by X:</label>
        <input
          type="range"
          min="1.0"
          max="2.5"
          step="0.1"
          value={separateX}
          onChange={(e) => setSeparateX(parseFloat(e.target.value))}
          className="w-full h-2 bg-gray-200 rounded-full"
          style={{
            accentColor: "#1F2937", // Tailwind color gray-800 for thumb
            color: "#1F2937",
          }}
        />
      </div>
      <div>
        <label className="block">Separate by Y:</label>
        <input
          type="range"
          min="1.0"
          max="2.5"
          step="0.1"
          value={separateY}
          onChange={(e) => setSeparateY(parseFloat(e.target.value))}
          className="w-full h-2 bg-gray-200 rounded-full"
          style={{
            accentColor: "#1F2937", // Tailwind color gray-800 for thumb
            color: "#1F2937",
          }}
        />
      </div>
      <div>
        <label className="block">Separate by Z:</label>
        <input
          type="range"
          min="1.0"
          max="2.5"
          step="0.1"
          value={separateZ}
          onChange={(e) => setSeparateZ(parseFloat(e.target.value))}
          className="w-full h-2 bg-gray-200 rounded-full"
          style={{
            accentColor: "#1F2937", // Tailwind color gray-800 for thumb
            color: "#1F2937",
          }}
        />
      </div>
    </div>
  </div>
  
  {/* Button to generate random input */}
  <button 
    onClick={handleGenerateRandom}
    className={dynamicStyle(false)}
  >
    Generate Random Input
  </button>
</div>
    <div className="text-gray-800 justify-center"> {/* Updated text color */}
      <form onSubmit={handleSubmit}>
        <div className='flex flex-col items-center'>
          <h4 className="mb-2 text-xl font-semibold text-gray-800">Algorithm Type</h4>
          <div>
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Hill Climbing')}
              onClick={() => handleAlgorithmClick('Hill Climbing')}
              data-twe-ripple-init
              data-twe-ripple-color="light"
            >
              Hill Climbing
            </button>
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Simulated Annealing')}
              onClick={() => handleAlgorithmClick('Simulated Annealing')}
              data-twe-ripple-init
              data-twe-ripple-color="light"
            >
              Simulated Annealing
            </button>
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Genetic Algorithm')}
              onClick={() => handleAlgorithmClick('Genetic Algorithm')}
              data-twe-ripple-init
              data-twe-ripple-color="light"
            >
              Genetic Algorithm
            </button>
          </div>
          {!loading && (<button type="submit"
            className='mt-4 mx-4 mb-[15px] rounded border-2 border-gray-200 px-7 pb-[8px] pt-[10px] text-sm font-bold uppercase leading-normal text-gray-800 transition duration-150 ease-in-out hover:border-neutral-300 hover:bg-neutral-200 focus:border-neutral-300 focus:text-gray-800 focus:outline-none focus:ring-0 active:border-neutral-200 active:text-gray-600'
            data-twe-ripple-init
            data-twe-ripple-color="light">
              Submit!
          </button>
          )}
        </div>
      </form>
      {errorMessage && (
        <div className='text-gray-800 text-center mb-4'>{errorMessage}</div> // Updated color
      )}
      {loading && (
        <div className="flex justify-center items-center mt-[25px] mb-[50px]">
          <BeatLoader color="#000000" loading={loading} css={override} size={15} />
          <p className="ml-2 text-gray-800">Loading...</p>
        </div>
      )}
      {submitted && results && (
        <div className='flex flex-col items-center'>
          {results.numberOfPaths === 0 ? (
            <p className="text-gray-800 text-center mt-4 text-xl">No path found from <strong>{awal}</strong> to <strong>{akhir}</strong></p>
          ) : (
            <>
              <div className='w-[50%]'>
                <p className="text-gray-800 text-center mt-4 text-xl">Found <strong>{results.numberOfPaths} paths</strong> from <strong>{awal}</strong> to <strong>{akhir}</strong> in <strong>{results.elapsedTime} seconds</strong>!</p>
                <p className="text-gray-800 text-center mt-4 text-xl">Articles Checked: <strong>{results.articlesChecked}</strong></p>
                <p className="text-gray-800 text-center mt-4 text-xl">Articles Traversed: <strong>{results.articlesTraversed}</strong></p>
              </div>
              <div className='w-[1140px] bg-gray-300 h-[2px] mt-2'/> {/* Updated color */}
              <h2 className='mt-5 text-2xl font-bold text-gray-800'> Connecting Graphs </h2>
              <div className='w-[900px] h-[450px] font-inter rounded-[10px] border-2 border-gray-300 mr-2 overflow-hidden'> {/* Updated border */}
                  <div className='translate-x-[-200px] translate-y-[100px] z-[-10px]'>
                      <Graph path={results.paths} />
                  </div>
                  <div className='flex translate-y-[-760px] w-[150px] z-[10px] h-fit rounded-[10px] border-2 border-gray-300 mt-2 ml-2'> {/* Updated border */}
                      <p className="text-gray-800">Drag to pan. Scroll to zoom.</p> {/* Updated text color */}
                  </div>
              </div>
              <div className='w-[1140px] bg-gray-300 h-[2px] mt-4'/> {/* Updated color */}
              <h2 className='mt-5 text-2xl font-bold text-gray-800'> Individual Paths </h2>
              <div className="w-full flex flex-col items-center justify-center">
                  <PathBox path={results.paths} />
              </div>
          </>
          )}
      </div>
    )}
  </div>
</div>
);
}

