'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import React from 'react';

export default function GroupNavlink({
  href = '/',
  children = 'ico',
  isMember = true,
}) {
  const pathname = usePathname();
  const isActive = pathname === href;

  return (
    <Link
      href={isMember ? href : '#'}
      className={`rounded-lg flex gap-1 p-3 ${
        isActive ? 'text-lavender-3 bg-lavender-6' : ''
      } ${
        !isMember
          ? 'cursor-not-allowed opacity-50'
          : 'hover:text-lavender-3 transition-colors duration-100'
      }`}
      onClick={(e) => {
        if (!isMember) e.preventDefault();
      }}
    >
      {children}
    </Link>
  );
}
