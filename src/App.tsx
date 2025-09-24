import { useState } from "react";
import "./App.css";
import About from "./components/About.tsx";
import Details from "./components/Details.tsx";

type GameResponse = Record<string, any> | { error: string } | null;

export default function App() {
  const [game, setGame] = useState<GameResponse>(null);

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
      <div className="game">
        <div className="dealerHand">
          <strong>Dealer:</strong> {game?.dealer?.prettyString?.[0] ?? "—"}{" "}
          {game?.dealer && "+ ?"}
        </div>

        <br />
        <br />

        <div className="playerHand">
          <strong>Player:</strong>{" "}
          {game?.player?.prettyString?.join(" ") ?? "—"}
        </div>
        {/* ACTION BUTTONS TO DO*/}
        <br />
        <button onClick={newGame}>Next Hand →</button>
      </div>

      <Details game={game} />
      <About />
    </>
  );
}
