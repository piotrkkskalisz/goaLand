import type { Club } from "../config/club";
import { TableClub } from "./table-club";

type TableProps = {
  clubs: Club[];
};

export function Table({ clubs }: TableProps) {
  return (
    <table className="text-primary-small w-full table-fixed border-collapse">
      <colgroup>
        <col className="w-[120px]" />
        <col />
        <col className="w-[100px]" />
        <col className="w-[120px]" />
        <col className="w-[120px]" />
        <col className="w-[120px]" />
        <col className="w-[120px]" />
        <col className="w-[150px]" />
        <col className="w-[150px]" />
      </colgroup>

      <thead>
        <tr className="text-center text-granit-200">
          <th className="font-normal">miejsce</th>
          <th className="text-left font-normal">nazwa drużyny</th>
          <th className="font-normal">punkty</th>
          <th className="font-normal">mecze</th>
          <th className="font-normal">zwycięstwa</th>
          <th className="font-normal">remisy</th>
          <th className="font-normal">porażki</th>
          <th className="font-normal">bilans bramkowy</th>
          <th className="font-normal">forma</th>
        </tr>
      </thead>

      <tbody>
        {clubs.map((club, index) => (
          <TableClub key={club.teamId} position={index + 1} club={club} />
        ))}
      </tbody>
    </table>
  );
}
