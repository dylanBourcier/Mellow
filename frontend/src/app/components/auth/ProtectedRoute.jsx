'use client';

import { useUser } from '@/app/context/UserContext';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export default function ProtectedRoute({ children, redirectTo = '/login' }) {
  const { user, loading } = useUser();
  const router = useRouter();

  useEffect(() => {
    if (!loading && !user) {
      router.replace(redirectTo);
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
