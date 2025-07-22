import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';

export default function GroupPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>GroupPage</div>
    </ProtectedRoute>
  );
}
