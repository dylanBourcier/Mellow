import React from 'react';
import UserInfo from '@/app/components/ui/UserInfo';

function CommentCard(props) {
    const commentContent = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.";
    const commentId = 1;
    const commentCreated_at = "15h 2023-10-01";
    return (
    <div className='flex flex-col gap-1 bg-white p-4 shadow-(--box-shadow) rounded-lg'>
        <div className='flex items-center justify-between gap-3'>
            <UserInfo></UserInfo>
            <span className='font-thin text-sm'>{commentCreated_at}</span>
        </div>
            <pre className='whitespace-pre-wrap break-words font-sans font-inter'>{commentContent}</pre>
    </div>
    );
}

export default CommentCard;