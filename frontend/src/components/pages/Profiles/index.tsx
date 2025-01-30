import "./index.css";

const Profiles = () => {
	return (
		<>
			<div className="profiles-container">
				<div className="profiles-title">Select a profile</div>
				<div className="profiles-profiles-container">
					<div className="profiles-profile">
						<div className="profiles-pfp" />
						<div className="profiles-name">Name</div>
					</div>
					<div className="profiles-profile">
						<div className="profiles-pfp" />
						<div className="profiles-name">Name</div>
					</div>
					<div className="profiles-profile">
						<div className="profiles-pfp" />
						<div className="profiles-name">Name</div>
					</div>
					<div className="profiles-profile">
						<div className="profiles-pfp" />
						<div className="profiles-name">New Profile</div>
					</div>
				</div>
			</div>
		</>
	);
};

export default Profiles;
