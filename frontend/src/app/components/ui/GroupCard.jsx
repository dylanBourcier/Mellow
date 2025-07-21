import React from 'react';
import PageTitle from './PageTitle';

function GroupCard(props) {
  const GroupId = '12345'; // Placeholder for group ID
  const GroupsName = 'Group Name'; // Placeholder for group name
  const GroupsDescription = 'This is a description of the group.'; // Placeholder for group description
  const GroupMembers = 10; // Placeholder for number of members in the group
  return (
    <div className="flex flex-col p-1">
      <div className="text-dark-grey font-heading text-xl">{GroupsName}</div>
      <div className="text-dark-grey">{GroupsDescription}</div>
      <div className="text-dark-grey">Members: {GroupMembers}</div>
    </div>
  );
}

export default GroupCard;
