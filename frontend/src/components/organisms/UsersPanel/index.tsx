import { useParams } from "react-router-dom";
import PanelTitle from "../../atoms/PanelTitle";
import SidebarCollapseIcon from "../../atoms/SidebarCollapseIcon";
import GuildCategory from "../../molecules/GuildCategory";
import UserCard from "../../molecules/UserCard";
import "./index.css";
import React from "react";
import * as backend from "../../../../wailsjs/go/main/App";
import { types } from "../../../../wailsjs/go/models";

const UsersPanel = () => {
	const { serverAddress } = useParams();

	const [users, setUsers] = React.useState<types.Users[]>([]);

	React.useEffect(() => {
		if (!serverAddress) return;

		const fetchUsers = async () => {
			const fetchedUsers = await backend.GetUsers(serverAddress);
			setUsers(() => {
				return fetchedUsers.sort((a, b) => a.name.localeCompare(b.name));
			});
		};

		fetchUsers();
	}, []);
	return (
		<div className="users-panel">
			<div className="panel-header">
				<SidebarCollapseIcon />
				<PanelTitle>Users</PanelTitle>
			</div>
			<div className="users-list custom-scrollbar">
				<GuildCategory categoryName="online">
					<div className="users-of-category-wrapper">
						{[...users].map((user) => {
							if (user.status === "online")
								return (
									<UserCard
										key={user.id}
										name={user.name}
										status={user.status}
									/>
								);
						})}
					</div>
				</GuildCategory>
				{[...users].filter((x) => x.status === "offline").length > 0 && (
					<GuildCategory categoryName="offline">
						<div className="users-of-category-wrapper">
							{[...users].map((user) => {
								if (user.status === "offline")
									return (
										<UserCard
											key={user.id}
											name={user.name}
											status={user.status}
										/>
									);
							})}
						</div>
					</GuildCategory>
				)}
			</div>
			<div className="server-profile">
				<div className="server-profile-wrapper">
					<UserCard name="Username (in dev)" status="online" />
				</div>
			</div>
		</div>
	);
};

export default UsersPanel;
