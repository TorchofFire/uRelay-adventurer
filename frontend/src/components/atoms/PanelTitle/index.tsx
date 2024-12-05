import { ReactNode } from "react";
import "./index.css";

interface Props {
	children?: ReactNode;
}

const PanelTitle = (props: Props) => {
	return <div className="panel-title">{props.children}</div>;
};

export default PanelTitle;
