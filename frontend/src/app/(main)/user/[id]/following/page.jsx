import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import React from 'react';
import UsersList from '@/app/components/layout/UsersList';

const metadata = {
  title: 'Following',
  description: 'List of following for a user',
};
export { metadata };

export default async function FollowingPage({ params }) {
  const { id } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <UsersList userId={id} type="following" />
    </ProtectedRoute>
  );
}
