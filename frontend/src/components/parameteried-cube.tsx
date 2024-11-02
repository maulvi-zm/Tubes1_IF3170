import { useState } from "react";
import { Canvas } from "@react-three/fiber";
import { OrbitControls, Text } from "@react-three/drei";
import { Button } from "@/components/ui/button";
import Draggable from "react-draggable";
import { Gamepad } from "lucide-react";

const SmallCube = ({
  position,
  rotation,
  number,
}: {
  position: [number, number, number];
  rotation?: [number, number, number];
  number: number;
}) => {
  const [hovered, setHovered] = useState(false);

  const faces = [
    { position: [0, 0, 0.501], rotation: [0, 0, 0] },
    { position: [0, 0, -0.501], rotation: [0, Math.PI, 0] },
    { position: [0.501, 0, 0], rotation: [0, Math.PI / 2, 0] },
    { position: [-0.501, 0, 0], rotation: [0, -Math.PI / 2, 0] },
    { position: [0, 0.501, 0], rotation: [-Math.PI / 2, 0, 0] },
    { position: [0, -0.501, 0], rotation: [Math.PI / 2, 0, 0] },
  ];

  return (
    <mesh
      position={position}
      rotation={rotation}
      onPointerOver={() => setHovered(true)}
      onPointerOut={() => setHovered(false)}
      castShadow
      receiveShadow
    >
      <boxGeometry args={[1, 1, 1]} />
      <meshStandardMaterial
        color={hovered ? "#f0f0f0" : "#ffffff"}
        metalness={0.5}
        roughness={0.9} // Adjust for a smoother surface
      />
      {faces.map((face, index) => (
        <Text
          key={index}
          position={face.position as [number, number, number]}
          rotation={face.rotation as [number, number, number]}
          fontSize={0.3}
          color="#000000"
          anchorX="center"
          anchorY="middle"
          renderOrder={1}
        >
          {number}
        </Text>
      ))}
    </mesh>
  );
};

