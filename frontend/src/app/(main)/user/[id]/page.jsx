import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import ProfileScreen from '@/app/components/layout/ProfileScreen';
import React from 'react';

export const metadata = {
  title: 'User Profile',
  description: 'User profile page',
};

export default async function UserProfilePage({ params }) {
  const { id } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <ProfileScreen userId={id}></ProfileScreen>
    </ProtectedRoute>
  );
}
