import type { Club } from "./club";

export const routes = {
  results: "wyniki",
  matches: "mecze",
  table: "tabela",
  goalScorers: "strzelcy",
} as const;

type Route = (typeof routes)[keyof typeof routes];

export const basePath = "/:competitionID/:startYear";

export function routePath(route: Route) {
  return `${basePath}/${route}`;
}


export const clubaPagePath = "/competitions/:competition/clubs/:clubID/:club_name"

export function createClubPageLink(club: Club){
  return `/competitions/${club.competitionID}/clubs/${club.teamId}/${club.teamName}"`
}