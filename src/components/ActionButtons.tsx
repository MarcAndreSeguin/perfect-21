type ActionType = "NONE" | "HIT" | "STAND" | "DOUBLE" | "SPLIT" | "SURRENDER";

interface ActionButtonsProps {
  onPress: (action: ActionType) => void;
  disabled?: boolean;
}

export default function ActionButtons({
  onPress,
  disabled,
}: ActionButtonsProps) {
  return (
    <div>
      <h3>Player Actions:</h3>
      <div className="playerActions">
        {(
          [
            "NONE",
            "HIT",
            "STAND",
            "DOUBLE",
            "SPLIT",
            "SURRENDER",
          ] as ActionType[]
        ).map((a) => (
          <button key={a} onClick={() => onPress(a)} disabled={disabled}>
            {a[0] + a.slice(1).toLowerCase()}
          </button>
        ))}
      </div>
    </div>
  );
}
