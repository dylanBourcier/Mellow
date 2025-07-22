import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';

function GroupEventsPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>GroupEventsPage</div>
    </ProtectedRoute>
  );
}

export default GroupEventsPage;