const ParameterizedCube = ({
  cameraPosition,
  moveCamera,
  state,
}: {
  cameraPosition: [number, number, number];
  moveCamera: (axis: "x" | "y" | "z", direction: number) => void;
  state:number[];
}) => {
  const [activeAxis, setActiveAxis] = useState<"none" | "x" | "y" | "z">(
    "none",
  );
  const [flattenAxis, setFlattenAxis] = useState<"none" | "x" | "y" | "z">(
    "none",
  );
  const separation = 2;
  const flattenSeparation = 5;
  const baseHeight = 3;

  const cubes = [];
  let number = 1;

  const getPositionWithOffset = (x: number, y: number, z: number) => {
    const relativeX = x - 2;
    const relativeY = y - 2;
    const relativeZ = z - 2;

    let xPos = relativeX;
    let yPos = relativeY;
    let zPos = relativeZ;

    switch (activeAxis) {
      case "x":
        xPos =
          relativeX + Math.sign(relativeX) * separation * Math.abs(relativeX);
        break;
      case "y":
        yPos =
          relativeY + Math.sign(relativeY) * separation * Math.abs(relativeY);
        break;
      case "z":
        zPos =
          relativeZ + Math.sign(relativeZ) * separation * Math.abs(relativeZ);
        break;
    }

    switch (flattenAxis) {
      case "x":
        xPos = 0;
        yPos = relativeY;
        zPos = relativeZ + (x - 2) * flattenSeparation;
        break;
      case "y":
        xPos = relativeX;
        yPos = 0;
        zPos = relativeZ + (y - 2) * flattenSeparation;
        break;
      case "z":
        xPos = 0;
        yPos = relativeY;
        zPos = relativeZ + (x - 2) * flattenSeparation;
        break;
    }

    return [xPos, yPos + baseHeight, zPos] as [number, number, number];
  };

  for (let x = 0; x < 5; x++) {
    for (let y = 0; y < 5; y++) {
      for (let z = 0; z < 5; z++) {
        cubes.push(
          <SmallCube
            key={`${x}-${y}-${z}`}
            position={getPositionWithOffset(x, y, z)}
            number={state[x+y*5+z*25  ]}
          />,
        );
      }
    }
  }

  const handleView = (view: typeof activeAxis) => {
    setActiveAxis((prev) => (prev === view ? "none" : view));
    setFlattenAxis("none");
  };

  const handleFlatten = (axis: typeof flattenAxis) => {
    setFlattenAxis((prev) => (prev === axis ? "none" : axis));
    setActiveAxis("none");
  };

  return (
    <>
      <Canvas
        camera={{ position: cameraPosition, fov: 50 }}
        className="font-inter font-bold"
        shadows
      >
        <color attach="background" args={["#f0f0f0"]} />
        <ambientLight intensity={0.7} />
        <directionalLight
          position={[0, 130, 0]}
          intensity={1}
          castShadow // Enable shadows for light
          shadow-mapSize-width={1024}
          shadow-mapSize-height={1024}
          shadow-camera-far={200}
          shadow-camera-left={-50}
          shadow-camera-right={50}
          shadow-camera-top={50}
          shadow-camera-bottom={-50}
          shadow-bias={-0.005} // Adjust bias to prevent artifacts
        />
        <OrbitControls />
        <mesh
          rotation={[-Math.PI / 2, 0, 0]}
          position={[0, -5, 0]}
          receiveShadow // This will receive the shadow
        >
          <planeGeometry args={[200, 200]} />
          <shadowMaterial opacity={0.1} /> {/* Add blurry shadow */}
        </mesh>

        <group>
          {cubes.map((cube) => (
            <SmallCube
              key={cube.key}
              {...cube.props}
              castShadow // Enable shadow casting for cubes
            />
          ))}
        </group>
        <gridHelper
          receiveShadow
          args={[40, 40, "#555555", "#cccccc"]}
          position={[0, -5, 0]}
        />
      </Canvas>
      <Draggable>
        <div className="absolute right-4 top-1/4 p-4 flex flex-col flex-wrap gap-4 max-w-[400px] bg-white rounded-lg cursor-grab">
          <div className="flex gap-2 m-auto items-center">
            <h2 className="text-2xl font-bold">Controls</h2>
            <Gamepad size={30} />
          </div>
          <div className="grid grid-cols-3 w-full gap-4">
            <Button
              onClick={() => handleView("x")}
              className={activeAxis === "x" ? "bg-blue-500" : ""}
            >
              Separate X
            </Button>
            <Button
              onClick={() => handleView("y")}
              className={activeAxis === "y" ? "bg-blue-500" : ""}
            >
              Separate Y
            </Button>
            <Button
              onClick={() => handleView("z")}
              className={activeAxis === "z" ? "bg-blue-500" : ""}
            >
              Separate Z
            </Button>
            <Button
              onClick={() => handleFlatten("x")}
              className={flattenAxis === "x" ? "bg-green-500" : ""}
            >
              Flatten X
            </Button>
            <Button
              onClick={() => handleFlatten("y")}
              className={flattenAxis === "y" ? "bg-green-500" : ""}
            >
              Flatten Y
            </Button>
            <Button
              onClick={() => handleFlatten("z")}
              className={flattenAxis === "z" ? "bg-green-500" : ""}
            >
              Flatten Z
            </Button>
          </div>
        </div>
      </Draggable>
    </>
  );
};

export default function Component({ numbers }: { numbers: number[] }) {
  const [cameraPosition, setCameraPosition] = useState<
    [number, number, number]
  >([10, 10, 10]);

  const moveCamera = (axis: "x" | "y" | "z", direction: number) => {
    setCameraPosition((prevPosition) => {
      const newPosition = [...prevPosition] as [number, number, number];
      if (axis === "x") newPosition[0] += direction;
      if (axis === "y") newPosition[1] += direction;
      if (axis === "z") newPosition[2] += direction;
      return newPosition;
    });
  };

  return (
    <>
      <div className="w-screen h-screen relative">
        <ParameterizedCube cameraPosition={cameraPosition} moveCamera={moveCamera} state={numbers}/>
      </div>
    </>
  );
}
