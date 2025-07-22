import GroupCreationForm from '@/app/components/layout/GroupCreationForm';
import PageTitle from '@/app/components/ui/PageTitle';
import React from 'react';

export default function CreateGroupPage() {
  return (
    <div className='flex flex-col gap-4'>
      <PageTitle>Create Group</PageTitle>
      <GroupCreationForm></GroupCreationForm>
    </div>
  );
}
