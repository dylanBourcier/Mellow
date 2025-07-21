'use client';

import React from 'react';
import GroupCard from '../ui/GroupCard';
import PageTitle from '../ui/PageTitle';

function JoinedGroupsList(props) {
  return (
    <div className="flex flex-col gap-2.5  items-start flex-1 bg-white rounded-2xl p-4 shadow-(--box-shadow) h-full">
      <h2 className=" text-lavender-5 text-shadow-(--text-shadow)">
        Groups joined
      </h2>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
      <GroupCard></GroupCard>
    </div>
  );
}

export default JoinedGroupsList;
