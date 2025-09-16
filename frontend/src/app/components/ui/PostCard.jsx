import Link from 'next/link';
import React from 'react';
import Image from 'next/image';
import UserInfo from './UserInfo';
import { formatDate } from '@/app/utils/date';
import { icons } from '@/app/lib/icons';

function PostCard({ post }) {
  const {
    post_id,
    title,
    content,
    creation_date,
    username,
    avatar_url,
    group_id,
    group_name,
    comments_count,
    user_id,
    visibility,
  } = post;
  console.log(group_id);

  const formattedDate = formatDate(creation_date);
  // Ensure that the content is not too long for display
  const displayContent =
    content.length > 100 ? content.substring(0, 100) + '...' : content;
  return (
    <div className="flex flex-col w-full p-3 lg:p-5 gap-3 bg-white shadow-(--box-shadow) rounded-lg hover:-translate-y-0.5 transition-all duration-300">
      <UserInfo
        authorAvatar={avatar_url}
        userName={username}
        groupName={group_name}
        groupId={group_id}
        userId={user_id}
      ></UserInfo>
      <Link className="flex flex-col gap-3" href={`/posts/${post_id}`}>
        <div>
          <h3 className="break-all">{title}</h3>
          <span className="hidden lg:inline">{displayContent}</span>
        </div>
        <div className="flex items-center justify-between text-sm">
          <span className="flex gap-0.5 align-bottom">
            {comments_count}
            <Image
              src="/img/comments.svg"
              alt="Comments Icon"
              width={16}
              height={16}
            ></Image>
          </span>
          <div className="flex items-center text-dark-grey-lighter">
            {visibility === 'private' && (
              <span className="mr-1" title="Private Post">
                {icons['lock']}
              </span>
            )}
            {visibility === 'followers' && (
              <span className="mr-1" title="Followers only post">
                {icons['follower']}
              </span>
            )}
            {visibility === 'public' && group_id == undefined && (
              <span className="mr-1" title="Public post">
                {icons['public']}
              </span>
            )}
            {visibility === 'public' && group_id != undefined && (
              <span className="mr-1" title="Group post">
                {icons['groups']}
              </span>
            )}
            <span className="mr-1">•</span>
            <span>{formattedDate}</span>
          </div>
        </div>
      </Link>
    </div>
  );
}

export default PostCard;
