"use client";

import { useState } from "react";
import { Canvas } from "@react-three/fiber";
import { OrbitControls, Text } from "@react-three/drei";
import { Button } from "@/components/ui/button";

const SmallCube = ({
  position,
  number,
}: {
  position: [number, number, number];
  number: number;
}) => {
  const [hovered, setHovered] = useState(false);

  const faces = [
    { position: [0, 0, 0.45], rotation: [0, 0, 0] },
    { position: [0, 0, -0.45], rotation: [0, Math.PI, 0] },
    { position: [0.45, 0, 0], rotation: [0, Math.PI / 2, 0] },
    { position: [-0.45, 0, 0], rotation: [0, -Math.PI / 2, 0] },
    { position: [0, 0.45, 0], rotation: [-Math.PI / 2, 0, 0] },
    { position: [0, -0.45, 0], rotation: [Math.PI / 2, 0, 0] },
  ];

  return (
    <mesh
      position={position}
      onPointerOver={() => setHovered(true)}
      onPointerOut={() => setHovered(false)}
    >
      <boxGeometry args={[0.9, 0.9, 0.9]} />
      <meshStandardMaterial color={hovered ? "hotpink" : "orange"} />
      {faces.map((face, index) => (
        <Text
          key={index}
          position={face.position as [number, number, number]}
          rotation={face.rotation as [number, number, number]}
          fontSize={0.4}
          color='#000000'
          anchorX='center'
          anchorY='middle'
        >
          {number}
        </Text>
      ))}
    </mesh>
  );
};

const NumberedCube = () => {
  const [separation, setSeparation] = useState({ x: 0, y: 0, z: 0 });

  const cubes = [];
  let number = 1;

  for (let x = 0; x < 5; x++) {
    for (let y = 0; y < 5; y++) {
      for (let z = 0; z < 5; z++) {
        cubes.push(
          <SmallCube
            key={`${x}-${y}-${z}`}
            position={
              [
                x - 2 + separation.x * x,
                y - 2 + separation.y * y,
                z - 2 + separation.z * z,
              ] as [number, number, number]
            }
            number={number++}
          />
        );
      }
    }
  }

  const handleSeparation = (axis: "x" | "y" | "z") => {
    setSeparation((prev) => ({
      ...prev,
      [axis]: prev[axis] === 0 ? 1 : 0,
    }));
  };

  return (
    <>
      <Canvas camera={{ position: [10, 10, 10], fov: 50 }}>
        <ambientLight intensity={0.5} />
        <pointLight position={[10, 10, 10]} />
        <OrbitControls />
        <group>{cubes}</group>
      </Canvas>
      <div className='absolute bottom-4 left-4 flex space-x-2'>
        <Button
          onClick={() => handleSeparation("x")}
          aria-label='Separate X axis'
        >
          Separate X
        </Button>
        <Button
          onClick={() => handleSeparation("y")}
          aria-label='Separate Y axis'
        >
          Separate Y
        </Button>
        <Button
          onClick={() => handleSeparation("z")}
          aria-label='Separate Z axis'
        >
          Separate Z
        </Button>
      </div>
    </>
  );
};

export default function Component() {
  return (
    <div className='w-full h-screen relative'>
      <NumberedCube />
    </div>
  );
}
