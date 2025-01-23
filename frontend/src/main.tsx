import React from "react";
import { createRoot } from "react-dom/client";
import "./style.css";
import App from "./App";
import { EventsOn } from "../wailsjs/runtime/runtime";
import { backendData } from "./types/backendData.namespace";
import { BrowserRouter } from "react-router-dom";

const container = document.getElementById("root");

const root = createRoot(container!);

EventsOn("guild_message", (data: backendData.GuildMessage) => {
	console.log(data);
});

EventsOn("system_message", (data: backendData.SystemMessage) => {
	console.log(data);
});

root.render(
	<BrowserRouter>
		<App />
	</BrowserRouter>
);
