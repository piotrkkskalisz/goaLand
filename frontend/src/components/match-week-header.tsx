type MatchWeekHeaderProps = {
  text: string;
};

export function MatchWeekHeader({ text }: MatchWeekHeaderProps) {
  return (
    <header className="flex h-[40px] w-[800px] items-center rounded-lg bg-green-800 px-[30px] text-secondary">
      {text}
    </header>
  );
}
