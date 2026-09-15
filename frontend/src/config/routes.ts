
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