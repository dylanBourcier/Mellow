'use client';

import React, { useEffect, useState } from 'react';
import Input from '../ui/Input';
import Button from '../ui/Button';
import GroupCard from '../ui/GroupCard';
import CustomToast from '../ui/CustomToast';
import Spinner from '../ui/Spinner';

export default function GroupsList() {
  const [groups, setGroups] = useState([]);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await fetch('/api/groups/not-joined');
        const data = await response.json();
        if (data.status !== 'success') {
          throw new Error(data.message);
        }
        if (data.data === null) {
          setGroups([]); // Handle case where no groups are returned
          return;
        }
        setGroups(data.data);
      } catch (error) {
        toast.custom((t) => (
          <CustomToast
            message="Failed to fetch groups. Please try again later."
            t={t}
            type="error"
          />
        ));
      } finally {
        setLoading(false);
      }
    };
    fetchGroups();
  }, []);

  // Handle search input change
  const handleSearchChange = (event) => {
    setSearchQuery(event.target.value);
  };

  // Filter groups based on search query
  const filteredGroups = groups.filter((group) =>
    group.title.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div className="flex flex-col bg-white rounded-2xl shadow-(--box-shadow) p-4 h-full flex-1">
      <h2 className="text-lavender-5 text-shadow-(--text-shadow)">
        Join a new group
      </h2>
      <div className="flex items-center gap-2">
        <Input
          type="search"
          placeholder="Search for groups..."
          value={searchQuery}
          onChange={handleSearchChange}
        />
      </div>
      {loading ? (
        <div className="flex gap-2">
          <Spinner></Spinner>Loading...
        </div>
      ) : groups.length === 0 ? (
        <div className="text-dark-grey-lighter text-center">
          No groups available to join.
        </div>
      ) : filteredGroups.length === 0 ? (
        <div className="text-dark-grey-lighter text-center">
          No groups match your search.
        </div>
      ) : (
        <div className="flex flex-col gap-2.5">
          {filteredGroups.map((group) => (
            <GroupCard key={group.group_id} props={group} withButton />
          ))}
        </div>
      )}
    </div>
  );
}
