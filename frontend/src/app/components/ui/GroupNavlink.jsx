'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import React from 'react';

export default function GroupNavlink({ href = '/', children = 'ico' }) {
  const pathname = usePathname();
  const isActive = pathname === href;
  return (
    <Link
      href={href}
      className={`rounded-lg flex gap-1 p-3 ${
        isActive
          ? 'text-lavender-3 bg-lavender-6'
          : 'hover:text-lavender-3 transition-colors duration-100'
      }`}
    >
      {children}
    </Link>
  );
}
