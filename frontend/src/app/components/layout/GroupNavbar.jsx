'use client';

import React from 'react';
import GroupNavlink from '../ui/GroupNavlink';
import { icons } from '@/app/lib/icons';

export default function GroupNavbar({ groupId, isMember }) {
  return (
    <nav className="flex w-full items-center justify-center gap-1 lg:gap-8">
      <GroupNavlink href={`/groups/${groupId}`}>
        {icons['posts']}Posts
      </GroupNavlink>
      <GroupNavlink isMember={isMember} href={`/groups/${groupId}/events`}>
        {icons['events']}Events
      </GroupNavlink>
      <GroupNavlink isMember={isMember} href={`/groups/${groupId}/chat`}>
        {icons['messages']}Chat
      </GroupNavlink>
    </nav>
  );
}
