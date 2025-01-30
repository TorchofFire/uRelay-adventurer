import { useNavigate, useParams } from "react-router-dom";
import PanelTitle from "../../atoms/PanelTitle";
import SidebarCollapseIcon from "../../atoms/SidebarCollapseIcon";
import GuildCategory from "../../molecules/GuildCategory";
import "./index.css";
import React from "react";
import { types } from "../../../../wailsjs/go/models";
import * as backend from "../../../../wailsjs/go/main/App";

const ChannelsPanel = () => {
	const nav = useNavigate();

	const { serverAddress, channelId } = useParams();

	const [channels, setChannels] = React.useState<types.GuildChannels[]>([]);
	const [categories, setCategories] = React.useState<types.GuildCategories[]>(
		[]
	);

	React.useEffect(() => {
		if (!serverAddress) return;

		const fetchChannelsAndCategories = async () => {
			const fetchedChannelsAndCategories = await backend.GetChannels(
				serverAddress
			);
			setChannels(() => {
				return fetchedChannelsAndCategories.channels
					.sort((a, b) => a.name.localeCompare(b.name))
					.sort((a, b) => a.display_priority - b.display_priority);
			});
			setCategories(() => {
				return fetchedChannelsAndCategories.categories
					.sort((a, b) => a.name.localeCompare(b.name))
					.sort((a, b) => a.display_priority - b.display_priority);
			});
		};

		fetchChannelsAndCategories();
	}, [serverAddress, channelId]);
	return (
		<div className="channels-panel">
			<div className="panel-header">
				<PanelTitle>Server Name</PanelTitle>
				<SidebarCollapseIcon />
			</div>
			{[...channels].filter((x) => !x.category_id).length > 0 && (
				<GuildCategory>
					{[...channels].map((channel) => {
						if (!channel.category_id) {
							return (
								<div
									key={channel.id}
									className={`channel ${
										channelId === channel.id.toString() ? "selected" : ""
									}`}
									onClick={() => nav(`/guild/${serverAddress}/${channel.id}`)}
								>
									<div className="hashtag-icon" />
									<div className="channel-name">{channel.name}</div>
								</div>
							);
						}
						return null;
					})}
				</GuildCategory>
			)}
			{[...categories].map((category) => {
				return (
					<GuildCategory key={category.id} categoryName={category.name}>
						{[...channels].map((channel) => {
							if (channel.category_id === category.id) {
								return (
									<div
										key={channel.id}
										className={`channel ${
											channelId === channel.id.toString() ? "selected" : ""
										}`}
										onClick={() => nav(`/guild/${serverAddress}/${channel.id}`)}
									>
										<div className="hashtag-icon" />
										<div className="channel-name">{channel.name}</div>
									</div>
								);
							}
							return null;
						})}
					</GuildCategory>
				);
			})}
		</div>
	);
};

export default ChannelsPanel;
