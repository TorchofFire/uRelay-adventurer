import { Route, Routes } from "react-router-dom";
import "./App.css";
import NavigationPanel from "./components/organisms/NavigationPanel";
import Guild from "./components/pages/Guild";

function App() {
	return (
		<>
			<NavigationPanel />
			<Routes>
				<Route path="/guild/:serverAddress/:channelId?" element={<Guild />} />
			</Routes>
		</>
	);
}

export default App;
