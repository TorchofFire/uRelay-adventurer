import "./index.css";

interface Props {
	name: string;
	pfp?: string;
	status?: string;
}

const UserCard = (props: Props) => {
	return (
		<div className="user-wrapper">
			<div
				className="pfp"
				style={{
					backgroundImage: props.pfp ? `url(${props.pfp})` : "none",
					backgroundColor: props.pfp ? "transparent" : "var(--text)",
				}}
			/>
			<div className="username-title">{props.name}</div>
			<div className={`status ${props.status}`} />
		</div>
	);
};

export default UserCard;
