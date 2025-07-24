import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import EventScreenPage from '@/app/components/layout/EventScreenPage';

function GroupEventsPage(event) {
  return (
    <ProtectedRoute redirectTo="/login">
      <EventScreenPage event={event}/>
    </ProtectedRoute>
  );
}

export default GroupEventsPage;
