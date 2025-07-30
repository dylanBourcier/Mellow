'use client';

import React, { use, useEffect } from 'react';
import PageTitle from '../ui/PageTitle';
import { useState } from 'react';
import { useUser } from '@/app/context/UserContext';
import Spinner from '../ui/Spinner';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';
import UserCard from '../ui/UserCard';

export default function UsersList({ userId, type }) {
  const [userList, setUserList] = useState([]);
  const [userInfo, setUserInfo] = useState(null); // State for user info
  const { user } = useUser();
  const [loading, setLoading] = useState(true);

  // Fetch user data based on userId and type (followers or following)
  useEffect(() => {
    async function fetchUsersListAndUserInfo() {
      try {
        const [usersResponse, userInfoResponse] = await Promise.all([
          fetch(`/api/users/${type}/${userId}`),
          fetch(`/api/users/${userId}`), // Fetch user info
        ]);

        if (!usersResponse.ok || !userInfoResponse.ok) {
          throw new Error('Failed to fetch data');
        }

        const usersData = await usersResponse.json();
        if (usersData.status !== 'success') {
          throw new Error(usersData.message || 'Failed to fetch users list');
        }
        const userInfoData = await userInfoResponse.json();
        if (userInfoData.status !== 'success') {
          throw new Error(userInfoData.message || 'Failed to fetch user info');
        }

        setUserList(usersData.data);
        setUserInfo(userInfoData.data); // Set user info

        setLoading(false);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }
    fetchUsersListAndUserInfo();
  }, [userId, type]);

  if (loading) {
    return (
      <div>
        <Spinner></Spinner>Loading ...
      </div>
    );
  }

  return (
    <div>
      <Link
        href={`/user/${userId}`}
        className="group flex items-center hover:underline hover:text-lavender-3 w-fit text-sm"
      >
        <span className="group-hover:animate-bounce">
          {icons['back_arrow']}
        </span>
        Back to {userInfo?.username || 'Unknown User'}
        {userInfo?.username?.endsWith('s') ? "'" : "'s"} profile
      </Link>
      <PageTitle>
        {`${userInfo?.username || 'Unknown User'}${
          userInfo?.username?.endsWith('s') ? "'" : "'s"
        } ${type}`}{' '}
        list
      </PageTitle>
      {userList && userList.length > 0 ? (
        <div className="flex flex-col gap-2.5">
          {userList.map((user) => (
            <UserCard key={user.user_id} user={user} />
          ))}
        </div>
      ) : (
        <span>
          {userInfo.username} has no {type} yet...
        </span>
      )}
    </div>
  );
}
