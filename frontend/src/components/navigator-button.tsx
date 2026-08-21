import { NavLink } from "react-router";

type NavigatorButtonProps = {
  text: string;
  to: string;
};

export function NavigatorButton({ text, to }: NavigatorButtonProps) {
  return (
    <NavLink
      className={({ isActive }) =>
        `flex h-[60px] w-[180px] items-center justify-center rounded-sm px-[10px] py-[15px] text-primary ${
          isActive ? "bg-green-700" : "bg-green-750 hover:bg-green-725"
        }`
      }
      to={to}
    >
      {text}
    </NavLink>
  );
}
