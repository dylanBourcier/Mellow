import React from 'react';
import FollowButton from './FollowButton';
import Image from 'next/image';
import Link from 'next/link';

export default function UserCard({ user }) {
  return (
    <div className="flex justify-between items-center rounded-2xl w-full p-2 bg-white shadow-(--box-shadow)">
      <Link
        href={`/user/${user.user_id}`}
        className=" group flex items-center gap-2"
      >
        <Image
          src={user?.image_url || '/img/DefaultAvatar.svg'}
          alt="User Avatar"
          width={50}
          height={50}
          className="w-12 h-12 rounded-full"
        />
        <span className="group-hover:underline">{user.username}</span>
      </Link>
      <FollowButton followStatus={user.follow_status} targetID={user.user_id} />
    </div>
  );
}
