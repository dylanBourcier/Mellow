import Link from 'next/link';
import React from 'react';

export default function GroupLink({ groupId, groupName = 'a group' }) {
  console.log('GroupLink', groupId, groupName);

  if (!groupId || groupId === 'undefined') {
    return <span className="text-lavender-5">{groupName}</span>;
  }
  return (
    <Link
      href={`/groups/${groupId}`}
      className="text-lavender-5 hover:underline"
    >
      {groupName}
    </Link>
  );
}
