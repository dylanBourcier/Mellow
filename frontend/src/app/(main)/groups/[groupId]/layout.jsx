import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import GroupLayoutHeader from '@/app/components/layout/GroupLayoutHeader';

export default async function GroupLayout({ children, params }) {
  const { groupId } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <div className="flex flex-col h-auto">
        <GroupLayoutHeader groupId={groupId} />
        {children}
      </div>
    </ProtectedRoute>
  );
}
