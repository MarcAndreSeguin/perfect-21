import { useState } from "react";
import "./App.css";
import About from "./components/About";
import Details from "./components/Details";
import PlayArea from "./components/PlayArea"

type GameResponse = Record<string, any> | { error: string } | null;

export default function App() {
  const [game, setGame] = useState<GameResponse>(null);
  const g = game as any;

  async function newGame() {
    try {
      const res = await fetch("https://perfect-21-api.onrender.com/play");
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const data = await res.json();
      setGame(data);
    } catch (e) {
      setGame({ error: String(e) });
    }
  }

  return (
    <>
      <h1>Perfect 21</h1>
      < PlayArea game={g}/>
        
        {/* ACTION BUTTONS TO DO*/}
        <br />
        <button onClick={newGame}>Next Hand →</button>
      
      
      <br /><br />
      <Details game={game} />
     
      <About />
    </>
  );
}
