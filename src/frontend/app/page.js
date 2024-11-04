"use client"
import React from 'react';
import { useState } from 'react';
import Navbar from '../components/navbar.js';
import { Canvas } from '@react-three/fiber';
import { OrbitControls } from '@react-three/drei';
import CubeMatrix from '../components/CubeMatrix.js';

export default function Home() {
  const matrixData = [
    [
      [25, 16, 80, 104, 90],
      [115, 98, 4, 1, 97],
      [42, 111, 85, 2, 75],
      [66, 72, 27, 102, 48],
      [67, 18, 119, 106, 5],
    ],
    [
      [91, 77, 71, 6, 70],
      [52, 64, 117, 69, 13],
      [30, 118, 21, 123, 23],
      [26, 39, 92, 44, 114],
      [116, 17, 14, 73, 95],
    ],
    [
      [47, 61, 45, 76, 86],
      [107, 43, 38, 33, 94],
      [89, 68, 63, 58, 37],
      [32, 93, 88, 83, 19],
      [40, 50, 81, 65, 79],
    ],
    [
      [31, 53, 112, 109, 10],
      [12, 82, 34, 87, 100],
      [103, 3, 105, 8, 96],
      [113, 57, 9, 62, 74],
      [56, 120, 55, 49, 35],
    ],
    [
      [121, 7, 108, 20, 59],
      [29, 28, 122, 125, 11],
      [51, 15, 41, 124, 84],
      [78, 54, 99, 24, 60],
      [36, 110, 46, 22, 101],
    ],
  ];

  const [separateX, setSeparateX] = useState(1.1);
  const [separateY, setSeparateY] = useState(1.1);
  const [separateZ, setSeparateZ] = useState(1.1);




  return (
    <div className="bg-gray-100 text-gray-800 font-sans pb-12 min-h-screen">
      <Navbar />

      {/* Canvas Section */}
      <div className="card h-[70vh] w-[80vw] mx-auto mb-12">
        <Canvas>
          <ambientLight intensity={0.5} />
          <directionalLight position={[5, 5, 5]} intensity={1} />
          <CubeMatrix data={matrixData} separateX={separateX} separateY={separateY} separateZ={separateZ} rotationSpeed={0.01} visibleLevel={[0,4]} />
          <OrbitControls enableRotate={false} enablePan={false} enableZoom={false} />
        </Canvas>
      </div>

      <div className="flex justify-center pb-[60px] ">
      <div align="center h-[50px]">
        <h1 className="text-5xl font-bold typing-animation"><u>Magic Cube</u> Solver</h1>
      </div>

      </div>

      {/* MagicCube Information Card */}
      <div className="card bg-white border border-gray-300 rounded-lg shadow-lg p-6 w-[70%] mx-auto mb-12">
        <h3 className="text-3xl font-bold text-center mb-4">
          What is <span className="underline">Magic Cube</span>?
        </h3>
        <div className="flex flex-col md:flex-row items-center justify-center">
          <img src={'/Simple_Magic_Cube.svg'} alt="Wikipedia logo" className="w-40 rounded-lg mb-6 md:mb-0 md:mr-8" />
          <p className="text-gray-700 text-justify leading-relaxed">
          A <b>Magic Cube</b> is a 3D extension of the magic square in mathematics. It is a cube grid (like 3x3x3 or larger) filled with numbers so that the sum of each row, column, and layer is equal, achieving a "magic constant." Variants include perfect magic cubes, where all diagonals also sum to this constant. Magic cubes are studied in number theory and combinatorics due to their symmetrical and numerical patterns.
          </p>
        </div>
      </div>

      {/* BFS and IDS Algorithm Explanation Cards */}
      <div className="flex flex-col md:flex-row justify-around w-full gap-6">

{/* Hill Climbing Card */}
<div  className="card bg-white border border-gray-300 rounded-lg shadow-lg p-6 w-full md:w-[30%] mx-auto">
  <h2 className="text-2xl font-bold mb-4">Hill Climbing</h2>
  <p className="text-gray-700 text-justify leading-relaxed mb-6">
  Hill Climbing is an optimization algorithm used to find the best solution within a search space by iteratively moving toward improvements. The algorithm starts with an initial or random solution, then evaluates neighboring solutions, moving towards the direction that increases the objective value locally. If a better neighboring solution is found, the algorithm moves to that solution. This process continues until it reaches a local or global peak, where no better neighbors are available. Although efficient, Hill Climbing can get stuck at local maxima and often does not guarantee a globally optimal solution.
</p>
  <div className="flex justify-center">
    <img src="/Hill_climb.png" alt="Hill Climb Algorithm" className="h-60 rounded-lg" />
  </div>
</div>

{/* Simulated Annealing Card */}
<div className="card bg-white border border-gray-300 rounded-lg shadow-lg p-6 w-full md:w-[30%] mx-auto">
  <h2 className="text-2xl font-bold mb-4">Simulated Annealing</h2>
  <p className="text-gray-700 text-justify leading-relaxed mb-6">
    Simulated Annealing is an optimization algorithm inspired by the annealing process in metallurgy, where materials are heated and then slowly cooled to reach a stable state. The algorithm explores the search space by randomly selecting neighboring solutions and deciding whether to move towards them based on a probability that decreases over time. Initially, it allows moves to worse solutions to escape local optima, but as the "temperature" lowers, the algorithm becomes more selective, converging towards an optimal or near-optimal solution. This makes Simulated Annealing effective for complex problems where finding a global optimum is challenging.
  </p>
  <div className="flex justify-center">
    <img src="/simulated_annealing.gif" alt="SA Algorithm" className="h-80 rounded-lg" />
  </div>
</div>

{/* Genetic Algorithm Card */}
<div className="card bg-white border border-gray-300 rounded-lg shadow-lg p-6 w-full md:w-[30%] mx-auto">
  <h2 className="text-2xl font-bold mb-4">Genetic Algorithm</h2>
  <p className="text-gray-700 text-justify leading-relaxed mb-6">
    A Genetic Algorithm (GA) is an optimization and search technique inspired by the principles of natural selection and genetics. It operates by creating a population of potential solutions and evolving them over generations to improve performance. Each generation involves processes like selection, crossover (recombination), and mutation to produce new solutions. The fittest solutions are more likely to be selected and passed on to the next generation, allowing the algorithm to explore and exploit the search space effectively. GAs are particularly useful for solving complex problems where traditional optimization methods struggle to find global optima.
  </p>
  <div className="flex justify-center">
    <img src="/genetic_algorithm.png" alt="Genetic Algorithm" className="h-80 rounded-lg" />
  </div>
</div>
</div>

    </div>
  );
}