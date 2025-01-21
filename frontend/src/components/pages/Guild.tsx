import { Route, Routes } from "react-router-dom";
import ChannelsPanel from "../organisms/ChannelsPanel";
import MessagesPanel from "../organisms/MessagesPanel";
import UsersPanel from "../organisms/UsersPanel";

const Guild = () => {
	return (
		<>
			<ChannelsPanel />
			<Routes>
				<Route path=":channelId" element={<MessagesPanel />} />
			</Routes>
			<UsersPanel />
		</>
	);
};

export default Guild;
