'use client';
import React from 'react';
import JoinedGroupsList from '@/app/components/layout/JoinedGroupsList';
import GroupsList from '@/app/components/layout/GroupsList';
import { useState } from 'react';

export default function GroupListContainer() {
  const [activeTab, setActiveTab] = useState('joined');

  return (
    <div className="h-full">
      {/* Tabs pour mobile */}
      <div className="flex lg:hidden justify-around mb-2 gap-1 ">
        <button
          className={`w-1/2 py-2 text-sm font-medium ${
            activeTab === 'joined'
              ? 'bg-lavender-6 rounded-2xl text-lavender-5'
              : 'text-grey'
          }`}
          onClick={() => setActiveTab('joined')}
        >
          Joined Groups
        </button>
        <button
          className={`w-1/2 py-2 text-sm font-medium ${
            activeTab === 'all'
              ? 'bg-lavender-6 rounded-2xl text-lavender-5'
              : 'text-grey'
          }`}
          onClick={() => setActiveTab('all')}
        >
          All Groups
        </button>
      </div>

      {/* Contenu : les deux composants visibles sur desktop */}
      <div className="flex flex-col lg:flex-row h-full gap-2">
        <div
          className={`lg:block ${
            activeTab === 'joined' ? 'block' : 'hidden'
          } lg:w-1/2`}
        >
          <JoinedGroupsList />
        </div>
        <div
          className={`lg:block ${
            activeTab === 'all' ? 'block' : 'hidden'
          } lg:w-1/2`}
        >
          <GroupsList />
        </div>
      </div>
    </div>
  );
}
