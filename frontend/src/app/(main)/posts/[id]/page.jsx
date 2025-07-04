import React from 'react';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';
import UserInfo from '@/app/components/ui/UserInfo';
import PageTitle from '@/app/components/ui/PageTitle';
import Input from '@/app/components/ui/Input';
import Button from '@/app/components/ui/Button';
import CommentCard from '@/app/components/ui/CommentCard';

const metadata = {
  title: 'Post Details - Mellow',
};
export { metadata };

export default function PostDetailsPage() {
  const Created_at = '15h 2023-10-01';
  const postContent="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.\nUt enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.\nDuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.\nExcepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.";
  const postTitle = "Post’s title with some useless characters to see the result in a more realistic way";
  const postId = 1;
  const nbComments = 5;
  return (
  <div className='flex flex-col gap-3'>
  <div className='flex flex-col gap-3 p-4 bg-white shadow-(--box-shadow) rounded-lg'>
    <Link href={"/"} className='group flex items-center hover:underline hover:text-lavender-3 text-sm'> <span className='group-hover:animate-bounce'>{icons["back_arrow"]}</span> <span>Back to home</span></Link>
    <div className='flex items-center justify-between gap-1'>
      <UserInfo></UserInfo><span className='font-thin text-sm'>{Created_at}</span>
    </div>
    <div className='flex flex-col gap-2'>
    <PageTitle className='text-left'>{postTitle}</PageTitle><pre className='whitespace-pre-wrap break-words font-sans font-inter leading-relaxed'>{postContent}</pre>
    </div>
  </div>
  <div className='flex flex-col gap-3 px-2 lg:px-8 py-2.5'>
    <PageTitle className='flex text-left'>Comments ({nbComments})</PageTitle>
    <div className='flex gap-1 items-center'>
    <Input type="text" placeholder="Post a comment..."className='border border-lavender-5'></Input>
    <Button>Comment</Button>
    </div>
    <CommentCard></CommentCard>
    <CommentCard></CommentCard>
    <CommentCard></CommentCard>
    <CommentCard></CommentCard>
  </div>
  </div>
);
}
