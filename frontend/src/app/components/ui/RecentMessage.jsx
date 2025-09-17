import Image from 'next/image';
import React from 'react';
import { formatDateTime } from '@/app/utils/date';
import Link from 'next/link';
import previewMessage from '@/app/utils/text';

export default function RecentMessage({ conversation }) {
  return (
    <Link
      className="flex w-full gap-2 p-3 relative"
      href={`/messages/${conversation.user_id}`}
    >
      <div className="relative flex min-w-12 min-h-12">
        <Image
          src={conversation.avatar || '/img/DefaultAvatar.svg'}
          alt={conversation.username}
          width={48}
          height={48}
          className="rounded-full w-12 h-12"
          unoptimized
        />
        {conversation.unread_count > 0 && (
          <span
            className="bg-red-500 w-4 h-4 text-xs flex items-center justify-center text-white rounded-full absolute top-0 right-0"
            style={{ fontFamily: 'inherit', lineHeight: '1', fontWeight: 600 }}
          >
            {Math.round(conversation.unread_count)}
          </span>
        )}
      </div>
      <div className="flex flex-col">
        <span>
          {conversation.username}{' '}
          <span className="font-extralight text-sm">
            {formatDateTime(conversation.last_sent_at)}
          </span>
        </span>
        <span className="text-sm text-gray-500 italic break-all">
          {previewMessage(conversation.last_message, 40)}
        </span>
      </div>
    </Link>
  );
}
