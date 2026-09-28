import { useEffect, useState } from "react";
import { Link, useLocation, useNavigate } from "react-router";
import { GetCompetitionEditions } from "../api/edition";
import type { Edition } from "../config/editions";
import { routes } from "../config/routes";

export function CompetitionTitle(props: Edition) {
  const location = useLocation();
  const navigate = useNavigate();
  const [editions, setEditions] = useState<Edition[]>([props]);

  useEffect(() => {
    GetCompetitionEditions(props.id)
      .then((items) => {
        if (items.length > 0) {
          setEditions(items);
          return;
        }
        setEditions([props]);
      })
      .catch(() => {
        setEditions([props]);
      });
  }, [props.id]);

  const currentPage = location.pathname.split("/").slice(3).join("/") || routes.matches;
  const competitionLocation = `/${props.id}/${props.startYear}/${routes.matches}`;

  const handleSeasonChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const nextYear = Number(event.target.value);
    navigate(`/${props.id}/${nextYear}/${currentPage}`);
  };

  const sortedEditions = [...editions].sort((a, b) => b.startYear - a.startYear);

  return (
    <div className="flex items-center gap-4 w-fit rounded-lg bg-card px-[50px] py-[5px] text-heading uppercase">
      <Link
        className=""
        to={competitionLocation}
      >
        {props.competitionName}
      </Link>

      <select
        value={props.startYear}
        onChange={handleSeasonChange}
          className="rounded-md bg-card px-3 py-2 text-sm text-heading outline-none focus:bg-card focus:outline-none focus:ring-0 [&>option]:bg-card [&>option]:text-heading"
        aria-label="Select season"
      >
        {sortedEditions.map((edition) => (
          <option key={`${edition.id}-${edition.startYear}`} value={edition.startYear}>
            {edition.startYear}/{edition.startYear + 1}
          </option>
        ))}
      </select>
    </div>
  );
}