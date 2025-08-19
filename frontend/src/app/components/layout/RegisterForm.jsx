'use client';

import React from 'react';
import Input from '@/app/components/ui/Input';
import Label from '@/app/components/ui/Label';
import Button from '@/app/components/ui/Button';
import { useForm } from 'react-hook-form';
import { useRouter } from 'next/navigation';
import FileInput from '@/app/components/ui/FileInput';
import Link from 'next/link';
import toast from 'react-hot-toast';
import CustomToast from '@/app/components/ui/CustomToast';
import { Controller } from 'react-hook-form';

export default function RegisterForm() {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    control,
    setValue,
    watch,
    formState: { errors },
  } = useForm(
    {
      defaultValues: {
        first_name: '',
        last_name: '',
        username: '',
        email: '',
        birthdate: '',
        password: '',
        confirm_password: '',
        avatar: null,
        about: '',
        privacy: 'public', // Default to public
      },
    }
  );

  // Watch the privacy field to conditionally render messages
  const privacy = watch('privacy'); // Default to 'public'

  const onSubmit = async (data) => {
    try {
      const formData = new FormData();
      
      
      // Champs texte
      formData.append('firstname', data.first_name);
      formData.append('lastname', data.last_name);
      formData.append('username', data.username);
      formData.append('email', data.email);
      formData.append('birthdate', data.birthdate);
      formData.append('password', data.password);
      formData.append('description', data.about || '');
      formData.append('privacy',data.privacy); // Default privacy setting
      
      // FileInput (avatar)
      if (data.avatar) {
        formData.append('avatar', data.avatar);
      }

      const res = await fetch('/api/auth/signup', {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });

      const result = await res.json();

      if (result.status === 'error') {
        if (result.message.includes('USER_ALREADY_EXISTS')) {
          throw new Error('User already exists with this email or username');
        }
        throw new Error(result.message || 'Registration failed');
      }

      toast.custom((t) => (
        <CustomToast
          message="Welcome to Mellow! 🎉 Redirecting to login..."
          type="success"
        />
      ));

      router.push('/login');
    } catch (err) {
      toast.custom((t) => (
        <CustomToast
          message={err.message || 'Oops! Something went wrong'}
          type="error"
        />
      ));
    }
  };

  return (
    <>
      <form
        className="flex flex-col gap-2.5 max-w-[600px] w-full"
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="flex flex-col lg:flex-row w-full gap-2">
          <div className="flex-1">
            <Label htmlFor="first_name" className="block mb-2">
              First Name* :
            </Label>
            <Input
              type="text"
              id="first_name"
              className={errors.first_name ? 'border border-error' : ''}
              {...register('first_name', { required: true })}
            />
            {errors.first_name && <span className="text-error">Required</span>}
          </div>
          <div className="flex-1">
            <Label htmlFor="last_name" className="block mb-2">
              Last Name* :
            </Label>
            <Input
              type="text"
              id="last_name"
              className={errors.last_name ? 'border border-error' : ''}
              {...register('last_name', { required: true })}
            />
            {errors.last_name && <span className="text-error">Required</span>}
          </div>
        </div>
        <div>
          <Label htmlFor="username">Username* :</Label>
          <Input
            type="text"
            id="username"
            className={errors.username ? 'border border-error' : ''}
            {...register('username', { required: true })}
          />
          {errors.username && <span className="text-error">Required</span>}
        </div>
        <div>
          <Label htmlFor="email" className="block mb-2">
            Email* :
          </Label>
          <Input
            type="email"
            id="email"
            className={errors.email ? 'border border-error' : ''}
            {...register('email', { required: true })}
          />
          {errors.email && <span className="text-error">Required</span>}
        </div>
        <div>
          <Label htmlFor="birthdate" className="block mb-2">
            Birthdate* :
          </Label>
          <Input
            type="date"
            id="birthdate"
            className={errors.birthdate ? 'border border-error' : ''}
            {...register('birthdate', { required: true })}
          />
          {errors.birthdate && <span className="text-error">Required</span>}
        </div>
        <div className="flex flex-col lg:flex-row w-full gap-2">
          <div className="flex-1">
            <Label htmlFor="password" className="block mb-2">
              Password* :
            </Label>
            <Input
              type="password"
              id="password"
              className={errors.password ? 'border border-error' : ''}
              {...register('password', { required: true })}
            />
            {errors.password && <span className="text-error">Required</span>}
          </div>
          <div className="flex-1">
            <Label htmlFor="confirm_password" className="block mb-2">
              Confirm Password* :
            </Label>
            <Input
              type="password"
              id="confirm_password"
              className={
                errors.confirm_password ||
                watch('password') !== watch('confirm_password')
                  ? 'border border-error'
                  : ''
              }
              {...register('confirm_password', {
                required: true,
                validate: (val) =>
                  val === watch('password') || "Passwords don't match",
              })}
            />
            {errors.confirm_password && (
              <span className="text-error">
                {errors.confirm_password.message || 'Required'}
              </span>
            )}
          </div>
        </div>
        <div>
          <Label htmlFor="avatar">Avatar :</Label>
          <Controller
            name="avatar"
            control={control}
            render={({ field }) => (
              <FileInput
                id="avatar"
                name="avatar"
                setValue={setValue}
                onChange={(file) => field.onChange(file)}
              />
            )}
          />
        </div>
        <>
          <Label htmlFor={'privacy'}>Privacy* :</Label>
          <div>
            <div className="flex gap-1 justify-center">
              <div className="flex-1">
                <input
                  type="radio"
                  name="privacy"
                  id="public"
                  className="hidden peer"
                  value="public"
                  {...register('privacy', { required: true })}
                  defaultChecked
                />
                <label
                  htmlFor="public"
                  className="flex align-middle justify-center py-2 border border-lavender-5 text-lavender-5 bg-transparent rounded-2xl peer-checked:bg-lavender-6 cursor-pointer transition-colors duration-200"
                >
                  Public
                </label>
              </div>
              <div className="flex-1">
                <input
                  type="radio"
                  name="privacy"
                  id="private"
                  className="hidden peer"
                  value="private"
                  {...register('privacy', { required: true })}
                />
                <label
                  htmlFor="private"
                  className="flex align-middle justify-center py-2 border border-lavender-5 text-lavender-5 bg-transparent rounded-2xl peer-checked:bg-lavender-6 cursor-pointer transition-colors duration-200"
                >
                  Private
                </label>
              </div>
            </div>
          </div>
          {privacy === 'public' && (
            <span className="italic text-sm">
              Your profil will be visible to everyone.
            </span>
          )}
          {privacy === 'private' && (
            <span className="italic text-sm">
              Your profil will be visible by your followers only.
            </span>
          )}
        </>
        <div>
          <Label htmlFor="about">About me :</Label>
          <Input
            type="textarea"
            id="about"
            {...register('about')}
            className="h-24"
          />
        </div>
        <Button type="submit">Register</Button>
        <p className="text-center text-sm">
          Already have a <span className="text-lavender-3">Mellow</span>{' '}
          account?{' '}
          <Link href="/login" className="underline hover:font-[500]">
            Log in
          </Link>
        </p>
      </form>
    </>
  );
}
