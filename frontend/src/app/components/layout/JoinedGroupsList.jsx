'use client';

import React, { useEffect } from 'react';
import GroupCard from '../ui/GroupCard';
import Button from '../ui/Button';
import { useState } from 'react';

function JoinedGroupsList(props) {
  const [groups, setGroups] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await fetch('/api/groups/joined');
        const data = await response.json();
        if (data.status !== 'success') {
          console.log('Error fetching groups:', data.message);
          
          throw new Error('Failed to fetch groups');
        }
        setGroups(data.data);
        
      } catch (error) {
        console.error('Error fetching groups:', error);
      } finally {
        setLoading(false);
      }
    };
    fetchGroups();
  }
  , []);


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
      {loading ? (
        <div className="text-center text-gray-500">Loading...</div>
      ) : groups.length==0 ? (
        <div className="text-center text-gray-500">
          You haven't joined any groups yet.
        </div>
      ) : (
        groups.map((group) => (
          <GroupCard key={group.group_id} props={group} currentUserId={group.user_id} />
        ))
      )}
    </div>
  );
}

export default JoinedGroupsList;
