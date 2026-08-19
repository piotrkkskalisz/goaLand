type NavigatorButtonProps = {
  isSelected: boolean;
  text: string;
};

export function NavigatorButton({
  isSelected,
  text,
}: NavigatorButtonProps) {
  return (
    <button
      aria-pressed={isSelected}
      className={`h-[60px] w-[180px] rounded-sm px-[10px] py-[15px] text-primary ${
        isSelected ? "bg-green-700" : "bg-green-750"
      }`}
      type="button"
    >
      {text}
    </button>
  );
}
