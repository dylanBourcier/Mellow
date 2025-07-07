import React from 'react';
import PostDetailscreen from '@/app/components/layout/PostDetailscreen';

const metadata = {
  title: 'Post Details - Mellow',
};
export { metadata };

export default function PostDetailsPage() {
  return(
    <div>
      <PostDetailscreen></PostDetailscreen>
    </div>
  )

}
