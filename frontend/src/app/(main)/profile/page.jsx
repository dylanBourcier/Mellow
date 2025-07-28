import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import ProfileScreen from '@/app/components/layout/ProfileScreen';
import React from 'react';

export const metadata = {
  title: 'Profile',
  description: 'User Profile Page',
};

export default function ProfilePage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>
        <ProfileScreen />
      </div>
    </ProtectedRoute>
  );
}
