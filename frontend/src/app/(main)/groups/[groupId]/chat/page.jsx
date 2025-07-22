import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';

export default function GroupChatPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>GroupChatPage</div>
    </ProtectedRoute>
  );
}
