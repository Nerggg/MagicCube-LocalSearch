// components/CubeMatrix.js
import { useRef } from 'react';
import * as THREE from 'three';

const CubeMatrix = ({ data, separateX, separateY, separateZ }) => {
  const groupRef = useRef();

  // Fungsi untuk membuat tekstur dari angka
  const createNumberTexture = (number) => {
    const size = 128; // Ukuran canvas
    const canvas = document.createElement('canvas');
    canvas.width = size;
    canvas.height = size;
    const context = canvas.getContext('2d');

    // Background putih
    context.fillStyle = 'white';
    context.fillRect(0, 0, size, size);

    // Set font dan warna untuk angka
    context.fillStyle = 'black';
    context.font = '60px Arial';
    context.textAlign = 'center';
    context.textBaseline = 'middle';

    // Tampilkan angka di tengah canvas
    context.fillText(number, size / 2, size / 2);

    // Membuat tekstur dari canvas
    const texture = new THREE.CanvasTexture(canvas);
    return texture;
  };

  // Ukuran kubus kecil
  const cubeSize = 0.5;

  // Offset dihitung dengan mempertimbangkan ukuran dan pemisahan
  const offsetX = (data[0][0].length - 1) * cubeSize * separateX * 0.5;
  const offsetY = (data[0].length - 1) * cubeSize * separateY * 0.5;
  const offsetZ = (data.length - 1) * cubeSize * separateZ * 0.5;

  return (
    <group ref={groupRef}>
      {data.map((layer, zIndex) =>
        layer.map((row, yIndex) =>
          row.map((number, xIndex) => {
            const textures = Array(6).fill(createNumberTexture(number));

            // Menggunakan offset yang telah disesuaikan untuk membuat poros tetap di tengah
            const position = [
              (xIndex * cubeSize * separateX) - offsetX,
              (yIndex * cubeSize * separateY) - offsetY,
              (zIndex * cubeSize * separateZ) - offsetZ,
            ];

            return (
              <mesh position={position} key={`${xIndex}-${yIndex}-${zIndex}`}>
                <boxGeometry args={[cubeSize, cubeSize, cubeSize]} />
                {textures.map((texture, index) => (
                  <meshStandardMaterial key={index} map={texture} attach={`material-${index}`} />
                ))}
              </mesh>
            );
          })
        )
      )}
    </group>
  );
};

export default CubeMatrix;
