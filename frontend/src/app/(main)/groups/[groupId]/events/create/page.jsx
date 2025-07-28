import React from 'react';
import CreateEventForm from '@/app/components/layout/CreateEventForm';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';

export default async function CreateEventPage({ params }) {
  const { groupId } = await params;

  return (
    <ProtectedRoute redirectTo="/login">
      <CreateEventForm groupId={groupId} />
    </ProtectedRoute>
  );
}
