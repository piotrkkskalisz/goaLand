import { Link } from "react-router/internal/react-server-client";
import type { MatchData } from "../config/matches";
import { createClubPageLink } from "../config/routes";

type MatchProps = MatchData & {
  size?: "small" | "large";
  isCurrent?: boolean;
};

export function Match({ size = "large", isCurrent = true, ...props }: MatchProps) {
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
        {isCurrent ? (
          <Link to={createClubPageLink(props.competitionID, props.homeTeam)} className={isSmall ? "text-secondary" : "text-primary"}>
            {props.homeTeam.name}
          </Link>
        ) : (
          <span className={isSmall ? "text-secondary" : "text-primary"}>{props.homeTeam.name}</span>
        )}
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
        {isCurrent ? (
          <Link
            to={createClubPageLink(props.competitionID, props.awayTeam)}
            className={`text-right ${isSmall ? "text-secondary" : "text-primary"}`}
          >
            {props.awayTeam.name}
          </Link>
        ) : (
          <span className={`text-right ${isSmall ? "text-secondary" : "text-primary"}`}>
            {props.awayTeam.name}
          </span>
        )}
        <span
          className={`bg-light ${
            isSmall ? "h-[27px] w-[27px]" : "h-[40px] w-[40px]"
          }`}
        />
      </div>
    </div>
  );
}
