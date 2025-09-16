import React from 'react';
import Link from 'next/link';
import Button from '../ui/Button';
import { icons } from '@/app/lib/icons'; // Assurez-vous que le chemin est correct

function GroupCard({ props, withButton = false, currentUserId }) {
  // const GroupMembers = 10; // Placeholder for number of members in the group
  if (!props) return null;
  const { group_id, title, description, user_id, member_count } = props;
  const SlicedGroupDescription =
    description.length > 100 ? description.slice(0, 100) + '...' : description;

  const isOwner = currentUserId === user_id;
  return (
    <Link
      href={'/groups/' + group_id}
      className="flex hover:bg-light-grey items-center rounded-lg w-full transition-all duration-100 ease-in-out justify-between"
    >
      <div className="flex flex-col p-2 group w-full">
        <div className="text-dark-grey font-heading text-xl group-hover:text-lavender-5 transition-all duration-100 ease-in-out">
          {title}
        </div>
        <div className="text-dark-grey-lighter text-sm">
          {SlicedGroupDescription}
        </div>
        <div className="text-dark-grey flex items-center gap-1 text-sm">
          {icons['groups16']}
          {member_count}
          {isOwner && (
            <>
              {' '}
              - <span className="text-lavender-5">Owner</span>
            </>
          )}
        </div>
      </div>
      {withButton && <Button className="whitespace-nowrap">Join Group</Button>}
    </Link>
  );
}

export default GroupCard;
