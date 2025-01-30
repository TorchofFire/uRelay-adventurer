import { useParams } from "react-router-dom";
import PanelTitle from "../../atoms/PanelTitle";
import SidebarCollapseIcon from "../../atoms/SidebarCollapseIcon";
import GuildCategory from "../../molecules/GuildCategory";
import UserCard from "../../molecules/UserCard";
import "./index.css";
import React from "react";
import * as backend from "../../../../wailsjs/go/main/App";
import { types } from "../../../../wailsjs/go/models";
import { EventsOff, EventsOn } from "../../../../wailsjs/runtime/runtime";

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

		const handleUser = (data: types.Users) => {
			setUsers((prevUsers) => {
				if (prevUsers.some((user) => user.id === data.id)) {
					const users = prevUsers.map((user) => {
						if (user.id === data.id) return data;
						return user;
					});
					return users;
				}
				return [...prevUsers, data];
			});
		};

		fetchUsers();

		EventsOn("user", handleUser);

		return () => {
			EventsOff("user");
		};
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
					<UserCard name="You (in dev)" status="online" />
				</div>
			</div>
		</div>
	);
};

export default UsersPanel;
