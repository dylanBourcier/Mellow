'use client';

import React from 'react';
import Input from '@/app/components/ui/Input';
import Label from '@/app/components/ui/Label';
import Button from '@/app/components/ui/Button';

import Link from 'next/link';
import { useState } from 'react';
import { useUser } from '@/app/context/UserContext';
import { useRouter } from 'next/navigation';

export default function LoginForm() {
  const [identifier, setIdentifier] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);

  const { setUser } = useUser();

  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ identifier, password }),
        credentials: 'include',
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }

      // Récupérer l'utilisateur directement après login
      const meRes = await fetch('/api/me', { credentials: 'include' });
      const meData = await meRes.json();

      // Mettre à jour le contexte
      router.push('/'); // Redirect to home page on successful login
      setUser(meData.data);
      // Handle successful login (e.g., redirect or show success message)
    } catch (error) {
      setError(error.message);
    }
  };

  return (
    <form
      className="flex flex-col gap-2.5 max-w-[600px] w-full "
      onSubmit={handleSubmit}
    >
      <div>
        <Label htmlFor="username">Username or email* :</Label>
        <Input
          type="text"
          id="username"
          name="username"
          placeholder="Enter your username..."
          value={identifier}
          onChange={(e) => setIdentifier(e.target.value)}
          required
        />
      </div>
      <div>
        <Label htmlFor="password" className="block mb-2">
          Password* :
        </Label>
        <Input
          type="password"
          id="password"
          placeholder="Enter your password..."
          name="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
      </div>
      <Button type="submit">Log in</Button>
      {error && <p className="text-red-500 text-sm">{error}</p>}
      <p className="text-center text-sm">
        You still don't have a <span className="text-lavender-3">Mellow</span>{' '}
        account ?{' '}
        <Link href="/register" className="underline hover:font-[500] ">
          Register
        </Link>
      </p>
    </form>
  );
}
