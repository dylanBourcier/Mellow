'use client';
import { usePathname } from 'next/navigation';
import React from 'react';
import Navlink from '../ui/Navlink';
import Image from 'next/image';
import Button from '../ui/Button';
import { useUser } from '@/app/context/UserContext';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';
import UserInfo from '../ui/UserInfo';
import LogoutButton from '../ui/LogoutButton';
import Spinner from '../ui/Spinner';

function Sidebar(props) {
  const pathname = usePathname();

  const { user, loading } = useUser();

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
        <Button className="mt-6 w-full" href="/posts/create">
          New Post
        </Button>
      </nav>
      <div className="flex w-full gap-2">
        {loading ? (
          <div className="flex w-full items-center gap-2 justify-center ">
            <Spinner size={32} color="#8B5CF6" />
            <span>Loading...</span>
          </div>
        ) : user ? (
          <div className="flex flex-col gap-2 w-full">
            {' '}
            <UserInfo
              userName={user.username}
              authorAvatar={user.image_url}
              userId={user.user_id}
            ></UserInfo>
            <LogoutButton className="flex-1" />
          </div>
        ) : (
          <>
            {!pathname.startsWith('/login') && (
              <Button href="/login" className="flex-1">
                Login
              </Button>
            )}
            {!pathname.startsWith('/register') && (
              <Button href="/register" isSecondary className="flex-1">
                Register
              </Button>
            )}
          </>
        )}
      </div>
    </div>
  );
}

export default Sidebar;
