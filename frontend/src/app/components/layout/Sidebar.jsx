import React from 'react';
import Navlink from '../ui/Navlink';

function Sidebar(props) {

    return (
        <div className="hidden lg:flex flex-col items-start justify-start h-full w-72 bg-white shadow-(--box-shadow) p-4 rounded-2xl">
            <div>logo</div>
            <nav className="flex flex-col flex-1 items-start justify-start w-full">
            <Navlink href="/" icon="home" isActive >Home</Navlink>
            <Navlink href="/search" icon="search">Search</Navlink>
            <Navlink href="/messages" icon="messages">Messages</Navlink>
            <Navlink href="/groups" icon="groups">Groups</Navlink>
            <Navlink href="/notifications" icon="notifications">Notifications</Navlink>
            <Navlink href="/profile" icon="profile">Profile</Navlink>
            </nav>
            <div className=''>logout</div>
        </div>
    );
}

export default Sidebar;