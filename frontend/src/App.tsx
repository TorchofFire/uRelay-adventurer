import { Route, Routes } from "react-router-dom";
import "./App.css";
import NavigationPanel from "./components/organisms/NavigationPanel";
import Guild from "./components/pages/Guild";
import Settings from "./components/pages/Settings";
import Profiles from "./components/pages/Profiles";

function App() {
	return (
		<>
			<NavigationPanel />
			<Routes>
				<Route path="/settings" element={<Profiles />} />
				{/*TODO: Switch out profiles for settings*/}
				<Route path="/guild/:serverAddress/:channelId?" element={<Guild />} />
			</Routes>
		</>
	);
}

export default App;
