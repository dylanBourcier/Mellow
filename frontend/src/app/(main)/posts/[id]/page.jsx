import React, { use } from 'react';
import PostDetailscreen from '@/app/components/layout/PostDetailscreen';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';

const metadata = {
  title: 'Post Details - Mellow',
};
export { metadata };

export default async function PostDetailsPage(props) {
  const { id } = await props.params;
  return (
    <ProtectedRoute redirectTo="/login">
      <div>
        <PostDetailscreen postid={id}></PostDetailscreen>
      </div>
    </ProtectedRoute>
  );
}
