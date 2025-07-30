import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import PostsGroupContainer from '@/app/components/layout/PostsGroupContainer';

export default async function GroupPage({ params }) {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>
        <PostsGroupContainer groupId={params.groupId}></PostsGroupContainer>
      </div>
    </ProtectedRoute>
  );
}
