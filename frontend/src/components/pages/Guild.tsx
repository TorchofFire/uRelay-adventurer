import { Route, Routes, useParams } from "react-router-dom";
import ChannelsPanel from "../organisms/ChannelsPanel";
import MessagesPanel from "../organisms/MessagesPanel";
import UsersPanel from "../organisms/UsersPanel";

const Guild = () => {
	const { channelId } = useParams();
	return (
		<>
			<ChannelsPanel />
			{channelId && <MessagesPanel key={channelId} />}
			<UsersPanel />
		</>
	);
};

export default Guild;
