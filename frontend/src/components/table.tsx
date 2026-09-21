import type { Club } from "../config/club";
import { TableClub } from "./table-club";

type TableProps = {
  clubs: Club[];
  size?: "small" | "large";
  selectedClub?: number;
};

export function Table({ clubs, size = "large", selectedClub }: TableProps) {
  const isSmall = size === "small";

  return (
    <table
      className={`table-fixed border-collapse ${
        isSmall ? "text-secondary" : "w-full text-primary-small"
      }`}
    >
      <colgroup>
        {isSmall ? (
          <>
            <col className="w-[70px]" />
            <col className="w-[200px]" />
            <col className="w-[60px]" />
            <col className="w-[80px]" />
            <col className="w-[60px]" />
            <col className="w-[160px]" />
          </>
        ) : (
          <>
            <col className="w-[120px]" />
            <col />
            <col className="w-[100px]" />
            <col className="w-[120px]" />
            <col className="w-[120px]" />
            <col className="w-[120px]" />
            <col className="w-[120px]" />
            <col className="w-[150px]" />
            <col className="w-[150px]" />
          </>
        )}
      </colgroup>

      <thead>
        <tr className="text-center text-granit-200">
          <th className="font-normal">miejsce</th>
          <th className="text-left font-normal">nazwa drużyny</th>
          {!isSmall && <th className="font-normal">punkty</th>}
          <th className="font-normal">mecze</th>
          {!isSmall && <th className="font-normal">zwycięstwa</th>}
          {!isSmall && <th className="font-normal">remisy</th>}
          {!isSmall && <th className="font-normal">porażki</th>}
          <th className="font-normal">bilans bramkowy</th>
          {isSmall && <th className="font-normal">punkty</th>}
          <th className="font-normal">forma</th>
        </tr>
      </thead>

      <tbody>
        {clubs.map((club) => (
          <TableClub
            key={club.teamId}
            club={club}
            size={size}
            isSelected={club.teamId === selectedClub}
          />
        ))}
      </tbody>
    </table>
  );
}
