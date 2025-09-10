'use client';

import React from 'react';
import { useUser } from '@/app/context/UserContext';
import { formatDateTime } from '@/app/utils/date'; // Assuming you have a utility function for formatting dates

export default function Message({ message }) {
  const { sender_id, content, creation_date } = message; // Assuming message has these properties
  const { user } = useUser(); // Assuming useUser is defined in your context
  const isSender = sender_id === user.user_id; // Check if the logged-in user is the sender

  return (
    <div
      className={`flex items-end gap-2 ${
        isSender ? 'self-end flex-row-reverse' : 'self-start'
      }`}
    >
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
  );
}
