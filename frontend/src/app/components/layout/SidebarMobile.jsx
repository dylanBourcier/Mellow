'use client';
import { usePathname } from 'next/navigation';
import React from 'react';
import Navlink from '../ui/Navlink';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';
import { useUser } from '@/app/context/UserContext';
import LogoutButton from '../ui/LogoutButton';
import Button from '../ui/Button';

function SidebarMobile() {
  const pathname = usePathname();
  const { user, loading } = useUser();
  return (
    <>
      {loading ? null : user ? (
        <LogoutButton isMobile className="absolute right-4 lg:hidden z-50" />
      ) : (
        <Link
          href="/login"
          className="absolute right-4 text-lavender-3 bg-lavender-6 rounded-2xl p-2 border border-lavender-5 shadow-(--box-shadow) lg:hidden z-50"
        >
          {icons['signin']}
        </Link>
      )}

      <div className="flex w-full lg:hidden h-16 fixed bottom-0 bg-white justify-evenly items-center">
        <Navlink href="/" icon="home" isActive={pathname === '/'} />
        <Navlink
          href="/search"
          icon="search"
          isActive={pathname.startsWith('/search')}
        />
        <div className="relative">
          <Navlink
            href="/messages"
            icon="messages"
            disabled={!user}
            isActive={pathname.startsWith('/messages')}
          />
          {user && user.unread_count > 0 && (
            <span className="absolute top-3 left-4 inline-flex items-center justify-center px-2 py-1 text-xs font-bold leading-none text-red-100 bg-red-600 rounded-full transform translate-x-1/2 -translate-y-1/2">
              {user.unread_count}
            </span>
          )}
        </div>

        <Link
          href={user ? '/posts/create' : '#'}
          className={`flex items-center justify-center w-12 h-12 ${
            user ? '' : 'pointer-events-none opacity-50'
          }`}
        >
          <span
            className={`flex-shrink-0 ${
              user
                ? 'text-lavender-3 hover:text-lavender-4 hover:scale-110 ease-out transition-all duration-200'
                : 'text-dark-grey'
            }`}
          >
            {icons['createPost']}
          </span>
        </Link>

        <Navlink
          href="/groups"
          icon="groups"
          disabled={!user}
          isActive={pathname.startsWith('/groups')}
        />
        <Navlink
          href="/notifications"
          icon="notifications"
          disabled={!user}
          isActive={pathname.startsWith('/notifications')}
        />
        {user ? (
          <Navlink
            href="/profile"
            img={
              user.image_url && user.image_url !== ''
                ? user.image_url
                : '/img/DefaultAvatar.svg'
            }
            isActive={pathname.startsWith('/profile')}
          />
        ) : (
          <Navlink
            href="/login"
            icon="signin"
            isActive={
              pathname.startsWith('/login') || pathname.startsWith('/register')
            }
          ></Navlink>
        )}
      </div>
    </>
  );
}

export default SidebarMobile;
