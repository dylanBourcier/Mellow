'use client';
import { usePathname } from 'next/navigation';
import React, { useEffect } from 'react';
import Navlink from '../ui/Navlink';
import Image from 'next/image';
import Button from '../ui/Button';
import { useUser } from '@/app/context/UserContext';

import UserInfo from '../ui/UserInfo';
import LogoutButton from '../ui/LogoutButton';
import Spinner from '../ui/Spinner';

function Sidebar(props) {
  const pathname = usePathname();
  const { user, loading } = useUser();
  const [unreadCount, setUnreadCount] = React.useState(0);
  useEffect(() => {
    if (user) {
      setUnreadCount(user.unread_count);
    } else {
      setUnreadCount(0);
    }
  }, [user?.unread_count]);

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
        <Navlink
          href="/"
          icon="home"
          disabled={!user}
          isActive={pathname === '/'}
        >
          Home
        </Navlink>
        <Navlink
          href="/search"
          icon="search"
          disabled={!user}
          isActive={pathname.startsWith('/search')}
        >
          Search
        </Navlink>
        <Navlink
          href="/messages"
          icon="messages"
          disabled={!user}
          isActive={pathname.startsWith('/messages')}
        >
          Messages
          {user && user.unread_count > 0 && (
            <span className="absolute top-3 left-4 inline-flex items-center justify-center px-2 py-1 text-xs font-bold leading-none text-red-100 bg-red-600 rounded-full transform translate-x-1/2 -translate-y-1/2">
              {unreadCount}
            </span>
          )}
        </Navlink>
        <Navlink
          href="/groups"
          icon="groups"
          disabled={!user}
          isActive={pathname.startsWith('/groups')}
        >
          Groups
        </Navlink>
        <Navlink
          href="/notifications"
          icon="notifications"
          disabled={!user}
          isActive={pathname.startsWith('/notifications')}
        >
          Notifications
        </Navlink>
        <Navlink
          href="/profile"
          icon={user ? undefined : 'profile'}
          img={user ? user.image_url || '/img/DefaultAvatar.svg' : undefined}
          disabled={!user}
          isActive={pathname.startsWith('/profile')}
        >
          {user ? user.lastname + ' ' + user.firstname : 'Profile'}
        </Navlink>
        <Button className="mt-6 w-full" disabled={!user} href="/posts/create">
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
            <UserInfo
              userName={user.username}
              authorAvatar={
                user.image_url && user.image_url !== ''
                  ? user.image_url
                  : '/img/DefaultAvatar.svg'
              }
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
