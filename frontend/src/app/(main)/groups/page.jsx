import JoinedGroupsList from '@/app/components/layout/JoinedGroupsList';
import GroupsList from '@/app/components/layout/GroupsList';
import React from 'react';

export default function GroupsPage() {
  return (
    <div className="flex flex-col lg:flex-row h-full gap-2">
      <JoinedGroupsList />
      <GroupsList />
    </div>
  );
}
