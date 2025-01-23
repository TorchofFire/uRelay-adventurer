export namespace backendData {
	export interface GuildMessage {
		guild_id: string;
		channel_id: number;
		sender_id: number;
		sender_name: string;
		message: string;
		id: number;
		sent_at: number;
	}
	export interface SystemMessage {
		guild_id: string;
		severity: "info" | "warning" | "danger";
		message: string;
		channel_id: number;
	}
}
