import Image from 'next/image';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';
import React from 'react';

export default function UserInfo({
    userName = 'johndoe',
    authorAvatar = '/img/DefaultAvatar.png',
    groupId = null,
    userId = 'default-user',
    groupName = null,
}) {
    const safeUserName = userName || 'johndoe';
    const safeAuthorAvatar = authorAvatar || '/img/DefaultAvatar.png';
    const safeUserId = userId || 'default-user';
    const safeGroupId = groupId || null;
    const safeGroupName = groupName || null;

    return (
        <div className="flex items-center gap-2">
            <Image
                src={safeAuthorAvatar}
                width={32}
                height={32}
                alt="Author Avatar"
                className="w-8 h-8 rounded-full"
            />
            <span className='flex gap-1'>
                <Link className="hover:underline" href={`/user/${safeUserId}`}>
                    {safeUserName}
                </Link>
                {safeGroupId && safeGroupName && (
                    <>
                        {' · '}
                        <Link 
                            className='font-semibold flex items-center gap-1 hover:underline' 
                            href={`/groups/${safeGroupId}`}
                        >
                            {icons?.["groups"] || '🏠'}
                            {safeGroupName}
                        </Link>
                    </>
                )}
            </span>
        </div>
    );
}