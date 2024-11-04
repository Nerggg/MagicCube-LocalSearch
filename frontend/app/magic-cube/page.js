"use client"
import React, { useState } from 'react';
import Navbar from '../../components/navbar.js';
import { css } from "@emotion/react";
import { BeatLoader } from "react-spinners";
import axios from 'axios';
import { Canvas } from '@react-three/fiber';
import { OrbitControls } from '@react-three/drei';
import CubeMatrix from '../../components/CubeMatrix.js';
import ObjectiveChart from '@/components/ObjectiveChart.js';

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

export default function Magiccube() {
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
  const [separateXresult, setSeparateXresult] = useState(1.1);
  const [separateYresult, setSeparateYresult] = useState(1.1);
  const [separateZresult, setSeparateZresult] = useState(1.1);
  const [matrixData, setMatrixData] = useState(generateRandomMatrixData(5, 5, 5, 1, 125));
  const [populationSize, setPopulationSize] = useState(0);
  const [maxGenerations, setMaxGenerations] = useState(0);
  const [maxStateGeneration, setMaxStateGeneration] = useState(0);
  const [restartChance, setRestartChance] = useState(0);
  const [restartAmount, setRestartAmount] = useState(0);
  const [temperature, setTemperature] = useState(0); 
  const [coolingRate, setCoolingRate] = useState(0); 
  const [maxIterations, setMaxIterations] = useState(0); 
  const [visibleLevel, setVisibleLevel] = useState([0, matrixData.length - 1]);
  const [visibleLevelresult, setVisibleLevelresult] = useState([0, 4]);

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

    // Error handling based on selected algorithm
    if (activeAlgorithm === 'Genetic Algorithm') {
      if (!populationSize || populationSize <= 0) {
        setErrorMessage("Please enter a valid Population Size.");
        setLoading(false)
        await delay(1500);
        setErrorMessage(null);
        return;
      }
      if (!maxGenerations || maxGenerations <= 0) {
        setErrorMessage("Please enter a valid Max Generations.");
        setLoading(false)
        await delay(1500);
        setErrorMessage(null);
        return;
      }
    }

    if (activeAlgorithm === 'Stochastic Hill Climbing') {
      if (!maxStateGeneration || maxStateGeneration <= 0) {
        setErrorMessage("Please enter a valid Max State Generation.");
        setLoading(false)
        await delay(1500);
        setErrorMessage(null);
        return;
      }
    }

    if (activeAlgorithm === 'Random Restart Hill Climbing') {
      if (!restartChance || restartChance <= 0 || restartChance > 100) {
        setErrorMessage("Please enter a valid Restart Chance (between 1 and 100).");
        setLoading(false)
        await delay(1500);
        setErrorMessage(null);
        return;
      }
      if (!restartAmount || restartAmount <= 0) {
        setErrorMessage("Please enter a valid Maximum Restart.");
        setLoading(false)
        await delay(1500);
        setErrorMessage(null);
        return;
      }
    }

    if (activeAlgorithm === 'Simulated Annealing') {
      if (!temperature || temperature <= 0) {
        setErrorMessage("Please enter a valid Temperature.");
        setLoading(false)
        await delay(1500);
        setErrorMessage(null);
        return;
      }
      if (coolingRate == null || coolingRate <= 0 || coolingRate >= 1) {
        setErrorMessage("Please enter a valid Cooling Rate (between 0 and 1).");
        setLoading(false)
        await delay(1500);
        setErrorMessage(null);
        return;
      }
      if (!maxIterations || maxIterations <= 0) {
        setErrorMessage("Please enter a valid Max Iterations.");
        setLoading(false)
        await delay(1500);
        setErrorMessage(null);
        return;
      }
    }

    setLoading(true);
    try {
      // TODO: frontend request to backend server endpoint
      const response = await axios.post('http://localhost:8080/search', {
        cube : matrixData,
        algorithm: activeAlgorithm,
        populationSize: activeAlgorithm === 'Genetic Algorithm' ? populationSize : undefined,
        maxGenerations: activeAlgorithm === 'Genetic Algorithm' ? maxGenerations : undefined,
        temperature: activeAlgorithm === 'Simulated Annealing' ? temperature : undefined,
        maxStateGeneration: activeAlgorithm === 'Stochastic Hill Climbing' ? maxStateGeneration : undefined,
        restartChance: activeAlgorithm === 'Random Restart Hill Climbing' ? restartChance : undefined,
        restartAmount: activeAlgorithm === 'Random Restart Hill Climbing' ? restartAmount : undefined,
        coolingRate: activeAlgorithm === 'Simulated Annealing' ? coolingRate : undefined,
        maxIterations: activeAlgorithm === 'Simulated Annealing' ? maxIterations : undefined
      });
      setLoading(false);
      setSubmitted(true);
      console.log("Response", response.data);
      setResults(response.data);
      console.log("Matrix Data", matrixData);
      setErrorMessage(null)
    } catch (error) {
      console.error(error);
      setLoading(false);
    }
  };


