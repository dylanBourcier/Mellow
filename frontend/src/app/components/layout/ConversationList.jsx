'use client'

import React, { use, useEffect } from 'react'
import RecentMessage from '../ui/RecentMessage';


export default function ConversationList() {

const [conversations, setConversations] = React.useState([]);

useEffect(() => {
    const fetchConversations = async () => {
        try {
            const response = await fetch('/api/messages');
            const data = await response.json();
            console.log(data);
            if (data.status !== 'success') {
                throw new Error(data.errorCode);
            }
            setConversations(data.data);
            console.log(data.data);

        } catch (error) {
            console.error('Error fetching conversations:', error);
        }
    };
    fetchConversations();
}, []);

return (
    <div className='flex flex-col'>
        {conversations.length > 0 ? (
            conversations.map((conv, idx) => (
                <RecentMessage key={idx} conversation={conv} />
            ))
        ) : (
            <p>No conversation for now</p>
        )}
    </div>
);
}
