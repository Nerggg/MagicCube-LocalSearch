// components/Cube.js
import { useRef } from 'react';
import { useFrame, useLoader } from '@react-three/fiber';
import { TextureLoader } from 'three';
import * as THREE from 'three';

const Cube = () => {
  const cubeRef = useRef();

  // Contoh matriks angka 5x5 untuk setiap sisi
  const matrices = [
    [
      [1, 2, 3, 4, 5],
      [6, 7, 8, 9, 10],
      [11, 12, 13, 14, 15],
      [16, 17, 18, 19, 20],
      [21, 22, 23, 24, 25]
    ],
    [
      [26, 27, 28, 29, 30],
      [31, 32, 33, 34, 35],
      [36, 37, 38, 39, 40],
      [41, 42, 43, 44, 45],
      [46, 47, 48, 49, 50]
    ],
    [
        [26, 27, 28, 29, 30],
        [31, 32, 33, 34, 35],
        [36, 37, 38, 39, 40],
        [41, 42, 43, 44, 45],
        [46, 47, 48, 49, 50]
      ],
      [
        [26, 27, 28, 29, 30],
        [31, 32, 33, 34, 35],
        [36, 37, 38, 39, 40],
        [41, 42, 43, 44, 45],
        [46, 47, 48, 49, 50]
      ],
      [
        [26, 27, 28, 29, 30],
        [31, 32, 33, 34, 35],
        [36, 37, 38, 39, 40],
        [41, 42, 43, 44, 45],
        [46, 47, 48, 49, 50]
      ],
      [
        [26, 27, 28, 29, 30],
        [31, 32, 33, 34, 35],
        [36, 37, 38, 39, 40],
        [41, 42, 43, 44, 45],
        [46, 47, 48, 49, 50]
      ],
    // Tambahkan matriks lain untuk setiap sisi (total 6 matriks untuk 6 sisi kubus)
  ];

  // Fungsi untuk membuat tekstur dari matriks
  const createTextureFromMatrix = (matrix) => {
    const size = 100; // Ukuran canvas
    const canvas = document.createElement('canvas');
    canvas.width = size;
    canvas.height = size;
    const context = canvas.getContext('2d');

    // Background putih
    context.fillStyle = 'white';
    context.fillRect(0, 0, size, size);

    // Set font dan warna untuk angka
    context.fillStyle = 'black';
    context.font = '12px Arial bold';
    context.textAlign = 'center';
    context.textBaseline = 'middle';

    const cellSize = size / 5; // Ukuran tiap cell dalam matriks

    // Menggambar angka dari matriks 5x5
    for (let row = 0; row < 5; row++) {
      for (let col = 0; col < 5; col++) {
        const number = matrix[row][col];
        const x = col * cellSize + cellSize / 2;
        const y = row * cellSize + cellSize / 2;
        context.fillText(number, x, y);
      }
    }

    // Membuat tekstur dari canvas
    const texture = new THREE.CanvasTexture(canvas);
    return texture;
  };

  // Buat tekstur dari setiap matriks
  const textures = matrices.map((matrix) => createTextureFromMatrix(matrix));

  // Update rotasi kubus di setiap frame
  useFrame(() => {
    cubeRef.current.rotation.x += 0.01;
    cubeRef.current.rotation.y += 0.01;
  });

  return (
    <mesh ref={cubeRef}>
      {/* Menggunakan boxGeometry untuk membuat kubus */}
      <boxGeometry args={[3, 3, 3]} />

      {/* Terapkan material dengan tekstur yang berbeda di setiap sisi */}
      {textures.map((texture, index) => (
        <meshStandardMaterial key={index} map={texture} attach={`material-${index}`} />
      ))}
    </mesh>
  );
};

export default Cube;
