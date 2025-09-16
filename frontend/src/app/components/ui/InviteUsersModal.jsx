'use client';
import React, { useState, useEffect } from 'react';
import { toast } from 'react-hot-toast';
import CustomToast from '../ui/CustomToast';
import Spinner from '../ui/Spinner';
import Image from 'next/image';

export default function InviteUsersModal({ onClose, groupId }) {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (query.length < 2) return;

    const fetchResults = async () => {
      setLoading(true);
      try {
        const res = await fetch(
          `/api/users/search?q=${encodeURIComponent(
            query
          )}&groupId=${groupId}&excludeGroupMembers=true`
        );
        const data = await res.json();
        if (data.status === 'error') {
          throw new Error(data.message);
        }
        setResults(data.data); // suppose que l’API renvoie { users: [...] }
      } catch (err) {
        toast.custom((t) => (
          <CustomToast
            type="error"
            t={t}
            message={`Failed to fetch users + ${err}`}
          />
        ));
        setResults([]);
      } finally {
        setLoading(false);
      }
    };

    const timeout = setTimeout(fetchResults, 300); // debounce

    return () => clearTimeout(timeout);
  }, [query]);

  const handleInvite = async (userId) => {
    try {
      const res = await fetch(`/api/groups/invite/${groupId}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ user_id: userId }),
      });

      const data = await res.json();
      if (data.status === 'success') {
        toast.success('Invitation sent!');
      } else {
        toast.error(data.message || 'Failed to invite');
      }
    } catch (err) {
      toast.error('An error occurred');
    }
  };

  return (
    <div className="fixed inset-0 z-50 bg-black/50 flex justify-center items-center">
      <div className="bg-white p-6 rounded-xl w-full max-w-lg shadow-(--box-shadow)">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-lg font-bold">Invite users</h2>
          <button onClick={onClose} className="text-xl font-bold">
            ×
          </button>
        </div>
        <input
          className="w-full border border-gray-300 rounded p-2 mb-4"
          type="text"
          placeholder="Search users..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
        {loading ? (
          <Spinner />
        ) : results ? (
          <ul className="max-h-64 overflow-auto">
            {results.map((user) => (
              <li
                key={user.user_id}
                className="flex justify-between items-center p-2 hover:bg-gray-100"
              >
                <div className="flex items-center gap-2">
                  <Image
                    src={user.image_url || '/img/DefaultAvatar.svg'}
                    alt="User Avatar"
                    width={40}
                    height={40}
                    className="w-10 h-10 rounded-full"
                  />
                  <span>{user.username}</span>
                </div>
                <button
                  onClick={(e) => {
                    handleInvite(user.user_id);
                    e.target.textContent = 'Invited';
                    e.target.disabled = true;
                    e.target.classList.add('opacity-50', 'cursor-not-allowed');
                  }}
                  className="bg-lavender-3 px-2 py-1 rounded text-white"
                >
                  Invite
                </button>
              </li>
            ))}
          </ul>
        ) : (
          <p className="text-gray-500">No results found</p>
        )}
      </div>
    </div>
  );
}
