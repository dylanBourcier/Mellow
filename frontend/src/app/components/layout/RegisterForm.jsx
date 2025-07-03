'use client';

import React from 'react';
import Input from '@/app/components/ui/Input';
import Label from '@/app/components/ui/Label';
import Button from '@/app/components/ui/Button';

export default function RegisterForm() {
  return (
    <form className="flex flex-col gap-4 max-w-[600px] w-full">
      <div>
        <Label htmlFor="username">Username* :</Label>
        <Input type="text" id="username" name="username" required />
      </div>
      <div>
        <Label htmlFor="email" className="block mb-2">
          Email* :
        </Label>
        <Input
          type="email"
          id="email"
          name="email"
          required
          className="w-full p-2 border border-light-grey rounded"
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
            Password* :
          </Label>
          <Input
            type="password"
            id="confirm_password"
            name="confirm_password"
            required
          />
        </div>
      </div>
      <Button type="submit">Register</Button>
    </form>
  );
}
