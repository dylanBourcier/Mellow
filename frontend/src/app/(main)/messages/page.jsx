import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import React from 'react';

export default function MessagesListPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>MessagesListPage</div>
    </ProtectedRoute>
  );
}
