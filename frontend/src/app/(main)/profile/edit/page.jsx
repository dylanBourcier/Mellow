import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import EditProfile from '@/app/components/layout/EditProfile';
import React from 'react';

function EditPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div className="flex flex-col items-center justify-center">
        <EditProfile />
      </div>
    </ProtectedRoute>
  );
}

export default EditPage;
