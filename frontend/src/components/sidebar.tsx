import { Link } from "react-router";

const leagues = [
  {
    name: "Premier League",
    code: "PL",
    startYear: 2026,
    flag: "linear-gradient(90deg, transparent 42%, #ce1124 42% 58%, transparent 58%), linear-gradient(transparent 35%, #ce1124 35% 65%, transparent 65%), #fff",
  },
  {
    name: "La liga",
    code: "PD",
    startYear: 2026,
    flag: "linear-gradient(#aa151b 0 25%, #f1bf00 25% 75%, #aa151b 75%)",
  },
  {
    name: "Serie A",
    code: "SA",
    startYear: 2026,
    flag: "linear-gradient(90deg, #009246 0 33%, #fff 33% 67%, #ce2b37 67%)",
  },
  {
    name: "Bundesliga",
    code: "BL1",
    startYear: 2026,
    flag: "linear-gradient(#000 0 33%, #dd0000 33% 67%, #ffce00 67%)",
  },
  {
    name: "Ligue 1",
    code: "FL1",
    startYear: 2026,
    flag: "linear-gradient(90deg, #002654 0 33%, #fff 33% 67%, #ce1126 67%)",
  },
];

export function Sidebar() {
  return (
    <aside className="flex h-[650px] w-[370px] flex-col items-center gap-[40px] px-[10px] py-[50px]">
      {leagues.map((league) => (
        <Link
          className="flex h-[60px] w-[350px] items-center gap-[30px] rounded-lg px-[10px] text-left transition-colors hover:bg-granit-800"
          key={league.code}
          to={`/${league.code}/${league.startYear}/wyniki`}
        >
          <span
            aria-hidden="true"
            className="h-[40px] w-[60px]"
            style={{ background: league.flag }}
          />
          <span className="text-primary text-text-primary">{league.name}</span>
        </Link>
      ))}
    </aside>
  );
}
