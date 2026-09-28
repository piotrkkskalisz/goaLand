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


export function createClubPageLink(competitionID: number, club: Club){
  return `/competitions/${competitionID}/clubs/${club.id}/${club.name}`;
}