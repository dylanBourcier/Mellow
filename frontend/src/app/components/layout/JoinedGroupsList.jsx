'use client';

import React from 'react';
import GroupCard from '../ui/GroupCard';
import PageTitle from '../ui/PageTitle';
import Button from '../ui/Button';

function JoinedGroupsList(props) {
  return (
    <div className="flex flex-col gap-2.5  items-start flex-1 bg-white rounded-2xl p-4 shadow-(--box-shadow) h-full">
      <h2 className=" text-lavender-5 text-shadow-(--text-shadow)">
        Groups joined
      </h2>
      <div className="flex w-full">
        <Button href={'/groups/create'} className="w-full">
          Create Group
        </Button>
      </div>

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
