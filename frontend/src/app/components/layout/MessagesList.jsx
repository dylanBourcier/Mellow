import React from 'react';
import Message from '../ui/Message';

export default function MessagesList({ messages, type }) {
  switch (type) {
    case 'private':
      return (
        <div className="flex flex-col gap-2.5">
          {messages.map((message) => {
            return <Message message={message} key={message.message_id} />;
          })}
        </div>
      );
    case 'group':
      return <ul>{/* TODO */}</ul>;
    default:
      return <p>No messages available</p>;
  }
}
