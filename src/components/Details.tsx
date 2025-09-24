type DetailsProps = {
  game: any; // Replace 'any' with a more specific type if available
};

export default function Details({ game }: DetailsProps) {
  return (
    <>
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
    </>
  );
}