import PanelTitle from '../../atoms/PanelTitle';
import SidebarCollapseIcon from '../../atoms/SidebarCollapseIcon';
import GuildCategory from '../../molecules/GuildCategory';
import UserCard from '../../molecules/UserCard';
import './index.css';

const UsersPanel = () => {
    return (
        <div className='users-panel'>
            <div className='panel-header'>
                <SidebarCollapseIcon />
                <PanelTitle>Users</PanelTitle>
            </div>
            <div className='users-list custom-scrollbar'>
                <GuildCategory categoryName='online'>
                    <div className='users-of-category-wrapper'>
                        <UserCard name='Username' status='online' />
                        <UserCard name='Username' status='dnd' />
                        <UserCard name='Username' status='online' />
                    </div>
                </GuildCategory>
                <GuildCategory categoryName='offline'>
                    <div className='users-of-category-wrapper'>
                        <UserCard name='Username' />
                        <UserCard name='Username' />
                        <UserCard name='Username' />
                    </div>
                </GuildCategory>
            </div>
            <div className='server-profile'>
                <div className='server-profile-wrapper'>
                    <UserCard name='Username' status='online' />
                </div>
            </div>
        </div>
    );
};

export default UsersPanel;