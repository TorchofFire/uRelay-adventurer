import moment from "moment";
import AutoResizingTextarea from "../../atoms/AutoResizingTextarea";
import FullMessage from "../../molecules/FullMessage";
import "./index.css";
import React from "react";

const MessagesPanel = () => {
	const textareaRef = React.useRef<HTMLTextAreaElement | null>(null);

	const handleMessagesSend = () => {
		if (!textareaRef.current) return;
		console.log(textareaRef.current.value);
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
				<FullMessage
					username="Torch"
					date={moment().unix()}
					message="hello there"
				/>
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
