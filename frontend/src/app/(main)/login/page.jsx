import LoginForm from '@/app/components/layout/LoginForm';
import PageTitle from '@/app/components/ui/PageTitle';
import Image from 'next/image';
import React from 'react';

export const metadata = {
  title: 'Login',
  description: 'Login to your Mellow account',
};

export default function LoginPage() {
  return (
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
  );
}
