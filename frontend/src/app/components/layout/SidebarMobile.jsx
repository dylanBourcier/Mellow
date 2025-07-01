import React from 'react';
import Navlink from '../ui/Navlink';
import Link from 'next/link';
import Image from 'next/image';
import { icons } from '@/app/lib/icons';

function SidebarMobile() {
  //
  return (
    <div className="flex w-full md:hidden h-16 absolute -bottom-0 bg-white justify-evenly items-center">
      <Navlink href="/" icon="home"></Navlink>
      <Navlink href="/search" icon="search"></Navlink>
      <Navlink href="/messages" icon="messages"></Navlink>

      <Link
        href="/posts/create"
        className="flex items-center justify-center w-12 h-12"
      >
        <span className="flex-shrink-0 text-lavender-3 hover:text-lavender-4 hover:scale-110 ease-out transition-all duration-200">
          {icons['createPost']}
        </span>
      </Link>

      <Navlink href="/groups" icon="groups"></Navlink>
      <Navlink href="/notifications" icon="notifications"></Navlink>
      <Navlink href="/profile" icon="profile"></Navlink>
    </div>
  );
}

export default SidebarMobile;
