import moment from "moment";
import AutoResizingTextarea from "../../atoms/AutoResizingTextarea";
import FullMessage from "../../molecules/FullMessage";
import "./index.css";
import React from "react";
import * as backend from "../../../../wailsjs/go/main/App";
import { useParams } from "react-router-dom";
import { backendData } from "../../../types/backendData.namespace";
import { EventsOn } from "../../../../wailsjs/runtime/runtime";

const MessagesPanel = () => {
	// TODO: figure in DMs
	const textareaRef = React.useRef<HTMLTextAreaElement | null>(null);

	const { serverAddress, channelId } = useParams();

	const [messages, setMessages] = React.useState<backendData.GuildMessage[]>(
		[]
	);

	React.useEffect(() => {
		if (!serverAddress || !channelId) return;

		const fetchMessages = async () => {
			const fetchedMessages = await backend.GetMessages(
				serverAddress || "",
				Number(channelId),
				0
			);
			setMessages(fetchedMessages);
		};

		fetchMessages();

		const handleGuildMessage = (data: backendData.GuildMessage) => {
			// Add the new message to the state
			setMessages((prevMessages) => [
				...prevMessages,
				{
					channel_id: data.channel_id,
					guild_id: data.guild_id,
					id: data.id,
					message: data.message,
					sender_id: data.sender_id,
					sender_name: data.sender_name,
					sent_at: data.sent_at,
				},
			]);
		};

		EventsOn("guild_message", handleGuildMessage);
	}, []);

	const handleMessagesSend = () => {
		if (!textareaRef.current || !serverAddress || !channelId) return;
		backend
			.SendMessage(serverAddress, textareaRef.current.value, Number(channelId))
			.catch((error) => {
				console.error(error);
			});
		textareaRef.current.value = "";
		textareaRef.current.style.height = "auto";
	};

	return (
		<div className="messages-panel">
			<div className="messages-panel-header">
				<div className="title-wrapper">
					<div className="big-title-icon hashtag" />
					<div className="channel-title-text">Channel Name</div>
				</div>
			</div>
			<div className="messages-container custom-scrollbar">
				{[...messages]
					.sort((a, b) => b.sent_at - a.sent_at)
					.map((msg, index) => (
						<FullMessage
							key={index}
							username={msg.sender_name}
							date={msg.sent_at}
							message={msg.message}
						/>
					))}
			</div>
			<div className="input">
				<AutoResizingTextarea
					ref={textareaRef}
					onEnterPress={handleMessagesSend}
				/>
				<div className="message-send" onClick={handleMessagesSend} />
			</div>
		</div>
	);
};

export default MessagesPanel;
