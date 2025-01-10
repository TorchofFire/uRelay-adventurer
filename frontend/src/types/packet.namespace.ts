export namespace packet {
	export interface GuildMessage {
		channel_id: number;
		sender_id: number;
		message: string;
		id: number;
	}
	export interface SystemMessage {
		severity: "info" | "warning" | "danger";
		message: string;
		channel_id: number;
	}
}
