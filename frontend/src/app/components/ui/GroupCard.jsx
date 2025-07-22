import React from 'react';
import Link from 'next/link';
import Button from '../ui/Button';

function GroupCard({ props, withButton = false }) {
  const GroupId = '12345'; // Placeholder for group ID
  const GroupsName = 'Group Name'; // Placeholder for group name
  const GroupsDescription =
    'This is a description of the group Which can be long, but like veryyyyy long and its annoying for the display.'; // Placeholder for group description
  const SlicedGroupDescription =
    GroupsDescription.length > 100
      ? GroupsDescription.slice(0, 100) + '...'
      : GroupsDescription;
  const GroupMembers = 10; // Placeholder for number of members in the group
  return (
    <Link
      href={'/groups/' + GroupId}
      className="flex hover:bg-light-grey items-center rounded-lg w-full transition-all duration-100 ease-in-out"
    >
      <div className="flex flex-col p-2 group">
        <div className="text-dark-grey font-heading text-xl group-hover:text-lavender-5 transition-all duration-100 ease-in-out">
          {GroupsName}
        </div>
        <div className="text-dark-grey-lighter text-sm">
          {SlicedGroupDescription}
        </div>
        <div className="text-dark-grey">Members: {GroupMembers}</div>
      </div>
      {withButton && <Button className="whitespace-nowrap">Join Group</Button>}
    </Link>
  );
}

export default GroupCard;
