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
        //We need to check if the user try to access a conversation with himself if its the case redirect him on where he was before
        if (user.user_id === id) {
          window.history.back();
          return;
        }
        // Fetch user information
        const userResponse = await fetch(`/api/users/${id}`);

        const userInfo = await userResponse.json();
        if (userInfo.status !== 'success') {
          throw new Error(userInfo.message);
        }

        // Fetch messages between logged-in user and the other user
        const messagesResponse = await fetch(
          `/api/messages/${id}?limit=50&offset=0`
        );
        const messagesData = await messagesResponse.json();

        if (messagesData.status !== 'success') {
          throw new Error(messagesData.message);
        }

        setUserData(userInfo.data);

        if (messagesData.data) {
          setMessages(messagesData.data);
        } else {
          setMessages([]);
        }
        // Scroll to bottom after loading messages
        setTimeout(() => {
          const chatContainer = document.querySelector('.overflow-y-auto');
          if (chatContainer) {
            chatContainer.scrollTop = chatContainer.scrollHeight;
          }
        }, 100);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, [id]);
  useEffect(() => {
    if (!id || !user) return;

    const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws';
    // user.user_id = ton ID connecté
    const room =
      user.user_id < id ? `${user.user_id}:${id}` : `${id}:${user.user_id}`;
    const ws = new WebSocket(
      `ws://localhost:3225/ws/chat?room=private:${room}`
    );

    ws.onopen = () => console.log('WS connected');

    ws.onmessage = (event) => {
      const newMsg = JSON.parse(event.data);
      console.log('Received WS message:', newMsg);
      setMessages((prev) => [...prev, newMsg]);
      // Scroll to bottom when a new message arrives
      setTimeout(() => {
        const chatContainer = document.querySelector('.overflow-y-auto');
        if (chatContainer) {
          chatContainer.scrollTop = chatContainer.scrollHeight;
        }
      }, 100);
    };

    ws.onclose = (event) => {
      console.log('WS closed', event.code, event.reason);
    };

    ws.onerror = (event) => {
      console.error('WS error:', event);
    };

    return () => ws.close();
  }, [id, user]);

  return (
    <div className="flex flex-col h-full max-h-[90vh]">
      {!userData ? (
        <div className="flex-1 flex justify-center items-center">
          <Spinner />
        </div>
      ) : (
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
      )}

      <section className="flex-1 overflow-y-auto py-2">
        {messages.length === 0 ? (
          <p>No messages yet. Start the conversation!</p>
        ) : (
          <MessagesList messages={messages} type="private" />
        )}
      </section>
      <div className="p-4 bg-white rounded-b-2xl shadow-(--box-shadow)">
        <form
          onSubmit={async (e) => {
            e.preventDefault();
            const form = e.target;
            const formData = new FormData(form);
            const content = formData.get('message');
            if (!content) return;

            try {
              const response = await fetch(`/api/messages/${id}`, {
                method: 'POST',
                headers: {
                  'Content-Type': 'application/json',
                },
                body: JSON.stringify({ content }),
              });
              const result = await response.json();
              if (result.status !== 'success') {
                throw new Error(result.message);
              }
              form.reset();
              // Message will be added here

              // Scroll to bottom after sending message
              setTimeout(() => {
                const chatContainer =
                  document.querySelector('.overflow-y-auto');
                if (chatContainer) {
                  chatContainer.scrollTop = chatContainer.scrollHeight;
                }
              }, 100);
            } catch (error) {
              console.error('Error sending message:', error);
            }
          }}
          className="flex items-center gap-2"
        >
          <input
            type="text"
            name="message"
            placeholder="Type your message..."
            className="flex-1 border border-gray-300 rounded-full px-4 py-2 focus:outline-none focus:ring-2 focus:ring-lavender-5"
          />
          <button
            type="submit"
            className="bg-lavender-5 text-white rounded-full px-4 py-2 hover:bg-lavender-4 focus:outline-none focus:ring-2 focus:ring-lavender-5"
          >
            Send
          </button>
        </form>
      </div>
    </div>
  );
}
