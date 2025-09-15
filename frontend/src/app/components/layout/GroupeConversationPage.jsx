'use client';

import React, { useEffect, useState } from 'react';
import { useUser } from '@/app/context/UserContext';
import Spinner from '../ui/Spinner';
import MessagesList from './MessagesList';

export default function GroupConversationPage({ groupId }) {
  const { user } = useUser(); // Logged-in user

  const [messages, setMessages] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Fetch messages between logged-in user and the other user
        const messagesResponse = await fetch(
          `/api/messages/group/${groupId}?limit=50&offset=0`
        );
        const messagesData = await messagesResponse.json();

        if (messagesData.status !== 'success') {
          throw new Error(messagesData.message);
        }

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
  }, [groupId]);
  useEffect(() => {
    if (!groupId || !user) return;

    const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws';
    // user.user_id = ton ID connecté
    const room = groupId; // ID de la conversation (ex: user1_user2)
    const ws = new WebSocket(
      `${protocol}://localhost:3225/ws/chat?room=group:${room}`
    );

    ws.onopen = () => console.log('WS connected');

    ws.onmessage = (event) => {
      const newMsg = JSON.parse(event.data);
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
  }, [groupId, user]);

  if (!messages) {
    return (
      <div>
        <Spinner></Spinner>Loading...
      </div>
    );
  }
  return (
    <div className="flex flex-col h-full max-h-[90vh]">
      <section className="flex-1 overflow-y-auto py-2">
        {messages.length === 0 ? (
          <p>No messages yet. Start the conversation!</p>
        ) : (
          <MessagesList messages={messages} type="group" />
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

              const response = await fetch(`/api/messages/group/${groupId}`, {
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
