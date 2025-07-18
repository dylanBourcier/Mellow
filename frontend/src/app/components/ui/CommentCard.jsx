import React from 'react';
import UserInfo from '@/app/components/ui/UserInfo';
import Image from 'next/image';
import { formatDate } from '@/app/utils/date';

function CommentCard({ comment }) {
  const { content, creation_date, user_id, username, avatar_url, image_url } =
    comment;
  const formattedDate = formatDate(creation_date);
  return (
    <div className="flex flex-col gap-1 bg-white p-4 shadow-(--box-shadow) rounded-lg">
      <div className="flex items-center justify-between gap-3">
        <UserInfo
          userId={user_id}
          authorAvatar={avatar_url}
          userName={username}
        ></UserInfo>
        <span className="font-thin text-sm">{formattedDate}</span>
      </div>
      <pre className="whitespace-pre-wrap break-words font-sans font-inter">
        {content}
      </pre>
      {image_url && (
        <Image
          src={image_url}
          alt="Comment Image"
          className="w-full h-auto rounded-xl object-cover"
          width={600}
          height={400}
          loading="lazy"
        />
      )}
    </div>
  );
}

export default CommentCard;
