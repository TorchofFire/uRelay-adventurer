import { useRef } from 'react';
import './index.css';

const AutoResizingTextarea = () => {
    const textareaRef = useRef<HTMLTextAreaElement | null>(null);

    const autoResize = () => {
        const textarea = textareaRef.current;
        if (!textarea) return;
        const padding = parseInt(window.getComputedStyle(textarea).paddingTop) + parseInt(window.getComputedStyle(textarea).paddingBottom);
        textarea.style.height = `${textarea.scrollHeight - padding}px`;
    };

    return (
        <textarea
            ref={textareaRef}
            className='input-box'
            rows={1}
            placeholder='Send a message'
            onInput={autoResize}
        />
    );
};

export default AutoResizingTextarea;
