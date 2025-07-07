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
      <div className="flex flex-col lg:flex-row w-full gap-2">
        <div className="flex-1">
          <Label htmlFor="first_name" className="block mb-2">
            First Name* :
          </Label>
          <Input type="text" id="first_name" name="first_name" required />
        </div>
        <div className="flex-1">
          <Label htmlFor="last_name" className="block mb-2">
            Last Name* :
          </Label>
          <Input type="text" id="last_name" name="last_name" required />
        </div>
      </div>
      <div>
        <Label htmlFor="username">Username* :</Label>
        <Input type="text" id="username" name="username" required />
      </div>
      <div>
        <Label htmlFor="email" className="block mb-2">
          Email* :
        </Label>
        <Input type="email" id="email" name="email" required />
      </div>
      <div>
        <Label htmlFor="birthdate" className="block mb-2">
          Birthdate* :
        </Label>
        <Input
          type="date"
          id="birthdate"
          name="birthdate"
          placeholder="********"
          required
        />
      </div>
      <div className="flex flex-col lg:flex-row w-full gap-2">
        <div className="flex-1">
          <Label htmlFor="password" className="block mb-2">
            Password* :
          </Label>
          <Input type="password" id="password" name="password" required />
        </div>
        <div className="flex-1">
          <Label htmlFor="confirm_password" className="block mb-2">
            Confirm Password* :
          </Label>
          <Input
            type="password"
            id="confirm_password"
            name="confirm_password"
            required
          />
        </div>
      </div>
      <div>
        <Label htmlFor="avatar">Avatar :</Label>
        <FileInput
          name="avatar"
          id="avatar"
          label="Chose a profile picture"
          setValue={setValue}
          register={register}
        />
      </div>
      <div>
        <Label htmlFor="about">About me :</Label>
        <Input
          type="textarea"
          id="about"
          name="about"
          placeholder="Tell us about yourself..."
          className="h-24"
        />
      </div>
      <Button type="submit">Register</Button>
      <p className="text-center text-sm">
        Already have a <span className="text-lavender-3">Mellow</span> account ?{' '}
        <Link href="/login" className="underline hover:font-[500] ">
          Log in
        </Link>
      </p>
    </form>
  );
}
