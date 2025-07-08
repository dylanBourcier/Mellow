import React from 'react';
import Image from 'next/image';
import PostCard from '../ui/PostCard';
import Link from 'next/link';


function ProfileScreen({ firstName, lastName, username, email, birthdate, followers, following,authorAvatar, description, myposts, userId}) {
    const postId = 1;// Example user ID
    const postTitle = 'Sample Post Title';
    const postContent =
      'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod.';
    const authorAvatar2 = '/img/lion.png'; // Example avatar image
    const userName = 'johndoe';
    const date = '2023-10-01';
    const Comments = 5;
    const props = {
      postId,
      postTitle,
      postContent,
      authorAvatar,
      userName,
      date,
      Comments,
    };
    return (
    <div className='flex flex-col gap-2.5 w-full'>
        <div className='flex flex-col items-center gap-2.5'>
            <Image
                src={authorAvatar}
                width={64}
                height={64}
                alt="Author Avatar"
            ></Image>
            <div className='flex items-center gap-2.5'>
                <div className='font-quickSand text-2xl font-medium'>
                    {firstName} {lastName}
                </div>
                <div className='font-quickSand text-xl'>
                    ({username})
                </div>
            </div>
                <div>
                    {description}
                </div>
        </div>
            <div className='flex flex-col gap-2.5 items-start py-4 font-inter text-sm'>
                <div>
                    Email: <span className='text-dark-grey-lighter font-inter'>{email}</span>
                </div>
                <div>
                    Birthdate: <span className='text-dark-grey-lighter font-inter'>{birthdate}</span>
                </div>
                <div className='flex gap-2.5 underline cursor-pointer font-inter text-sm'>
                    <div>
                        <Link href={`../user/${userId}/followers`}>Followers: <span >{followers}</span></Link>
                    </div>
                    <div>
                        <Link href={`../user/${userId}/following`}>Following: <span>{following}</span></Link>
                    </div>
                </div>
            </div>
        <PostCard postInfos={props} />
        <PostCard postInfos={props} />
        <PostCard postInfos={props} />
        <PostCard postInfos={props} />
    </div>
    );
}

export default ProfileScreen;