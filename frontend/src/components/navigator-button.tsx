import { NavLink } from "react-router";

type NavigatorButtonProps = {
  text: string;
  to: string;
  size?: "small" | "large";
};

export function NavigatorButton({
  text,
  to,
  size = "large",
}: NavigatorButtonProps) {
  return (
    <NavLink
      className={({ isActive }) =>
        `flex items-center justify-center rounded-xs text-primary ${
          size === "small"
            ? "h-[30px] w-[100px]"
            : "h-[60px] w-[180px]"
        } ${
          isActive ? "bg-green-700" : "bg-green-750 hover:bg-green-725"
        }`
      }
      to={to}
    >
      {text}
    </NavLink>
  );
}
