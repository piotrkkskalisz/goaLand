import { Link } from "react-router";
import type { ClubStats } from "../config/club";
import { createClubPageLink } from "../config/routes";

type TableClubProps = {
  club: ClubStats;
  size?: "small" | "large";
  isSelected?: boolean;
  isCurrent?: boolean;
};

const formResult = {
  win: { label: "Z", className: "bg-green-725" },
  draw: { label: "R", className: "bg-yellow-700" },
  loss: { label: "P", className: "bg-red-800" },
};

export function TableClub({
  club,
  size = "large",
  isSelected = false,
  isCurrent = true,
}: TableClubProps) {
  const playedMatches = club.wins + club.draws + club.losses;
  const isSmall = size === "small";

  return (
    <tr
      className={`border-t border-granit-200 text-center ${
        isSmall ? "h-[35px]" : "h-[50px]"
      } ${isSelected ? "bg-green-750" : ""}`}
    >
      <td>{club.position}.</td>
      <td className="text-left">
        <span className="flex items-center gap-[8px]">
          {club.isLive && (
            <span className="h-[8px] w-[8px] rounded-full bg-live" />
          )}
          {isCurrent ? (
            <Link to={createClubPageLink(club.competitionID, club)}>
              {club.name}
            </Link>
          ) : (
            <span>{club.name}</span>
          )}
        </span>
      </td>
      {!isSmall && <td>{club.points}</td>}
      <td>{playedMatches}</td>
      {!isSmall && <td>{club.wins}</td>}
      {!isSmall && <td>{club.draws}</td>}
      {!isSmall && <td>{club.losses}</td>}
      <td>{`${club.goalsScored}:${club.goalsConceded}`}</td>
      {isSmall && <td>{club.points}</td>}
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
