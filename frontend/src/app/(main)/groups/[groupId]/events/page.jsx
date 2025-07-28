import React from 'react';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import EventScreen from '@/app/components/layout/EventsScreen';

async function GroupEventsPage({ params }) {
  const { groupId } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <EventScreen groupId={groupId} />
    </ProtectedRoute>
  );
}

export default GroupEventsPage;
