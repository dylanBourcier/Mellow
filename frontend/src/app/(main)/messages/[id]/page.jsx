import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import React from 'react';
import UserConversation from '@/app/components/layout/UserConversation';

const metadata = {
  title: 'Conversation',
  description: 'Conversation page',
};
export { metadata };

export default async function ConversationPage({ params }) {
  const { id } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <UserConversation id={id} />
    </ProtectedRoute>
  );
}
