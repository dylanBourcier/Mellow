'use client';

import React from 'react';
import CustomToast from '../ui/CustomToast';
import { useUser } from '@/app/context/UserContext';
import { toast } from 'react-hot-toast';
import { useState } from 'react';
import PageTitle from '../ui/PageTitle';
import { useEffect } from 'react';
import Spinner from '../ui/Spinner';
import { useRouter } from 'next/navigation';
import Button from '../ui/Button';
import GroupNavbar from './GroupNavbar';
import { icons } from '@/app/lib/icons';

export default function GroupLayoutHeader({ groupId }) {
  // Get user data and group data based on groupId
  const { user } = useUser();
  const [loading, setLoading] = useState(true);
  const [groupData, setGroupData] = useState(null);
  const router = useRouter();

  useEffect(() => {
    const fetchGroupData = async () => {
      try {
        const response = await fetch(`/api/groups/${groupId}`);
        const data = await response.json();
        if (data.status === 'error') {
          throw new Error(data.message);
        }
        if (data.data == null) {
          throw new Error('Group not found');
        }
        setGroupData(data.data);
      } catch (error) {
        toast.custom((t) => (
          <CustomToast
            type="error"
            t={t}
            message={'Failed to fetch group data'}
          ></CustomToast>
        ));
      } finally {
        setLoading(false);
      }
    };

    fetchGroupData();
  }, [groupId]);

  if (loading) {
    return (
      <div className="flex items-center h-11 justify-center">
        <Spinner />
        Loading...
      </div>
    );
  }

  if (!groupData && !loading) {
    router.push('/groups'); // Redirect to groups page if no group data
    return null; // Prevent rendering if no group data
  }
  return (
    <div className="flex flex-col w-full gap-1 lg:gap-2 p-2">
      <div className="flex">
        <PageTitle className="">{groupData.title}</PageTitle>
      </div>
      <p>{groupData.description}</p>
      {user?.user_id === groupData.user_id && (
        <div className="flex gap-2 items-center justify-between w-full">
          <span className="">
            {groupData.member_count} member
            {groupData.member_count > 1 ? 's' : ''}
          </span>
          {user && user.user_id === groupData.user_id && (
            <div className="flex items-center gap-2  self-end ml-auto lg:absolute lg:right-0">
              <button className="p-1.5 rounded-xl bg-lavender-1 border border-lavender-1 text-white flex gap-2">
                {icons['add_person']}
                <span className="">Add People</span>
              </button>
              <button className="p-1.5 rounded-xl border border-dark-grey">
                {icons['edit']}
              </button>
              <button className="p-1.5 rounded-xl border text-red-400 bg-red-100 border-red-400">
                {icons['trash']}
              </button>
            </div>
          )}
        </div>
      )}
      <div className="w-[50%] h-[1px] bg-lavender-2 self-center mt-1 mb-1"></div>
      <GroupNavbar groupId={groupId} />
    </div>
  );
}
