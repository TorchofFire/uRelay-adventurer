import { ReactNode } from "react";
import "./index.css";
import { useNavigate } from "react-router-dom";

interface Props {
	to?: string;
}

const NavIcon = (props: Props) => {
	const nav = useNavigate();

	const handleClick = () => {
		nav(props.to || "");
	};

	return <div className="nav-icon" onClick={handleClick} />;
};

export default NavIcon;
