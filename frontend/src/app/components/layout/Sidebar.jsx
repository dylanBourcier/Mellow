'use client';
import { usePathname } from 'next/navigation';
import React from 'react';
import Navlink from '../ui/Navlink';
import Image from 'next/image';
import Button from '../ui/Button';

function Sidebar(props) {
  const pathname = usePathname();
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
        <Navlink href="/" icon="home" isActive={pathname === '/'}>
          Home
        </Navlink>
        <Navlink
          href="/search"
          icon="search"
          isActive={pathname.startsWith('/search')}
        >
          Search
        </Navlink>
        <Navlink
          href="/messages"
          icon="messages"
          isActive={pathname.startsWith('/messages')}
        >
          Messages
        </Navlink>
        <Navlink
          href="/groups"
          icon="groups"
          isActive={pathname.startsWith('/groups')}
        >
          Groups
        </Navlink>
        <Navlink
          href="/notifications"
          icon="notifications"
          isActive={pathname.startsWith('/notifications')}
        >
          Notifications
        </Navlink>
        <Navlink
          href="/profile"
          icon="profile"
          isActive={pathname.startsWith('/profile')}
        >
          Profile
        </Navlink>
        <Button className="mt-6 w-full" href="/posts/create">New Post</Button>
      </nav>
      <div className="w-full">
        <Button href="/login" className="w-full">
          Login
        </Button>
      </div>
    </div>
  );
}

export default Sidebar;
