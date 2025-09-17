'use client';

import React from 'react';
import { useUser } from '@/app/context/UserContext';
import { formatDateTime } from '@/app/utils/date'; // Assuming you have a utility function for formatting dates
import Image from 'next/image';
import Link from 'next/link';

export default function Message({ message, type }) {
  const { sender_id, content, creation_date } = message; // Assuming message has these properties
  const { user } = useUser(); // Assuming useUser is defined in your context
  const isSender = sender_id === user.user_id; // Check if the logged-in user is the sender
  return (
    <div className={`flex flex-col gap-1`}>
      {!isSender && type == 'group' && (
        <Link
          className="flex gap-1 hover:underline"
          href={`/user/${message.sender_id}`}
        >
          <div>
            <Image
              src={message.image_url || '/img/DefaultAvatar.svg'}
              width={24}
              height={24}
              alt="User Avatar"
              className="w-6 h-6 rounded-full "
            />
          </div>
          <div>{message.username}</div>
        </Link>
      )}
      {isSender && type == 'group' && (
        <div className="flex self-end">
          <Image
            src={user.image_url || '/img/DefaultAvatar.svg'}
            width={24}
            height={24}
            alt={user.username + ' avatar'}
            className="w-6 h-6 rounded-full "
          />
        </div>
      )}
      <div
        className={`flex gap-2 items-end ${
          isSender ? 'self-end flex-row-reverse' : 'self-start'
        }`}
      >
        <div
          className={`p-3 w-fit lg:max-w-[500px] max-w-[80%] rounded-2xl shadow-2xs break-all ${
            isSender ? 'bg-lavender-1' : 'bg-white'
          }`}
        >
          {content}
        </div>
        <span className="text-sm whitespace-nowrap">
          {formatDateTime(creation_date)}
        </span>
      </div>
    </div>
  );
}
