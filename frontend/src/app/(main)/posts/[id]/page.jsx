import React, { use } from 'react';
import PostDetailscreen from '@/app/components/layout/PostDetailscreen';

const metadata = {
  title: 'Post Details - Mellow',
};
export { metadata };

export default async function PostDetailsPage(props) {
  const { id } = await props.params;
  return (
    <div>
      <PostDetailscreen postid={id}></PostDetailscreen>
    </div>
  );
}
