import { useState } from "react";
import "./App.css";
import About from "./components/About";
import Details from "./components/Details";
import PlayArea from "./components/PlayArea";
import ActionButtons from "./components/ActionButtons";

type GameResponse = Record<string, any> | { error: string } | null;
type ActionType = "NONE" | "HIT" | "STAND" | "DOUBLE" | "SPLIT" | "SURRENDER";

export default function App() {
  const [game, setGame] = useState<GameResponse>(null);
  const [result, setResult] = useState<null | "correct" | "incorrect">(null);

  const g = game as any;

  async function newGame() {
    setResult(null);
    try {
      const res = await fetch("https://perfect-21-api.onrender.com/play");
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const data = await res.json();
      setGame(data);
    } catch (e) {
      setGame({ error: String(e) });
      setResult(null);
      setGame(null);
    }
  }

  function checkAnswer(playerAction: ActionType) {
    if (!g) return;
    const ok = playerAction === g.correctAction;
    setResult(ok ? "correct" : "incorrect");
  }

  return (
    <>
      <h1>Perfect 21</h1>
      <PlayArea game={g} />

      {result && (
        <p
          style={{
            marginTop: 3,
            fontWeight: 600,
            color: result === "correct" ? "green" : "crimson",
          }}
        >
          {result === "correct"
            ? "✅ Correct!"
            : `❌ Incorrect — correct action was ${g?.correctAction}`}
        </p>
      )}

      <br />
      <ActionButtons onPress={checkAnswer} disabled={!g || !!result} />
      <br />
      <button onClick={newGame}>Next Hand →</button>

      <br />
      <br />
      <Details game={game} />

      <About />
    </>
  );
}
