'use client';

import { createContext, useContext, useEffect, useState } from 'react';

const UserContext = createContext();

export function UserProvider({ children }) {
  const [user, setUser] = useState(null); // utilisateur connecté
  const [loading, setLoading] = useState(true); // en cours de fetch

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch('/api/me', { credentials: 'include' });
        if (!res.ok) throw new Error('Not authenticated');
        const json = await res.json();
        setUser(json.data);
      } catch (err) {
        setUser(null); // pas connecté
      } finally {
        setLoading(false);
      }
    };

    fetchUser();
  }, []);

  return (
    <UserContext.Provider value={{ user, setUser, loading }}>
      {children}
    </UserContext.Provider>
  );
}

export function useUser() {
  return useContext(UserContext);
}
