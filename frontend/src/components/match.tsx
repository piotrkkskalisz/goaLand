import type { MatchData } from "../config/matches";

type MatchProps = MatchData & {
  size?: "small" | "large";
};

export function Match({ size = "large", ...props }: MatchProps) {
  const isLive = props.status === "live";
  const score =
    props.status === "scheduled"
      ? "-"
      : `${props.homeScore}:${props.awayScore}`;
  const isSmall = size === "small";

  return (
    <div
      className={`grid items-center ${
        isSmall
          ? "h-[40px] w-[677px] grid-cols-[97px_250px_47px_250px] px-[13px]"
          : "h-[60px] w-[890px] grid-cols-[146px_312px_70px_312px] px-[20px]"
      }`}
    >
      <div
        className={`flex items-center justify-between text-center ${
          isSmall ? "pr-[7px] text-third" : "pr-[10px]"
        }`}
      >
        <div>
          <div>{props.date}</div>
          <div>{props.time}</div>
        </div>

        {isLive && (
          <div className="text-live-text">
            <div
              className={`mx-auto rounded-full bg-live ${
                isSmall ? "h-[10px] w-[10px]" : "h-[15px] w-[15px]"
              }`}
            />
            <div>LIVE</div>
          </div>
        )}
      </div>

      <div
        className={`flex items-center ${
          isSmall ? "gap-[7px] px-[7px]" : "gap-[10px] px-[10px]"
        }`}
      >
        <span
          className={`bg-light ${
            isSmall ? "h-[27px] w-[27px]" : "h-[40px] w-[40px]"
          }`}
        />
        <span className={isSmall ? "text-secondary" : "text-primary"}>
          {props.homeTeam}
        </span>
      </div>

      <div
        className={`text-center ${
          isSmall ? "text-secondary" : "text-primary"
        } ${isLive ? "text-live" : ""}`}
      >
        {score}
      </div>

      <div
        className={`flex items-center justify-end ${
          isSmall ? "gap-[7px] px-[7px]" : "gap-[10px] px-[10px]"
        }`}
      >
        <span className={isSmall ? "text-secondary" : "text-primary"}>
          {props.awayTeam}
        </span>
        <span
          className={`bg-light ${
            isSmall ? "h-[27px] w-[27px]" : "h-[40px] w-[40px]"
          }`}
        />
      </div>
    </div>
  );
}
