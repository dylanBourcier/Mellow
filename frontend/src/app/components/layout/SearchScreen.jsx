"use client"

import React from 'react'
import Label from '../ui/Label'
import Input from '../ui/Input'
import { useState, useEffect } from 'react';
import UserCard from '../ui/UserCard';
import { toast } from 'react-hot-toast';
import CustomToast from '../ui/CustomToast';

export default function SearchScreen() {
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
          )}&groupId=${""}&excludeGroupMembers=false`
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


  return (
          <div>
          <Label htmlFor="search" className="mb-2">
            Search :
          </Label>
          <Input
            id="search"
            name="search"
            placeholder="Search for users..."
            value={query}
            onChange={(e) => setQuery(e.target.value)}
          ></Input>
          {loading && <p>Loading...</p>}
          {!loading && query.length >= 2 && results&&results.length === 0 && (
            <p>No results found</p>
          )}
          <ul className='mt-4 flex flex-col gap-2'>
            {results &&results.length > 0 && results.map((user) => (
              <UserCard key={user.user_id} user={user} />
            ))}
          </ul>
        </div>

  )
}
