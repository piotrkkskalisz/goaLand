import type { ReactNode } from "react";
import type { MatchData } from "../config/matches";

type winner = "team1" | "team2";

type TopMatchProps = {
  team1: string;
  team2: string;
  winner?: winner
  separator: string;
  children: ReactNode;
};

function TopMatch({team1, team2, winner, separator, children}: TopMatchProps) {
  return (
    <div className="h-[48px] w-[128px] shrink-0 flex flex-col items-center -space-y-[2px]">
      <div className="text-secondary relative flex w-[128px]  h-[20px] items-center justify-center rounded-[5px]  bg-sections">
        <span className={winner === "team1" ? "font-bold" : undefined}>
          {team1}
        </span>
        {separator}
        <span className={winner === "team2" ? "font-bold" : undefined}>
          {team2}
        </span>
      </div>
      <div className="flex h-[30px] w-[80px] items-center justify-center leading-[13px] bg-sections [clip-path:polygon(0_0,100%_0,87.5%_100%,12.5%_100%)]">
        <div className="text-third text-center">{children}</div>
      </div>
    </div>
  );
}


export function UpcomingMatch({homeTeamCode, awayTeamCode, date, time,
}: MatchData) {
  return (
    <TopMatch team1={homeTeamCode} separator=" - " team2={awayTeamCode}>
      <div>{date}</div>
      <div>{time}</div>
    </TopMatch>
  );
}

export function LiveMatch({homeTeamCode, awayTeamCode, homeScore, awayScore,
}: MatchData) {
  return (
    <TopMatch
      team1={`${homeTeamCode} ${homeScore}`}
      separator=" : "
      team2={`${awayScore} ${awayTeamCode}`}
    >
      <span className="text-third text-red-600">trwa</span>
    </TopMatch>
  );
}

export function FinishedMatch({homeTeamCode, awayTeamCode, homeScore, awayScore,
}: MatchData) {
  let matchWinner: winner | undefined
  if (homeScore! > awayScore!){
    matchWinner = "team1"
  }
  if (homeScore! < awayScore!){
    matchWinner = "team2"
  }

  return (
    <TopMatch
      team1={`${homeTeamCode} ${homeScore}`}
      separator=" : "
      team2={`${awayScore} ${awayTeamCode}`}
      winner={matchWinner}
    >
      <span className="text-third"> koniec </span>
    </TopMatch>
  );
}
