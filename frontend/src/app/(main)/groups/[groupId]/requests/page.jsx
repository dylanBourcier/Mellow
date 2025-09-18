import ProtectedRoute from '@/app/components/auth/ProtectedRoute';
import GroupRequestsPanel from '@/app/components/layout/GroupRequestsPanel';

export default async function GroupRequestsPage({ params }) {
  const { groupId } = await params;
  return (
    <ProtectedRoute redirectTo="/login">
      <GroupRequestsPanel groupId={groupId} />
    </ProtectedRoute>
  );
}

