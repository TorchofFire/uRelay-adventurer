import NavIcon from "../../atoms/NavIcon";
import "./index.css";

const NavigationPanel = () => {
	return (
		<div className="navigation-panel">
			<div className="fixed-nav">
				<NavIcon to="/settings" />
				<NavIcon to="/direct-messages" />
				<NavIcon to="/network" />
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
