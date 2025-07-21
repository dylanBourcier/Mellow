'use client';

import React from 'react';
import GroupCard from '../ui/GroupCard';
import PageTitle from '../ui/PageTitle';

function JoinedGroupsList(props) {
  return (
    <div className="flex flex-col gap-2.5 items-start flex-2 bg-white rounded-2xl p-4 shadow-(--box-shadow) h-full">
      <PageTitle className="flex gap-2.5">Groups joined</PageTitle>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
    </div>
  );
}

export default JoinedGroupsList;
