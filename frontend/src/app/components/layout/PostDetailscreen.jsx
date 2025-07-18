'use client';

import React, { use, useEffect, useState } from 'react';
import Link from 'next/link';
import UserInfo from '../ui/UserInfo';
import { icons } from '@/app/lib/icons';
import PageTitle from '../ui/PageTitle';
import Input from '../ui/Input';
import Button from '../ui/Button';
import Spinner from '../ui/Spinner';
import Image from 'next/image';
import { formatDate } from '@/app/utils/date';
import { useUser } from '@/app/context/UserContext';
import FileInput from '../ui/FileInput';

function PostDetailscreen({ postid }) {
  const [post, setPost] = useState(null);
  const [error, setError] = useState(null);
  const { user, loading } = useUser();

  useEffect(() => {
    if (!postid) return;

    const fetchPost = async () => {
      try {
        const response = await fetch(`/api/posts/${postid}`);
        if (!response.ok) {
          throw new Error(`Failed to fetch post (status: ${response.status})`);
        }

        const result = await response.json();
        if (result.status == 'error') {
          throw new Error(result.message || 'Failed to fetch post data');
        }

        if (!result?.data) {
          throw new Error('No post data returned from server');
        }

        setPost(result.data);
        post;
      } catch (err) {
        console.error('Error fetching post:', err);
        setError('Could not load the post. Please try again later.');
      }
    };

    fetchPost();
  }, [postid]);

  if (error) {
    return <div className="text-red-600">{error}</div>;
  }

  if (!post) {
    return (
      <div className="min-h-screen flex items-center gap-2 justify-center">
        <Spinner></Spinner>Loading...
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-3 w-full">
      <div className="flex flex-col gap-3 p-4 bg-white shadow-(--box-shadow) rounded-lg">
        <Link
          href={'/'}
          className="group flex items-center hover:underline hover:text-lavender-3 text-sm"
        >
          {' '}
          <span className="group-hover:animate-bounce">
            {icons['back_arrow']}
          </span>{' '}
          <span>Back to home</span>
        </Link>
        <div className="flex items-center justify-between gap-1">
          <UserInfo
            userName={post.username}
            userId={post.user_id}
            authorAvatar={post.avatar_url}
          ></UserInfo>
          <span className="font-thin text-sm">
            {formatDate(post.creation_date)}
          </span>
        </div>
        <div className="flex flex-col gap-2">
          <PageTitle className="text-left">{post.title}</PageTitle>
          <pre className="whitespace-pre-wrap break-words font-sans font-inter leading-relaxed">
            {post.content}
          </pre>
        </div>
        {post.image_url && (
          <div>
            <Image
              src={post.image_url}
              alt="Post Image"
              className="w-full h-auto rounded-lg object-cover"
              width={600}
              height={400}
              loading="lazy"
            ></Image>
          </div>
        )}
      </div>
      <div className="flex flex-col gap-3 px-2 lg:px-8 py-2.5">
        <PageTitle className="flex text-left">
          Comments ({post.comment?.length || 0})
        </PageTitle>
        {user && (
          <div className="flex gap-1 items-center">
            <Input
              type="text"
              placeholder="Post a comment..."
              className="border border-lavender-5"
            ></Input>
            <FileInput label={icons['image']} usePreview={false} isMini></FileInput>
            <Button>Comment</Button>
          </div>
        )}
        {/* {(post.comment || []).map((comment, index) => ( */}
        {/* <CommentCard key={index} comment={comment}></CommentCard> */}
        {/* ))} */}
      </div>
    </div>
  );
}

export default PostDetailscreen;
