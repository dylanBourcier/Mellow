'use client';

import React from 'react';
import UserInfo from '../ui/UserInfo';
import Button from '../ui/Button';
import { useState } from 'react';

// EventID      uuid.UUID `json:"event_id"`
// UserID       uuid.UUID `json:"user_id"`
// GroupID      uuid.UUID `json:"group_id"`
// CreationDate time.Time `json:"creation_date"`
// EventDate    time.Time `json:"event_date"`
// Title        string    `json:"title"`

function EventScreenPage({ event }) {
  const [vote, setVote] = useState(null);

  //     const { event_id, user_id, group_id, creation_date, event_date, title,username
  // , avatar_url
  //      } = event;
  const getButtonStyle = (type) =>
    `flex flex-1 hover:cursor-pointer px-4 py-1 rounded-xl font-light border transition-all duration-200 ${
      vote === type
        ? 'bg-lavender-1 text-white border-lavender-1'
        : 'text-lavender-5 border-lavender-5'
    }`;

  const handleVoteSubmit = () => {
    if (!vote) {
      alert('Please select Yes or No before voting.');
      return;
    }

    // 👉 Remplace ceci par ton appel API ensuite
    console.log(`User voted: ${vote}`);

    // Optionnel : désactiver les boutons ou donner un retour visuel
  };

  const creation_date = '25-10-2023 12:00:00'; // Example date, replace with actual data
  const username = 'John Doe'; // Example username, replace with actual data
  const avatar_url = '/img/DefaultAvatar.png'; // Example avatar URL, replace with actual data
  const event_date = '26/10/2023'; // Example event date, replace with actual data
  const title =
    'Sample Event Title That Spans Across One and a Half Lines and Provides More Context About the Event for Better Understanding ?'; // Example title, replace with actual data
  return (
    <div className="flex flex-col w-full bg-white shadow-(--box-shadow) rounded-lg hover:-translate-y-0.5 transition-all duration-300">
      <div className="flex  w-full p-3 gap-3 items-center justify-between">
        <UserInfo userName={username} authorAvatar={avatar_url} />
        <span>{creation_date}</span>
      </div>
      <div className="flex w-full px-3 gap-3 items-start justify-between">
        <h4 className="text-dark-grey p-0 font-heading text-lg leading-tight">
          {title}{' '}
          <span className="text-dark-grey font-bold">on {event_date}</span>{' '}
        </h4>
      </div>
      <div className="flex w-full items-center gap-2 px-3 mt-1 mb-1.5">
        <button
          onClick={() => setVote('yes')}
          className={getButtonStyle('yes')}
        >
          Yes
        </button>
        <button onClick={() => setVote('no')} className={getButtonStyle('no')}>
          No
        </button>
      </div>
      <div className="flex w-full items-center gap-2 px-3 mb-3">
        <Button className="w-full" onClick={handleVoteSubmit} disabled={!vote}>
          Vote
        </Button>
      </div>
    </div>
  );
}

export default EventScreenPage;
