import RegisterForm from '@/app/components/layout/RegisterForm';
import PageTitle from '@/app/components/ui/PageTitle';
import Image from 'next/image';
import React from 'react';

export const metadata = {
  title: 'Register',
  description: 'Register a new account on Mellow',
};

export default function RegisterPage() {
  return (
    <div className="flex flex-col items-center justify-center">
      <PageTitle>Register</PageTitle>
      <h2 className="flex justify-center items-center gap-2">
        <span>Welcome to</span>
        <Image
          src="/img/Logo&Name.svg"
          width={120}
          height={48}
          alt="Mellow logo"
        ></Image>
      </h2>
      <RegisterForm />
    </div>
  );
}
