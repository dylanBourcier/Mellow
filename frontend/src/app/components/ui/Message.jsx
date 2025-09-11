'use client';

import React from 'react';
import { useUser } from '@/app/context/UserContext';
import { formatDateTime } from '@/app/utils/date'; // Assuming you have a utility function for formatting dates
import Image from 'next/image';

export default function Message({ message,type }) {
  const { sender_id, content, creation_date } = message; // Assuming message has these properties
  const { user } = useUser(); // Assuming useUser is defined in your context
  const isSender = sender_id === user.user_id; // Check if the logged-in user is the sender

  console.log(message);
  
  return (
    <div
      className={`flex flex-col gap-1 ${
        isSender ? 'self-end flex-row-reverse' : 'self-start'
      }`}
    >{!isSender &&type=="group"&&(<div className='flex gap-1'><div>{message.image_url ||<Image src={message.image_url || "/img/DefaultAvatar.svg"} width={24} height={24} alt="User Avatar" className="w-6 h-6 rounded-full " />}</div><div>{message.username}</div></div>)}
    <div className={`flex gap-2 items-end ${isSender ? 'self-end flex-row-reverse' : 'self-start'}`}>
      <div
        className={`p-3 w-fit max-w-[500px] rounded-2xl shadow-2xs break-words ${
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
