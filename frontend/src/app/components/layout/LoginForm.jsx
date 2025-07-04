'use client';

import React from 'react';
import Input from '@/app/components/ui/Input';
import Label from '@/app/components/ui/Label';
import Button from '@/app/components/ui/Button';
import { useForm } from 'react-hook-form';
import FileInput from '@/app/components/ui/FileInput';
import Link from 'next/link';

export default function RegisterForm() {
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm();

  const onSubmit = (data) => {
    console.log(data);
    // Envoyer à ton backend via fetch + FormData
  };

  return (
    <form className="flex flex-col gap-2.5 max-w-[600px] w-full">
      <div>
        <Label htmlFor="username">Username or email* :</Label>
        <Input type="text" id="username" name="username" required />
      </div>
      <div>
        <Label htmlFor="password" className="block mb-2">
          Password* :
        </Label>
        <Input type="password" id="password" name="password" required />
      </div>
      <Button type="submit">Log in</Button>
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
