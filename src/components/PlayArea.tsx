
type PlayAreaProps = {
  game: any; // Replace 'any' with a more specific type if available
};

function PlayArea({ game }: PlayAreaProps) {
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

export default PlayArea;