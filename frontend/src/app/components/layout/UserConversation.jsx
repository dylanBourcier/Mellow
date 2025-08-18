'use client';

import React, { useEffect, useState } from 'react';
import { useUser } from '@/app/context/UserContext';
import Spinner from '../ui/Spinner';
import Image from 'next/image';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';
import MessagesList from '../layout/MessagesList'; // Assuming you have a MessagesList component

export default function UserConversation({ id }) {
  const { user } = useUser(); // Logged-in user
  const [userData, setUserData] = useState(null);

  const [messages, setMessages] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Fetch user information
        const userResponse = await fetch(`/api/users/${id}`);
        const userInfo = await userResponse.json();
        if (userInfo.status !== 'success') {
          throw new Error(userInfo.message);
        }

        // Fetch messages between logged-in user and the other user
        const messagesResponse = await fetch(`/api/messages/${id}?limit=50&offset=0`);
        const messagesData = await messagesResponse.json();
        console.log(messagesData);
        
        if (messagesData.status !== 'success') {
          throw new Error(messagesData.message);
        }

        setUserData(userInfo.data);
        if (messagesData.data){

          setMessages(messagesData.data);
        }else{
          setMessages([]);
        }

        //Do some fake messages for testing
        // setMessages([
        //   {
        //     message_id: 1,
        //     sender_id: user.user_id,
        //     receiver_id: id,
        //     content: 'Hello, how are you?',
        //     creation_date: new Date().toISOString(),
        //   },
        //   {
        //     message_id: 2,
        //     sender_id: id,
        //     receiver_id: user.user_id,
        //     content: 'I am fine, thank you! How about you?',
        //     creation_date: new Date().toISOString(),
        //   },
        //   {
        //     message_id: 3,
        //     sender_id: user.user_id,
        //     receiver_id: id,
        //     content: 'I am doing well, thanks for asking!',
        //     creation_date: new Date().toISOString(),
        //   },
        // ]);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, [id]);

  if (!userData) {
    return (
      <div>
        <Spinner></Spinner>Loading...
      </div>
    );
  }
  return (
    <div>
      <section className="flex justify-center items-center py-2 bg-white relative rounded-2xl shadow-(--box-shadow)">
        <Link
          href={'/messages'}
          className="group flex items-center hover:underline hover:text-lavender-3 text-sm absolute left-2"
        >
          {' '}
          <span className="group-hover:animate-bounce">
            {icons['back_arrow']}
          </span>{' '}
          <span className="hidden lg:block">Back to conversations</span>
        </Link>
        <Link
          className="flex group items-center gap-2"
          href={`/user/${userData.user_id}`}
        >
          <Image
            src={userData.image_url || '/img/DefaultAvatar.svg'}
            alt={userData.username || 'User Avatar'}
            width={40}
            height={40}
            className="rounded-full w-10 h-10 object-cover border border-transparent group-hover:border-lavender-5 transition-all duration-200"
          />
          <span className="group-hover:underline group-hover:text-lavender-5 transition-all duration-200">
            {userData.firstname} {userData.lastname}
          </span>
        </Link>
        <div></div>
      </section>
      {messages.length === 0 ? (
        <p>No messages yet. Start the conversation!</p>
      ) : (
        <MessagesList messages={messages} type="private" />
      )}
    </div>
  );
}
