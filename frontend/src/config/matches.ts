import type { Club } from "./club";

export type MatchData = {
  id: number;
  competitionID: number;

  date: string;
  time: string;
  status: "scheduled" | "live" | "finished";
  
  homeTeam: Club;
  awayTeam: Club;

  homeScore?: number;
  awayScore?: number;
};


export type Round = {
  stage: string;
  matchday: number | null;
};


export type RoundMatches = {
  round: Round;
  matches: MatchData[];
};