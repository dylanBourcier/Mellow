import Image from 'next/image';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';
import React from 'react';

export default function UserInfo({
  userName = 'username',
  authorAvatar = '/img/DefaultAvatar.svg',
  groupId = null,
  userId = 'default-user',
  groupName = null,
}) {
  return (
    <div className="flex items-center gap-2">
      <Image
        src={authorAvatar || '/img/DefaultAvatar.svg'}
        width={40}
        height={40}
        alt="Author Avatar"
        className="w-10 h-10 rounded-full"
      />
      <span className="flex gap-1">
        <Link className="hover:underline" href={`/user/${userId}`}>
          {userName}
        </Link>
        {groupId && groupName && (
          <>
            {' · '}
            <Link
              className="font-semibold flex items-center gap-1 hover:underline"
              href={`/groups/${groupId}`}
            >
              {icons?.['groups'] || '🏠'}
              {groupName}
            </Link>
          </>
        )}
      </span>
    </div>
  );
}
