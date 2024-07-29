import { SiGo, SiReact, SiTypescript, SiVite } from "react-icons/si";
import "./App.css";

function App() {
  return (
    <>
      <div className="flex flex-col items-center justify-center w-full font-mono">
        <h1 className="text-3xl font-bold">
          Hello from {" "}
          <span className="underline">go-echo-playground!</span>
        </h1>
        <div className="flex flex-row gap-20 pt-4">
          <div className="flex flex-col items-start justify-start">
            <h2 className="text-2xl text-center">Backend</h2>
            <SiGo className="size-20 text-blue-500" />
          </div>
          <div className="flex flex-col items-start justify-start gap-2">
            <h2 className="text-2xl text-center">Frontend</h2>
            <SiVite className="size-20 text-blue-500" />
            <SiTypescript className="size-20 text-blue-500" />
            <SiReact className="size-20 text-blue-500" />
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
