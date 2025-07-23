import React from 'react';

export default function CreateEventPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>CreateEventPage</div>
    </ProtectedRoute>
  );
}
