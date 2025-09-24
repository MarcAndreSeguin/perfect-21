import { useState } from "react";
import "./App.css";

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
        <button onClick={newGame}>Next Hand →</button>
      </div>

      <details className="mt-3">
        <summary className="cursor-pointer select-none text-sm">
          Show API Response 
        </summary>
        <pre
          style={{ textAlign: "left" }}
          className="mt-2 font-mono tabular-nums leading-6 whitespace-pre rounded bg-neutral-950 text-neutral-100 p-4 overflow-auto text-sm"
        >
          {game
            ? JSON.stringify(game, null, 2)
            : "No data yet. Click Next to fetch."}
        </pre>
      </details>

      <p className="about">By Marc-Andre Seguin | 2025 | JazzStack.io</p>
      <p className="about">
        Based on H17 perfect strategy by{" "}
        <a
          href="https://www.blackjackapprenticeship.com/blackjack-strategy-charts/"
          target="_blank"
          rel="noopener noreferrer"
        >
          BlackjackApprenticeship.com
        </a>
      </p>
    </>
  );
}
