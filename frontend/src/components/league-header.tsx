type LeagueHeaderProps = {
  leagueName: string;
  flag: string;
};

export function LeagueHeader({ leagueName, flag }: LeagueHeaderProps) {
  return (
    <header className="flex h-[40px] w-[800px] items-center gap-[30px] rounded-lg bg-green-750 pl-[15px] pr-[10px]">
      <span
        aria-hidden="true"
        className="h-[25px] w-[35px]"
        style={{ background: flag }}
      />
      <span>{leagueName}</span>
    </header>
  );
}
