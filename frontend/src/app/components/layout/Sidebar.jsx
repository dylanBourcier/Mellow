import React from 'react';
import Navlink from '../ui/Navlink';
import Image from 'next/image';
import Button from '../ui/Button';

function Sidebar(props) {
  return (
    <div
      className="hidden lg:flex fixed top-6 flex-col self-start items-start justify-start h-[95dvh] box-border w-72 bg-white shadow-(--box-shadow) p-4 rounded-2xl"
      style={{ left: 'max(1.5rem, calc((100vw - 1280px) / 2))' }}
    >
      <nav className="flex flex-col flex-1 items-start justify-start h-auto w-full gap-2">
        <div className="">
          <Image
            src="/img/Logo&Name.svg"
            alt="Logo"
            width={152}
            height={56}
          ></Image>
        </div>
        <Navlink href="/" icon="home">
          Home
        </Navlink>
        <Navlink href="/search" icon="search">
          Search
        </Navlink>
        <Navlink href="/messages" icon="messages">
          Messages
        </Navlink>
        <Navlink href="/groups" icon="groups">
          Groups
        </Navlink>
        <Navlink href="/notifications" icon="notifications">
          Notifications
        </Navlink>
        <Navlink href="/profile" icon="profile">
          Profile
        </Navlink>
        <Button className="mt-6 w-full">Post</Button>
      </nav>
      <div className="w-full">
        <Button className="w-full">Login</Button>
      </div>
    </div>
  );
}

export default Sidebar;
