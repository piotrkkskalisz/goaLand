import type { Player } from "../config/team-players";


type PlayerProps = {
  player: Player;
};

type PositionProps = {
  position: string;
};

export function Player({ player }: PlayerProps) {
  return (
    <div className="flex w-[400px] items-center justify-between gap-5 bg-sections pl-[50px] pr-[30px] pb-[4px] pt-[4px]">
      <span> {player.playerName} </span>
      <span className="text-right text-granit-200">{player.nationality} </span>
    </div>
  );
}

export function Position({ position }: PositionProps) {
  return (
    <div className="flex w-[400px] items-center bg-green-800 font-bold px-[30px] pb-[4px] pt-[4px]">
      <span>{position} </span>
    </div>
  );
}
