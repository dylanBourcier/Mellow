'use client';
import React from 'react';
import LoginForm from '@/app/components/layout/LoginForm';
import PageTitle from '@/app/components/ui/PageTitle';
import Image from 'next/image';
import GuestOnly from '../auth/GuestOnly';

export default function LoginScreen() {
  return (
    <GuestOnly redirectTo="/">
      <div className="flex flex-col items-center justify-center">
        <PageTitle>Login</PageTitle>
        <h2 className="flex justify-center items-center gap-2">
          <span>Welcome back to </span>
          <Image
            src="/img/Logo&Name.svg"
            width={120}
            height={48}
            alt="Mellow logo"
          ></Image>
        </h2>
        <LoginForm />
      </div>
    </GuestOnly>
  );
}
