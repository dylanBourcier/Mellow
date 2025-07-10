import Link from "next/link";
import React from "react";
import Image from "next/image";
import UserInfo from "./UserInfo";

function PostCard({ postInfos }) {
  // You can replace these with props or state as needed
  const {
    postId,
    postTitle,
    postContent,
    authorAvatar,
    userName,
    date,
    Comments,
  } = postInfos;
  return (
    <div>
      <div className="flex flex-col p-6 gap-3 bg-white shadow-(--box-shadow) rounded-lg hover:-translate-y-0.5 transition-all duration-300">
        <UserInfo
          authorAvatar={authorAvatar}
          userName={userName}
          userId={24}
          groupId={45}
          groupName="test"
        ></UserInfo>
        <Link href={`/posts/${postId}`}>
          <div>
            <h3>{postTitle}</h3>
            <span>{postContent}</span>
          </div>
          <div className="flex items-center justify-between text-sm">
            <span className="flex gap-0.5 align-bottom">
              {Comments}
              <Image
                src="/img/comments.svg"
                alt="Comments Icon"
                width={16}
                height={16}
              ></Image>
            </span>
            <span>{date}</span>
          </div>
        </Link>
      </div>
    </div>
  );
}

export default PostCard;
