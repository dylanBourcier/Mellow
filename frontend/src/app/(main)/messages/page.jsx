import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import React from 'react';
import ConversationList from '@/app/components/layout/ConversationList';
import PageTitle from '@/app/components/ui/PageTitle';

export const metadata = {
  title: 'Messages',
  description: 'View and manage your messages',
};


export default function MessagesListPage() {
  
  return (
    <ProtectedRoute redirectTo="/login">
      <PageTitle>Messages</PageTitle>
      <ConversationList />
    </ProtectedRoute>
  );
}
