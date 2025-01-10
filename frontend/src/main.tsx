import React from "react";
import { createRoot } from "react-dom/client";
import "./style.css";
import App from "./App";
import { EventsOn } from "../wailsjs/runtime/runtime";
import { packet } from "./types/packet.namespace";

const container = document.getElementById("root");

const root = createRoot(container!);

EventsOn("guild_message", (data: packet.GuildMessage) => {
	console.log(data);
});

EventsOn("system_message", (data: packet.SystemMessage) => {
	console.log(data);
});

root.render(
	<React.StrictMode>
		<App />
	</React.StrictMode>
);
