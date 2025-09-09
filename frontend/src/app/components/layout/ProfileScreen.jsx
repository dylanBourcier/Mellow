'use client';

import { use, useEffect } from 'react';
import Image from 'next/image';
import Link from 'next/link';
import Button from '../ui/Button';
import { icons } from '@/app/lib/icons';
import { useState } from 'react';
import toast from 'react-hot-toast';
import CustomToast from '../ui/CustomToast';
import Spinner from '../ui/Spinner';
import { formatDateShort } from '@/app/utils/date';
import { useUser } from '@/app/context/UserContext';
import { useRouter } from 'next/navigation';
import UserPostsContainer from './UserPostsContainer';
import FollowButton from '../ui/FollowButton';

function ProfileScreen({ userId }) {
  const router = useRouter();
  const { user } = useUser();
  if (user.user_id === userId) {
    router.push('/profile');
  }
  if (!userId) {
    userId = user.user_id; // Fallback to current user if no ID is provided
  }

  const [userData, setUserData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [isLimited, setIsLimited] = useState(false);
  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const res = await fetch(`/api/users/${userId}`, {
          credentials: 'include',
        });
        const data = await res.json();
        if (data.status !== 'success') {
          throw new Error(data.message || 'Failed to fetch user data');
        }

        setUserData(data.data);

        setIsLimited(data.message === 'Limited');
      } catch (err) {
        toast.custom((t) => (
          <CustomToast
            t={t}
            type="error"
            message={'Error fetching user data! ' + err.message}
          />
        ));
      } finally {
        setLoading(false);
      }
    };

    fetchUserData();
  }, [userId]);

  if (loading) {
    return (
      <div className="flex items-center justify-center h-full">
        <Spinner />
        <span>Loading profile...</span>
      </div>
    );
  }
  if (!loading && !userData) {
    return (
      <div className="flex items-center justify-center h-full">
        <span className="text-red-500">User not found</span>
      </div>
    );
  }
  if (isLimited) {
    return (
      <div className="flex flex-col gap-2.5 w-full">
        <div className="flex flex-col items-center gap-2.5 relative">
          {userData.user_id !== user.user_id && (
            <div className="absolute top-3 right-3 flex gap-2.5">
              <Button
                href={`/messages/${userId}`}
                childrenClassName="flex items-center gap-1"
              >
                <span>{icons['messages']}</span>{' '}
                <span className="hidden lg:block">Message</span>
              </Button>
              <button className="flex items-center gap-1 px-2 py-2 rounded-2xl bg-red-100 text-error hover:bg-red-200 border border-red-300 ">
                <span>{icons['report']}</span>{' '}
                <span className="hidden lg:block cursor-pointer">Report</span>
              </button>
            </div>
          )}

          <Image
            src={userData?.image_url || '/img/DefaultAvatar.svg'}
            width={128}
            height={128}
            className="w-16 h-16 rounded-full border border-lavender-5 shadow-(--box-shadow) hover:transform hover:scale-150 hover:rotate-360 transition-all duration-300"
            alt="Author Avatar"
          ></Image>
          <div className="flex items-center gap-2.5">
            <div className="font-quickSand text-2xl font-medium">
              {userData.username}
            </div>
          </div>
        </div>
        <div className="flex flex-col gap-2.5 items-start py-4 font-inter text-sm">
          <div className="flex gap-2.5 font-inter text-sm justify-center w-full">
            {userData.user_id !== user.user_id && (
              <FollowButton
                targetID={userData.user_id}
                followStatus={userData.follow_status}
                privacy={userData.privacy}
              />
            )}
          </div>
        </div>
        <div className="flex flex-col items-center justify-center text-dark-grey italic">
          {icons['lock']}
          <span>
            This profile is private. Follow to see his informations and posts.
          </span>
        </div>
      </div>
    );
  } else {
    return (
      <>
        <div className="flex flex-col gap-2.5 w-full  bg-white p-4 rounded-2xl shadow-(--box-shadow) mb-6">
          <div className="flex flex-col items-center gap-2 relative">
            {userData.user_id !== user.user_id && (
              <div className="absolute top-3 right-3 flex gap-2.5">
                <Button
                  href={`/messages/${userId}`}
                  childrenClassName="flex items-center gap-1"
                >
                  <span>{icons['messages']}</span>{' '}
                  <span className="hidden lg:block">Message</span>
                </Button>
                <button className="flex items-center gap-1 px-2 py-2 rounded-2xl bg-red-100 text-error hover:bg-red-200 border border-red-300 ">
                  <span>{icons['report']}</span>{' '}
                  <span className="hidden lg:block cursor-pointer">Report</span>
                </button>
              </div>
            )}

            <Image
              src={userData?.image_url || '/img/DefaultAvatar.svg'}
              width={128}
              height={128}
              className="w-16 h-16 rounded-full border border-lavender-5 shadow-(--box-shadow) hover:transform hover:scale-150 hover:rotate-360 transition-all duration-300"
              alt="Author Avatar"
            ></Image>
            {userData.user_id === user.user_id && (
              <div className="whitespace-nowrap italic text-dark-grey-lighter font-light">
                Your profile is <strong>{userData.privacy}</strong>
              </div>
            )}

            <div className="flex items-center gap-2">
              <div className="font-quickSand text-2xl font-medium">
                {userData.username}
              </div>
              <div className="font-quickSand text-xl">
                ({userData.lastname} {userData.firstname})
              </div>
            </div>
            {userData?.description && (
              <div className="flex w-full text-sm">
                Description : {userData.description}
              </div>
            )}
          </div>
          <div className="flex flex-col gap-1 items-start  font-inter text-sm">
            <div>
              {userData.email} - {/* Calculate the age of the user */}
              {formatDateShort(userData.birthdate)} (
              {new Date().getFullYear() -
                new Date(userData.birthdate).getFullYear()}{' '}
              years old)
            </div>
            <div className="flex gap-2.5 font-inter text-sm justify-between w-full">
              <div className="flex gap-2.5 underline cursor-pointer font-inter text-sm items-center ">
                <div>
                  <Link
                    href={`/user/${userId}/followers`}
                    className="whitespace-nowrap"
                  >
                    <span>
                      {userData.followers_count} Follower
                      {userData.followers_count > 1 ? 's' : ''}
                    </span>
                  </Link>
                </div>
                <div>
                  <Link
                    href={`/user/${userId}/following`}
                    className="whitespace-nowrap"
                  >
                    <span>{userData.followed_count} Following</span>
                  </Link>
                </div>
              </div>
              {userData.user_id !== user.user_id && (
                <FollowButton
                  targetID={userData.user_id}
                  followStatus={userData.follow_status}
                  privacy={userData.privacy}
                  onFollowChange={(delta) => {
                    setUserData((prevData) => ({
                      ...prevData,
                      followers_count: prevData.followers_count + delta,
                    }));
                  }}
                />
              )}

              {userData.user_id === user.user_id && (
                <div className="flex items-center gap-2.5">
                  <Button
                    childrenClassName="flex items-center  gap-1"
                    href="/profile/edit"
                    wFull={false}
                    isSecondary={true}
                  >
                    <span>{icons['edit']}</span>
                    <span className="px-1.5 text-center hidden lg:inline">
                      Edit Profile
                    </span>
                  </Button>
                </div>
              )}
            </div>
          </div>
        </div>
        <UserPostsContainer userId={userId} />
      </>
    );
  }
}

export default ProfileScreen;
