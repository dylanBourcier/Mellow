import React, { use } from 'react';
import PostDetailscreen from '@/app/components/layout/PostDetailscreen';

const metadata = {
  title: 'Post Details - Mellow',
};
export { metadata };

export default async function PostDetailsPage(props) {
  const params=await props.params;
  const id= params.id;
  return(
    <div>
      <PostDetailscreen postid={id} ></PostDetailscreen>
    </div>
  )

}
