'use client';

import { useUser } from '@/app/context/UserContext';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';
import Spinner from '../ui/Spinner';
import toast from 'react-hot-toast';
import CustomToast from '../ui/CustomToast';

export default function ProtectedRoute({ children, redirectTo = '/login' }) {
  const { user, loading } = useUser();
  const router = useRouter();

  useEffect(() => {
    if (!loading && !user) {
      // If the user is not authenticated, redirect to the login page with a toast message
      router.replace(redirectTo);
      toast.custom((t) => (
        <CustomToast
          type={'error'}
          message={'You must be logged in to access this page'}
        ></CustomToast>
      ));
    }
  }, [user, loading, redirectTo]);

  if (loading || !user) {
    return (
      <div className="min-h-screen flex items-center  gap-2 justify-center">
        <Spinner size={32} color="#8B5CF6" />
        Loading...
      </div>
    );
  }

  return children;
}
