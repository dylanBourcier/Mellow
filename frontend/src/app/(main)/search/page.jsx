import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import Input from '@/app/components/ui/Input';
import Label from '@/app/components/ui/Label';
import PageTitle from '@/app/components/ui/PageTitle';
import React from 'react';

export const metadata = {
  title: 'Search',
  description: 'Search for content',
};
export default function SearchPage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>
        <PageTitle>Find new friends</PageTitle>
        <div>
          <Label htmlFor="search" className="mb-2">
            Search :
          </Label>
          <Input
            id="search"
            name="search"
            placeholder="Search for users or groups..."
          ></Input>
        </div>
      </div>
    </ProtectedRoute>
  );
}
