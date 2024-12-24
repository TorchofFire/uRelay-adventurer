import "./index.css";
import React from "react";

interface Props {
	onEnterPress: () => void;
}

const AutoResizingTextarea = React.forwardRef<HTMLTextAreaElement, Props>(
	({ onEnterPress }, ref) => {
		const paddingRef = React.useRef<number>(0);
		React.useEffect(() => {
			if (ref && "current" in ref && ref.current) {
				const textarea = ref.current;
				const styles = window.getComputedStyle(textarea);
				paddingRef.current =
					parseInt(styles.paddingTop, 10) + parseInt(styles.paddingBottom, 10);
			}
		}, [ref]);

		const autoResize = (
			element: HTMLTextAreaElement | React.FormEvent<HTMLTextAreaElement>
		) => {
			const textarea =
				element instanceof HTMLTextAreaElement
					? element
					: element.currentTarget;
			if (!textarea) return;
			textarea.style.height = "auto"; // resets scrollHeight to conform to text
			textarea.style.height = `${textarea.scrollHeight - paddingRef.current}px`;
		};

		const handleKeyDown = (event: React.KeyboardEvent<HTMLTextAreaElement>) => {
			if (event.key === "Enter" && !event.shiftKey) {
				event.preventDefault();
				onEnterPress();
			}
		};

		return (
			<textarea
				ref={ref}
				className="input-box"
				rows={1}
				onKeyDown={handleKeyDown}
				placeholder="Send a message"
				onInput={autoResize}
			/>
		);
	}
);

export default AutoResizingTextarea;
