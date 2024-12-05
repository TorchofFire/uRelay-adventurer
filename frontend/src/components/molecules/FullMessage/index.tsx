import moment from "moment";
import "./index.css";

interface Props {
	username: string;
	date: number;
	message: string;
}

// TODO: have the actual message JSON be here to be copied by the user thru a right click menu

const FullMessage = (props: Props) => {
	const date = moment.unix(props.date);
	let dateFormatted = date.format("MMMM Do, YYYY | hh:mm A");
	if (moment().isSame(date, "year"))
		dateFormatted = date.format("MMMM Do | hh:mm A");
	if (moment().isSame(date, "day"))
		dateFormatted = `Today at ${date.format("hh:mm A")}`;

	return (
		<div className="full-message">
			<div className="message-pfp" />
			<div className="message-wrapper">
				<div className="message-user-wrapper">
					<div className="message-sender">{props.username}</div>
					<div className="message-date">{dateFormatted}</div>
				</div>
				<p className="message-text">{props.message}</p>
			</div>
		</div>
	);
};

export default FullMessage;
