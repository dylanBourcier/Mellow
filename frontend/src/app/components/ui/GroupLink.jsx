import Link from 'next/link';
import React from 'react';

export default function GroupLink({ groupId, groupName }) {
  return (
    <Link
      href={`/groups/${groupId}`}
      className="text-lavender-5 hover:underline"
    >
      {groupName}
    </Link>
  );
}
