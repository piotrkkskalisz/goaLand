export type MatchResult = "win" | "draw" | "loss";

export type Club = {
  id: number;
  name: string;
  code: string,
}
export type ClubStats = Club & {
  position: number;
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
