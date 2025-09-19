'use client';

import React from 'react';
import Button from '../ui/Button';
import Link from 'next/link';
import Image from 'next/image';
import { formatDate } from '@/app/utils/date';
import GroupLink from './GroupLink';

function NotificationsCard({ notification, onAccept, onDecline }) {
  const { type, sender_username, sender_avatar_url, creation_date } =
    notification;

  // Déterminer dynamiquement le message
  let message;
  switch (type) {
    case 'follow_request':
      message = 'has sent you a follow request';
      break;
    case 'accepted_follow_request':
      message = 'has accepted your follow request';
      break;
    case 'rejected_follow_request':
      message = 'has rejected your follow request';
      break;
    case 'accepted_group_request':
      message = (
        <>
          has accepted your request to join{' '}
          <GroupLink
            groupId={notification.group_id}
            groupName={notification.group_name}
          />
        </>
      );
      break;
    case 'rejected_group_request':
      message = (
        <>
          has rejected your request to join{' '}
          <GroupLink
            groupId={notification.group_id}
            groupName={notification.group_name}
          />
        </>
      );
      break;
    case 'accepted_group_invite':
      message = (
        <>
          has accepted your invitation to join{' '}
          <GroupLink
            groupId={notification.group_id}
            groupName={notification.group_name}
          />
        </>
      );
      break;
    case 'rejected_group_invite':
      message = (
        <>
          has rejected your invitation to join{' '}
          <GroupLink
            groupId={notification.group_id}
            groupName={notification.group_name}
          />
        </>
      );
      break;
    case 'new_follower':
      message = 'has followed you';
      break;
    case 'event_created':
      message = (
        <>
          has created an event in{' '}
          <GroupLink
            groupId={notification.group_id}
            groupName={notification.group_name}
          />
        </>
      );
      break;
    case 'group_invite':
      message = (
        <>
          has invited you to join{' '}
          <GroupLink
            groupId={notification.group_id}
            groupName={notification.group_name}
          />
        </>
      );
      break;
    case 'group_request':
      message = (
        <>
          has requested to join{' '}
          <GroupLink
            groupId={notification.group_id}
            groupName={notification.group_name}
          />
        </>
      );
      break;
    default:
      message = 'has sent you a notification';
  }

  const showActions =
    type === 'follow_request' ||
    type === 'group_request' ||
    type === 'group_invite';

  return (
    <div className="bg-white flex justify-between items-center shadow-md rounded-lg w-full px-3 py-3">
      <div className="flex items-center gap-3">
        <Image
          width={44}
          height={44}
          src={sender_avatar_url || '/img/DefaultAvatar.svg'}
          alt="Avatar"
          className="w-11 h-11 rounded-full"
        />

        <div className="flex flex-col text-sm text-inter text-dark-grey gap-1">
          <span>
            <Link
              href={`/user/${notification.sender_id}`}
              className="text-lavender-5 italic hover:underline"
            >
              {sender_username}
            </Link>{' '}
            <span className="italic">{message}</span>
          </span>
          {showActions &&
            (notification.seen === false ? (
              <div className="flex gap-2">
                <Button onClick={onAccept}>Accept</Button>
                <Button
                  onClick={onDecline}
                  className="bg-light-grey text-lavender-5 hover:bg-white border border-lavender-5"
                >
                  Reject
                </Button>
              </div>
            ) : (
              <span className="text-xs text-gray-500">
                Already answered to this request
              </span>
            ))}
        </div>
      </div>

      <span className="text-sm text-inter text-dark-grey">
        {formatDate(creation_date)}
      </span>
    </div>
  );
}

export default NotificationsCard;
