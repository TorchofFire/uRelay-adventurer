import { Link } from "react-router-dom";
import NavIcon from "../../atoms/NavIcon";
import "./index.css";

const NavigationPanel = () => {
	return (
		<div className="navigation-panel">
			<div className="fixed-nav">
				<NavIcon />
				<NavIcon />
				<NavIcon />
			</div>
			<div id="nav-break" />
			<div className="server-list">
				<NavIcon to="/guild/localhost:8080" />
				<NavIcon />
				<NavIcon />
				<NavIcon />
				<NavIcon />
				<NavIcon />
			</div>
		</div>
	);
};

export default NavigationPanel;