// Button style base
const baseStyle = "mx-4 rounded border-2 px-7 pb-[8px] pt-[10px] text-sm font-medium uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0 bg-gray-200 text-gray-700 border-gray-300";
const dynamicStyle = (isActive) => 
  `${baseStyle} ${isActive ? 'border-gray-400 text-gray-900 bg-gray-300' : 'hover:border-gray-400 hover:bg-gray-300 hover:text-gray-900'}`;

return (
  <div className="bg-gray-100 text-gray-800 font-sans pb-12 min-h-screen">
    <Navbar />    
    <div className="flex flex-col h-[60vh] w-[70vw] items-center mt-[20px] mx-auto mb-12 relative">
      {/* Canvas Container with Border */}
      <div className="border-2 border-gray-300 rounded-lg p-4 mb-[30px] w-full h-full relative">
        <Canvas>
          <ambientLight intensity={0.5} />
          <directionalLight position={[5, 5, 5]} intensity={1} />
          <CubeMatrix
            data={matrixData}
            separateX={separateX}
            separateY={separateY}
            separateZ={separateZ}
            rotationSpeed={0}
            visibleLevel={visibleLevel}
          />
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
              style={{ accentColor: "#1F2937", color: "#1F2937" }}
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
              style={{ accentColor: "#1F2937", color: "#1F2937" }}
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
              style={{ accentColor: "#1F2937", color: "#1F2937" }}
            />
          </div>
          <div>
            <label className="block">Select Z Level:</label>
            <input
              type="range"
              min="0"
              max={matrixData.length - 1}
              step="1"
              value={visibleLevel[1]}
              onChange={(e) => setVisibleLevel([0, parseInt(e.target.value)])}
              className="w-full h-2 bg-gray-200 rounded-full"
              style={{ accentColor: "#1F2937", color: "#1F2937" }}
            />
            <p>Current Level: {visibleLevel[1]}</p>
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

    <div className="text-gray-800 justify-center mt-[20px]">
      <form onSubmit={handleSubmit}>
        <div className='flex flex-col items-center'>
          <h4 className="mb-2 text-xl font-semibold text-gray-800">Algorithm Type</h4>
          <div className="flex flex-wrap justify-center gap-4 mt-4">
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Steepest Ascent Hill Climbing')}
              onClick={() => handleAlgorithmClick('Steepest Ascent Hill Climbing')}
            >
              Steepest Ascent Hill Climbing
            </button>
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Sideways Move Hill Climbing')}
              onClick={() => handleAlgorithmClick('Sideways Move Hill Climbing')}
            >
              Sideways Move Hill Climbing
            </button>
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Random Restart Hill Climbing')}
              onClick={() => handleAlgorithmClick('Random Restart Hill Climbing')}
            >
              Random Restart Hill Climbing
            </button>
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Stochastic Hill Climbing')}
              onClick={() => handleAlgorithmClick('Stochastic Hill Climbing')}
            >
              Stochastic Hill Climbing
            </button>
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Simulated Annealing')}
              onClick={() => handleAlgorithmClick('Simulated Annealing')}
            >
              Simulated Annealing
            </button>
            <button
              type="button"
              className={dynamicStyle(activeAlgorithm === 'Genetic Algorithm')}
              onClick={() => handleAlgorithmClick('Genetic Algorithm')}
            >
              Genetic Algorithm
            </button>
          </div>
          {activeAlgorithm === 'Genetic Algorithm' && (
            <div className="mt-4">
              <label className="block mb-2 text-gray-800">Population Size:</label>
              <input
                type="number"
                value={populationSize}
                onChange={(e) => setPopulationSize(parseInt(e.target.value))}
                className="w-full h-10 px-3 text-base placeholder-gray-600 border rounded-lg focus:shadow-outline"
              />
              <label className="block mt-4 mb-2 text-gray-800">Max Generations:</label>
              <input
                type="number"
                value={maxGenerations}
                onChange={(e) => setMaxGenerations(parseInt(e.target.value))}
                className="w-full h-10 px-3 text-base placeholder-gray-600 border rounded-lg focus:shadow-outline"
              />
            </div>
          )}
          {activeAlgorithm === 'Stochastic Hill Climbing' && (
            <div className="mt-4">
              <label className="block mb-2 text-gray-800">Max State Generation:</label>
              <input
                type="number"
                value={maxStateGeneration}
                onChange={(e) => setMaxStateGeneration(parseInt(e.target.value))}
                className="w-full h-10 px-3 text-base placeholder-gray-600 border rounded-lg focus:shadow-outline"
              />
            </div>
          )}
          {activeAlgorithm === 'Random Restart Hill Climbing' && (
            <div className="mt-4">
              <label className="block mb-2 text-gray-800">Restart Chance (in %):</label>
              <input
                type="number"
                value={restartChance}
                onChange={(e) => setRestartChance(parseFloat(e.target.value))}
                className="w-full h-10 px-3 text-base placeholder-gray-600 border rounded-lg focus:shadow-outline"
              />
              <label className="block mb-2 text-gray-800">Maximum Restart:</label>
              <input
                type="number"
                value={restartAmount}
                onChange={(e) => setRestartAmount(parseInt(e.target.value))}
                className="w-full h-10 px-3 text-base placeholder-gray-600 border rounded-lg focus:shadow-outline"
              />
            </div>
          )}
          {activeAlgorithm === 'Simulated Annealing' && (
            <div className="mt-4">
              <label className="block text-gray-800">Temperature:</label>
              <input
                type="number"
                value={temperature}
                onChange={(e) => setTemperature(parseInt(e.target.value))}
                className="w-full h-10 px-3 text-base placeholder-gray-600 border rounded-lg focus:shadow-outline"
              />
              <label className="block text-gray-800 mt-4">Cooling Rate:</label>
              <input
                type="number"
                step="0.001"
                value={coolingRate}
                onChange={(e) => setCoolingRate(parseFloat(e.target.value))}
                className="w-full h-10 px-3 text-base placeholder-gray-600 border rounded-lg focus:shadow-outline"
              />
              <label className="block text-gray-800 mt-4">Max Iterations:</label>
              <input
                type="number"
                value={maxIterations}
                onChange={(e) => setMaxIterations(parseInt(e.target.value))}
                className="w-full h-10 px-3 text-base placeholder-gray-600 border rounded-lg focus:shadow-outline"
              />
            </div>
          )}
          {!loading && (
          <button type="submit"
              className="mt-4 mx-4 mb-[15px] rounded border-2 border-neutral-300 px-7 pb-[8px] pt-[10px] text-sm font-bold uppercase leading-normal text-gray-700 bg-neutral-200 transition duration-150 ease-in-out 
              hover:border-neutral-400 hover:bg-neutral-300 hover:text-gray-900 
              focus:border-neutral-400 focus:bg-neutral-300 focus:text-gray-900 focus:outline-none focus:ring-0 
              active:border-neutral-500 active:bg-neutral-400 active:text-gray-80"
          >
              Submit!
          </button>
          )}
        </div>
      </form>
      {errorMessage && (
        <div className='text-gray-800 text-center mb-4'>{errorMessage}</div>
      )}
      {loading && (
        <div className="flex justify-center items-center mt-[25px] mb-[50px]">
          <BeatLoader color="#000000" loading={loading} css={override} size={15} />
          <p className="ml-2 text-gray-800">Loading...</p>
        </div>
      )}
      {submitted && results && (
        <div className="flex flex-col items-center">
          <div className="w-[1140px] bg-gray-300 h-[2px] mt-4" />
          <div className="flex flex-col h-[60vh] w-[70vw] items-center mt-[20px] mx-auto mb-12 relative">
          <h1 className="text-3xl font-bold text-gray-800 underline mb-4">RESULTS</h1>
            {/* Canvas Container with Border */}
            <div className="border-2 border-gray-300 rounded-lg p-4 mb-[30px] w-full h-full relative">
              <Canvas>
                <ambientLight intensity={0.5} />
                <directionalLight position={[5, 5, 5]} intensity={1} />
                <CubeMatrix
                  data={results.finalState}
                  separateX={separateXresult}
                  separateY={separateYresult}
                  separateZ={separateZresult}
                  visibleLevel={visibleLevelresult}
                  rotationSpeed={0}
                />
                <OrbitControls enableRotate={true} enablePan={true} enableZoom={true} />
              </Canvas>
              <div className="absolute bottom-4 left-4 bg-white bg-opacity-75 border border-gray-300 rounded-lg p-2 shadow-md">
                <p className="text-gray-800">Drag to pan. Scroll to zoom.</p>
              </div>
              <div className="absolute bottom-4 right-4 bg-white bg-opacity-75 border border-gray-300 rounded-lg p-4 shadow-md">
              <div>
                <label className="block">Separate by X:</label>
                <input
                  type="range"
                  min="1.0"
                  max="2.5"
                  step="0.1"
                  value={separateXresult}
                  onChange={(e) => setSeparateXresult(parseFloat(e.target.value))}
                  className="w-full h-2 bg-gray-200 rounded-full"
                  style={{ accentColor: "#1F2937", color: "#1F2937" }}
                />
              </div>
              <div>
                <label className="block">Separate by Y:</label>
                <input
                  type="range"
                  min="1.0"
                  max="2.5"
                  step="0.1"
                  value={separateYresult}
                  onChange={(e) => setSeparateYresult(parseFloat(e.target.value))}
                  className="w-full h-2 bg-gray-200 rounded-full"
                  style={{ accentColor: "#1F2937", color: "#1F2937" }}
                />
              </div>
              <div>
                <label className="block">Separate by Z:</label>
                <input
                  type="range"
                  min="1.0"
                  max="2.5"
                  step="0.1"
                  value={separateZresult}
                  onChange={(e) => setSeparateZresult(parseFloat(e.target.value))}
                  className="w-full h-2 bg-gray-200 rounded-full"
                  style={{ accentColor: "#1F2937", color: "#1F2937" }}
                />
              </div>
              <div>
                <label className="block">Select Z Level:</label>
                <input
                  type="range"
                  min="0"
                  max={matrixData.length - 1}
                  step="1"
                  value={visibleLevelresult[1]}
                  onChange={(e) => setVisibleLevelresult([0, parseInt(e.target.value)])}
                  className="w-full h-2 bg-gray-200 rounded-full"
                  style={{ accentColor: "#1F2937", color: "#1F2937" }}
                />
                <p>Current Level: {visibleLevelresult[1]}</p>
              </div>
            </div>
            </div>
          </div>
          <p className="text-gray-800 text-center my-4 text-xl">Found <strong>{activeAlgorithm}</strong> solution in <strong>{(results.duration / 1000).toFixed(2)} seconds</strong>!</p>
          <p className="text-gray-800 text-center mb-4 text-xl">Final State Objective Value: <strong>{results.finalValue}</strong></p>
          {(activeAlgorithm === 'Steepest Ascent Hill Climbing' || activeAlgorithm === 'Stochastic Hill Climbing' || activeAlgorithm === 'Sideways Move Hill Climbing' || activeAlgorithm === 'Genetic Algorithm') &&(
            <p className="text-gray-800 text-center mb-4 text-xl">Iterations needed: <strong>{results.iterOF.length}</strong></p>
          )}
           {(activeAlgorithm === 'Genetic Algorithm') &&(
            <p className="text-gray-800 text-center mb-4 text-xl">Population Size: <strong>{results.populationSize}</strong></p>
          )}
           {(activeAlgorithm === 'Simulated Annealing') &&(
            <p className="text-gray-800 text-center mb-4 text-xl">Stucks Count: <strong>{results.stuckCount}</strong></p>
          )}
          <ObjectiveChart iterOF={results.iterOF} />
        </div>
      )}
    </div>
  </div>
);}
