import { ReactNode } from "react";
import "./index.css";

interface Props {
	children?: ReactNode;
	categoryName: string;
}

const GuildCategory = (props: Props) => {
	return (
		<div className="guild-category">
			<div className="category-label">
				<div className="angle-down-icon" />
				<div className="category-name">{props.categoryName.toUpperCase()}</div>
			</div>
			{props.children}
		</div>
	);
};

export default GuildCategory;
