// components/CubeMatrix.js
import { useRef } from 'react';
import { useFrame } from '@react-three/fiber';
import * as THREE from 'three';

const CubeMatrix = ({ data, separateX, separateY, separateZ, rotationSpeed, visibleLevel }) => {
  const groupRef = useRef();

  const createNumberTexture = (number) => {
    const size = 128;
    const canvas = document.createElement('canvas');
    canvas.width = size;
    canvas.height = size;
    const context = canvas.getContext('2d');

    context.fillStyle = 'white';
    context.fillRect(0, 0, size, size);

    context.fillStyle = 'black';
    context.font = '60px Arial';
    context.textAlign = 'center';
    context.textBaseline = 'middle';

    context.fillText(number, size / 2, size / 2);

    const texture = new THREE.CanvasTexture(canvas);
    return texture;
  };

  const cubeSize = 0.5;

  const offsetX = (data[0][0].length - 1) * cubeSize * separateX * 0.5;
  const offsetY = (data[0].length - 1) * cubeSize * separateY * 0.5;
  const offsetZ = (data.length - 1) * cubeSize * separateZ * 0.5;

  useFrame(() => {
    if (groupRef.current) {
      groupRef.current.rotation.x += rotationSpeed;
      groupRef.current.rotation.y += rotationSpeed;
    }
  });

  return (
    <group ref={groupRef}>
      {data.map((layer, zIndex) =>
        zIndex >= visibleLevel[0] && zIndex <= visibleLevel[1] ? ( // Check if layer is within visible range
          layer.map((row, yIndex) =>
            row.map((number, xIndex) => {
              const textures = Array(6).fill(createNumberTexture(number));

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
        ) : null
      )}
    </group>
  );
};

export default CubeMatrix;
