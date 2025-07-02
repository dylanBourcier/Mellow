'use client';
import { usePathname } from 'next/navigation';
import React from 'react';
import Navlink from '../ui/Navlink';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';

function SidebarMobile() {
  const pathname = usePathname();
  return (
    <div className="flex w-full lg:hidden h-16 fixed -bottom-0 bg-white justify-evenly items-center">
      <Navlink href="/" icon="home" isActive={pathname === '/'}></Navlink>
      <Navlink
        href="/search"
        icon="search"
        isActive={pathname.startsWith('/search')}
      ></Navlink>
      <Navlink
        href="/messages"
        icon="messages"
        isActive={pathname.startsWith('/messages')}
      ></Navlink>

      <Link
        href="/posts/create"
        className="flex items-center justify-center w-12 h-12"
      >
        <span className="flex-shrink-0 text-lavender-3 hover:text-lavender-4 hover:scale-110 ease-out transition-all duration-200">
          {icons['createPost']}
        </span>
      </Link>

      <Navlink
        href="/groups"
        icon="groups"
        isActive={pathname.startsWith('/groups')}
      ></Navlink>
      <Navlink
        href="/notifications"
        icon="notifications"
        isActive={pathname.startsWith('/notifications')}
      ></Navlink>
      <Navlink
        href="/profile"
        icon="profile"
        isActive={pathname.startsWith('/profile')}
      ></Navlink>
    </div>
  );
}

export default SidebarMobile;
