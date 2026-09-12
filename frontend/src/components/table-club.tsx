import type { Club } from "../config/club";

type TableClubProps = {
  position: number;
  club: Club;
};

const formResult = {
  win: { label: "Z", className: "bg-green-725" },
  draw: { label: "R", className: "bg-yellow-700" },
  loss: { label: "P", className: "bg-red-800" },
};

export function TableClub({ position, club }: TableClubProps) {
  const playedMatches = club.wins + club.draws + club.losses;

  return (
    <tr className="border-t border-granit-200 text-center h-[50px]">
      <td>{position}.</td>
      <td className="text-left">
        <span className="flex items-center gap-[8px]">
          {club.isLive && (
            <span className="h-[8px] w-[8px] rounded-full bg-live" />
          )}
          {club.teamName}
        </span>
      </td>
      <td> {club.points}</td>
      <td>{playedMatches}</td>
      <td>{club.wins}</td>
      <td>{club.draws}</td>
      <td>{club.losses}</td>
      <td>{`${club.goalsScored}:${club.goalsConceded}`}</td>
      <td>
        <span className="flex items-center justify-center gap-[5px]">
          {club.form.map((result, index) => {
            let size = 25 
            if (club.form.length === index + 1) {
              size = 30
            }
            return (<span
              key={`${result}-${index}`}
              className={`flex items-center justify-center ${formResult[result].className}`}
                style={{
                  height: `${size}px`,
                  width: `${size}px`,
                }}
            >
              {formResult[result].label}
            </span>
          )})}
        </span>
      </td>
    </tr>
  );
}
