export type MatchResult = "win" | "draw" | "loss";

export type Club = {
  teamId: number;
  teamName: string;
  points: number;
  wins: number;
  draws: number;
  losses: number;
  goalsScored: number;
  goalsConceded: number;
  form: MatchResult[];
  isLive: boolean;
};
