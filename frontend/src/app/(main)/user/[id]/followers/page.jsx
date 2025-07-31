import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import React from 'react';
import UsersList from '@/app/components/layout/UsersList';

const metadata = {
  title: 'Followers',
  description: 'List of followers for a user',
};
export { metadata };

export default async function FollowersPage({ params }) {
  const { id } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <UsersList userId={id} type="followers" />
    </ProtectedRoute>
  );
}
