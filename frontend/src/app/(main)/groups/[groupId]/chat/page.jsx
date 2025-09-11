import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import GroupConversationPage from '@/app/components/layout/GroupeConversationPage';

export default function GroupChatPage({ params }) {
  const { groupId } = params;
  return (
    <ProtectedRoute redirectTo="/login">
      <GroupConversationPage groupId={groupId} />
    </ProtectedRoute>
  );
}
