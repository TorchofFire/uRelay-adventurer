import './App.css';
import ChannelsPanel from './components/organisms/ChannelsPanel';
import MessagesPanel from './components/organisms/MessagesPanel';
import NavigationPanel from './components/organisms/NavigationPanel';
import UsersPanel from './components/organisms/UsersPanel';

function App() {
    return (
      <>
        <NavigationPanel />
        <ChannelsPanel />
        <MessagesPanel />
        <UsersPanel />
      </>
    );
  }

export default App
