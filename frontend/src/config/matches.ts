export type MatchData = {
  id: number;
  date: string;
  time: string;
  status: "scheduled" | "live" | "finished";
  homeTeam: string;
  awayTeam: string;
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