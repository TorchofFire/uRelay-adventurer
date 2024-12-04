import PanelTitle from '../../atoms/PanelTitle';
import SidebarCollapseIcon from '../../atoms/SidebarCollapseIcon';
import GuildCategory from '../../molecules/GuildCategory';
import './index.css';

const ChannelsPanel = () => {
    return (
        <div className='channels-panel'>
            <div className='panel-header'>
                <PanelTitle>Server Name</PanelTitle>
                <SidebarCollapseIcon />
            </div>
            <GuildCategory categoryName={'placeholder'}>
                <div className='channels-wrapper'>
                    <div className='channel'>
                        <div className='hashtag-icon' />
                        <div className='channel-name'>Name</div>
                    </div>
                    <div className='channel'>
                        <div className='hashtag-icon' />
                        <div className='channel-name'>Name</div>
                    </div>
                </div>
            </GuildCategory>
            <GuildCategory categoryName={'placeholder'}>
                <div className='channels-wrapper'>
                    <div className='channel'>
                        <div className='hashtag-icon' />
                        <div className='channel-name'>Name</div>
                    </div>
                    <div className='channel'>
                        <div className='hashtag-icon' />
                        <div className='channel-name'>Name</div>
                    </div>
                </div>
            </GuildCategory>
        </div>
    );
};

export default ChannelsPanel;