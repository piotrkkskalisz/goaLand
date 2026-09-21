export type MatchResult = "win" | "draw" | "loss";

export type Club = {
  position: number;
  teamId: number;
  teamName: string;
  competitionID: number;
  points: number;
  wins: number;
  draws: number;
  losses: number;
  goalsScored: number;
  goalsConceded: number;
  form: MatchResult[];
  isLive: boolean;
};
