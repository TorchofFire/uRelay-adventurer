import moment from "moment";
import AutoResizingTextarea from "../../atoms/AutoResizingTextarea";
import FullMessage from "../../molecules/FullMessage";
import "./index.css";
import React from "react";
import * as backend from "../../../../wailsjs/go/main/App";
import { useParams } from "react-router-dom";
import { EventsOff, EventsOn } from "../../../../wailsjs/runtime/runtime";
import LoadingWheel from "../../atoms/LoadingWheel";
import { types } from "../../../../wailsjs/go/models";

const MessagesPanel = () => {
	// TODO: figure in DMs
	const textareaRef = React.useRef<HTMLTextAreaElement | null>(null);
	const loadingWheelRef = React.useRef<HTMLDivElement | null>(null);

	const { serverAddress, channelId } = useParams();
	console.log(serverAddress, channelId);

	const [messages, setMessages] = React.useState<types.GuildMessageEmission[]>(
		[]
	);
	const [loadingWheel, setLoadingWheel] = React.useState(true);
	const [channelName, setChannelName] = React.useState("");

	React.useEffect(() => {
		if (!serverAddress) return;

		const fetchChannels = async () => {
			const fetchedData = await backend.GetChannels(serverAddress);

			const currentChannel = fetchedData.channels.find(
				(channel) => channel.id.toString() === channelId
			);
			if (currentChannel) {
				setChannelName(currentChannel.name);
			}
		};

		fetchChannels();
	}, [serverAddress, channelId]);

	React.useEffect(() => {
		if (!serverAddress || !channelId) return;

		textareaRef.current?.focus();

		const fetchMessages = async () => {
			const lastMessageId =
				messages.length > 0
					? [...messages].sort((a, b) => a.id - b.id)[0]?.id
					: 0;
			const fetchedMessages = await backend.GetMessages(
				serverAddress,
				Number(channelId),
				lastMessageId
			);
			setMessages((prevMessages) => {
				return [...prevMessages, ...fetchedMessages];
			});
			if (fetchedMessages.length === 0) {
				setLoadingWheel(false);
			}
		};

		const handleGuildMessage = (data: types.GuildMessageEmission) => {
			setMessages((prevMessages) => [...prevMessages, data]);
		};

		const observer = new IntersectionObserver(
			(entries) => {
				const loadingWheelEntry = entries[0];
				if (loadingWheelEntry?.isIntersecting) {
					fetchMessages();
				}
			},
			{ root: null, threshold: 1.0 }
		);

		if (loadingWheelRef.current) {
			observer.observe(loadingWheelRef.current);
		}

		EventsOn("guild_message", handleGuildMessage);

		return () => {
			observer.disconnect();
			EventsOff("guild_message");
		};
	}, [messages, serverAddress, channelId]);

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
					<div className="channel-title-text">{channelName}</div>
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
				{loadingWheel && (
					<div ref={loadingWheelRef}>
						<LoadingWheel />
					</div>
				)}
				{!loadingWheel && (
					<div className="empty-at-top">
						{messages.length > 100
							? "Wow, you've reached the top"
							: messages.length > 0
							? "This is the beginning of the conversation"
							: "Be the first to send a message!"}
					</div>
				)}
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
