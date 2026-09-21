export type Player = {
  playerID: number;
  playerName: string;
  nationality: string;
};

export type PlayersOnPosition = {
  players: Player[];
  position: string;
};