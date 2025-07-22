import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import React from 'react';

export default function ConversationPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>ConversationPage</div>
    </ProtectedRoute>
  );
}
