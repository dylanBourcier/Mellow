'use client';

import React from 'react';

import Button from '../ui/Button';
import Link from 'next/link';
import Image from 'next/image';
import { formatDate } from '@/app/utils/date';

function NotificationsCard({ notification, onAccept, onDecline }) {
  // Déterminer dynamiquement le message
  let message = '';

  if (notification.type === 'follow_request') {
    message = (
      <span>
        <Link href="/profile">
          <span className="text-lavender-5 italic hover:underline">
            {notification.sender_username}
          </span>
        </Link>{' '}
        <span className=" italic">has sent you a follow request</span>
      </span>
    );
  } else if (notification.type === 'new_follower') {
    message = (
      <span>
        <Link href="/profile">
          <span className="text-lavender-5 italic hover:underline">
            {notification.sender_username}
          </span>
        </Link>{' '}
        <span className=" italic">has followed you</span>
      </span>
    );
  }

  return (
    <div className="bg-white flex justify-between items-center shadow-md rounded-lg w-full">
      <div className="flex items-center px-3 py-3">
        <div className="flex items-center">
          <Image
            width={44}
            height={44}
            src={notification.sender_avatar_url || '/img/DefaultAvatar.svg'}
            alt="Avatar"
            className="w-11 h-11 rounded-full inline-block mr-2"
          />
        </div>
        <div className="flex flex-col gap-1 text-inter text-dark-grey text-sm">
          <span>{message}</span>
          {notification.type === 'follow_request' && (
            <div className="flex gap-2 ">
              <Button onClick={onAccept}>Accept</Button>
              <Button
                onClick={onDecline}
                className="bg-light-grey text-lavender-5 hover:bg-white border border-lavender-5"
              >
                Reject
              </Button>
            </div>
          )}
        </div>
      </div>
      <div className="flex flex-row items-end text-inter text-dark-grey text-sm mr-2">
        <span>{formatDate(notification.creation_date)}</span>
      </div>
    </div>
  );
}

export default NotificationsCard;
