import React from 'react';
import NotificationScreen from '@/app/components/layout/NotificationScreen';
import ProtectedRoute from '@/app/components/auth/ProtectedRoute';

export const metadata = {
  title: 'Notifications',
  description: 'Notifications page',
};

export default function NotficationsPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <NotificationScreen />
    </ProtectedRoute>
  );
}
