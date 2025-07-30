import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import PostsGroupContainer from '@/app/components/layout/PostsGroupContainer';

export default async function GroupPage({ params }) {
  const { groupId } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <div>
        <PostsGroupContainer groupId={groupId}></PostsGroupContainer>
      </div>
    </ProtectedRoute>
  );
}
