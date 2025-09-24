type Hand = {
  prettyString?: string[];
  isBlackJack?: boolean;
};

type Game = {
  dealer?: Hand;
  player?: Hand;
};

export default function PlayArea({ game }: { game: Game }) {
  return (
    <div className="playArea">
        <div className="dealerHand">
          <strong>Dealer:</strong> {game?.dealer?.prettyString?.[0] ?? "—"}{" "}
          {game?.dealer && "+ ?"}
        </div>
        <div className="dealerBJ">
          Dealer blackjack? <em>{String(game?.dealer?.isBlackJack)}</em>
        </div>

        <br />
        <br />

        <div className="playerHand">
          <strong>Player:</strong>{" "}
          {game?.player?.prettyString?.join(" ") ?? "—"}
        </div>
      </div> 
  );
};
