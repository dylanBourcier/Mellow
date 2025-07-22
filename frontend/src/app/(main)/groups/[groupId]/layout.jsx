import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import GroupLayoutHeader from '@/app/components/layout/GroupLayoutHeader';

export default async function GroupLayout({ children, params }) {
  const { groupId } = await params;
  return (
    <div>
      <ProtectedRoute redirectTo="/login">
        <GroupLayoutHeader groupId={groupId} />
        <div>{children}</div>
      </ProtectedRoute>
    </div>
  );
}
