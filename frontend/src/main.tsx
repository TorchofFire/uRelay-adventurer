import React from "react";
import { createRoot } from "react-dom/client";
import "./style.css";
import App from "./App";
import { EventsOn } from "../wailsjs/runtime/runtime";
import { BrowserRouter } from "react-router-dom";
import { types } from "../wailsjs/go/models";

const container = document.getElementById("root");

const root = createRoot(container!);

EventsOn("guild_message", (data: types.GuildMessageEmission) => {
	console.log(data);
});

EventsOn("system_message", (data: unknown) => {
	// TODO: add type once type is done
	console.log(data);
});

root.render(
	<BrowserRouter>
		<App />
	</BrowserRouter>
);
