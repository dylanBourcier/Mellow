import GroupCreationForm from '@/app/components/layout/GroupCreationForm';
import PageTitle from '@/app/components/ui/PageTitle';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import React from 'react';

export default function CreateGroupPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div className="flex flex-col gap-4">
        <PageTitle>Create Group </PageTitle>
        <GroupCreationForm></GroupCreationForm>
      </div>
    </ProtectedRoute>
  );
}
