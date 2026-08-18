type MatchBaseProps = {
  date: string;
  time: string;
  homeTeam: string;
  awayTeam: string;
};

type LiveMatchProps = MatchBaseProps & {
  status: "live";
  homeScore: number;
  awayScore: number;
  minute: number;
};

type FinishedMatchProps = MatchBaseProps & {
  status: "finished";
  homeScore: number;
  awayScore: number;
};

type ScheduledMatchProps = MatchBaseProps & {
  status: "scheduled";
};

export type MatchProps =
  | LiveMatchProps
  | FinishedMatchProps
  | ScheduledMatchProps;

export function Match(props: MatchProps) {
  const isLive = props.status === "live";
  const score =
    props.status === "scheduled"
      ? "-"
      : `${props.homeScore}:${props.awayScore}`;

  return (
    <div className="grid h-[60px] w-[890px] grid-cols-[146px_312px_70px_312px] items-center px-[20px]">
      <div className="flex items-center justify-between pr-[10px] text-center">
        <div>
          <div>{props.date}</div>
          <div>{props.time}</div>
        </div>

        {isLive && (
          <div className="text-live-text">
            <div className="mx-auto h-[15px] w-[15px] rounded-full bg-live" />
            <div>LIVE</div>
          </div>
        )}
      </div>

      <div className="flex items-center gap-[10px] px-[10px]">
        <span className="h-[40px] w-[40px] bg-light" />
        <span className="text-primary">{props.homeTeam}</span>
      </div>

      <div className={`text-center text-primary ${isLive ? "text-live" : ""}`}>
        <div>{score}</div>
        {isLive && <div className="text-secondary">{props.minute}'</div>}
      </div>

      <div className="flex items-center justify-end gap-[10px] px-[10px]">
        <span className="text-primary">{props.awayTeam}</span>
        <span className="h-[40px] w-[40px] bg-light" />
      </div>
    </div>
  );
}
