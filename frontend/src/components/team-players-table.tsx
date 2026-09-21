
import type { PlayersOnPosition } from "../config/team-players";
import { Player, Position } from "./player";

type PlayersOnPositionProps = {
  playersInTeam: PlayersOnPosition[];
};

export function TeamPlayersTable( {playersInTeam}: PlayersOnPositionProps) {
  return (
    <div className="w-[630px] flex flex-col items-center">
      {playersInTeam.map(({players, position}) => (
        <div key={position}>
          <Position position={position} />
          {players.map((player) => (
            <Player key={player.playerID} player={player} />
          ))}
        </div>
      ))}
    </div>
  );
}
