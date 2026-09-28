export type Competition = {
  id: number;
  competitionName: string;
  isCurrent: boolean;
};

export type Edition = Competition & {
  startYear: number;
};


export const premierLeagueCompetition = {
  id: 2021,
  competitionName: "Premier League",
} as const;

export const laLigaCompetition = {
  id: 2014,
  competitionName: "La Liga",
} as const;

export const bundesligaCompetition = {
  id: 2002,
  competitionName: "Bundesliga",
} as const;

export const serieACompetition = {
  id: 2019,
  competitionName: "Serie A",
} as const;

export const ligue1Competition = {
  id: 2015,
  competitionName: "Ligue 1",
} as const;

export const mainCompetitions = [
  premierLeagueCompetition,
  laLigaCompetition,
  bundesligaCompetition,
  serieACompetition,
  ligue1Competition,
] as const;

/*
export function getEditionFromStrings(
  competitionID: string | undefined,
  startYear: string | undefined){
  return getEdition(Number(competitionID), Number(startYear) );
}

export function getEdition(competitionID: number, startYear: number): Edition | undefined {
  const competition = mainCompetitions.find(
    (competition) => competition.id === competitionID,
  );

  if (!competition) {
    return undefined;
  }

  return {
    ...competition,
    startYear,
  };
}
*/