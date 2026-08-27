export type Edition = {
  id: number;
  competitionName: string;
  startYear: number;
};

export const premierLeagueEdition = {
  id: 2021,
  competitionName: "Premier League",
  startYear: 2026,
} as const;

export const laLigaEdition = {
  id: 2014,
  competitionName: "La Liga",
  startYear: 2026,
} as const;

export const bundesligaEdition = {
  id: 2002,
  competitionName: "Bundesliga",
  startYear: 2026,
} as const;

export const serieAEdition = {
  id: 2019,
  competitionName: "Serie A",
  startYear: 2026,
} as const;

export const ligue1Edition = {
  id: 2015,
  competitionName: "Ligue 1",
  startYear: 2026,
} as const;

export const mainEditions = [
  premierLeagueEdition,
  laLigaEdition,
  bundesligaEdition,
  serieAEdition,
  ligue1Edition,
] as const;

export function getEditionFromStrings(
  competitionID: string | undefined,
  startYear: string | undefined){
  return getEdition(Number(competitionID), Number(startYear) );
}

export function getEdition(competitionID: number, startYear: number){
  return mainEditions.find(
    (edition) =>
      edition.id === competitionID &&
      edition.startYear === startYear,
  );
}