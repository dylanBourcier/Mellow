'use client';

import { useUser } from '@/app/context/UserContext';
import { usePathname, useRouter } from 'next/navigation';
import toast from 'react-hot-toast';
import CustomToast from '@/app/components/ui/CustomToast';
import { icons } from '@/app/lib/icons';

export default function LogoutButton({ className, isMobile = false }) {
  const { setUser } = useUser();
  const router = useRouter();
  const pathname = usePathname();

  const handleLogout = async () => {
    try {
      await fetch('/api/auth/logout', {
        method: 'POST',
        credentials: 'include',
      });

      setUser(null); // vider le contexte
      toast.custom((t) => (
        <CustomToast message="Logged out successfully" type="success" />
      ));
      if (pathname !== '/') router.push('/login');
    } catch (err) {
      toast.custom((t) => <CustomToast message="Logout failed" type="error" />);
    }
  };

  return (
    <button
      onClick={handleLogout}
      className={`px-2 py-2 rounded-2xl bg-red-100 text-error hover:bg-red-200 border border-red-300 ${className}`}
    >
      {isMobile ? <span>{icons['logout']}</span> : <span>Logout</span>}
    </button>
  );
}
