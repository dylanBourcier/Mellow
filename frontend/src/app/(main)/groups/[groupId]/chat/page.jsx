import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import GroupConversationPage from '@/app/components/layout/GroupeConversationPage';

export default async function GroupChatPage({ params }) {
  const { groupId } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <GroupConversationPage groupId={groupId} />
    </ProtectedRoute>
  );
}
