import Link from 'next/link';
import React from 'react';
import Image from 'next/image';

function PostCard() {
    const  postId=1;
    const postTitle="Sample Post Title";
    const postContent="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod.";
    const firstname="John";
    const lastname="Doe";
    const authorAvatar="/img/DefaultAvatar.png";
    const userName="johndoe";
    const date="2023-10-01";
    const Comments=5;
    return (
        <div>
            <Link href="/post/1" className="flex flex-col p-6 gap-3 bg-white shadow-(--box-shadow) rounded-lg hover:-translate-y-0.5 transition-all duration-300">
            <div className='flex items-center gap-1'>
            <Image src={authorAvatar}
            alt="Author Avatar"
            width={32}
            height={32}>
            </Image>
            <span className='underline'>{firstname} {lastname} · {userName}</span>
            </div>
            <div>
            <h3>{postTitle}</h3>
            <span>{postContent}</span>
            </div>
            <div className='flex items-center justify-between text-sm'>
                <span className='flex gap-0.5 '>{Comments}<Image src="/img/comments.svg"
                alt="Comments Icon"
                width={16}
                height={16}>

                </Image></span>
                <span>{date}</span>
            </div>
            
            </Link>
        </div>
    );
}

export default PostCard;