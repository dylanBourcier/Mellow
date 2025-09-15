import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import PageTitle from '@/app/components/ui/PageTitle';
import React from 'react';
import SearchScreen from '@/app/components/layout/SearchScreen';

export const metadata = {
  title: 'Search',
  description: 'Search for content',
};
export default function SearchPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>
        <PageTitle>Find new friends</PageTitle>
        <SearchScreen />
      </div>
    </ProtectedRoute>
  );
}
