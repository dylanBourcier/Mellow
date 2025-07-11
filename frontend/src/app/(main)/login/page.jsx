import LoginForm from '@/app/components/layout/LoginForm';
import LoginScreen from '@/app/components/layout/LoginScreen';
import PageTitle from '@/app/components/ui/PageTitle';
import Image from 'next/image';
import React from 'react';

export const metadata = {
  title: 'Login',
  description: 'Login to your Mellow account',
};

export default function LoginPage() {
  return <LoginScreen />;
}
