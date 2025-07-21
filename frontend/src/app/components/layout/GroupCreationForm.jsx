'use client';

import React from 'react';
import Input from '@/app/components/ui/Input';
import Label from '@/app/components/ui/Label';
import Button from '@/app/components/ui/Button';
import { useForm } from 'react-hook-form';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import toast from 'react-hot-toast';
import CustomToast from '@/app/components/ui/CustomToast';

export default function GroupCreationForm() {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const onSubmit = async (data) => {
    try {
      const formData = new FormData();

      // Champs texte
      formData.append('title', data.title);
      formData.append('description', data.description || '');
      

      const res = await fetch('/api/groups', {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });

      const result = await res.json();

      if (result.status === 'error') {
        throw new Error(result.message || 'Group creation failed');
      }

      toast.custom((t) => (
        <CustomToast
          message="Group created successfully!"
          t={t}
          type="success"
        />
      ));

      router.push('/groups');
    } catch (err) {
      toast.custom((t) => (
        <CustomToast
          message={err.message || 'An error occurred while creating the group.'}
          t={t}
          type="error"
        />
      ));
    }
  };

  return (
    <div className="flex flex-col items-center ">
      <form
        className="flex flex-col gap-2.5 max-w-[600px] w-full"
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="flex flex-col lg:flex-row w-full gap-2">
          <div className="flex-1">
            <Label htmlFor="first_name" className="block mb-2">
              Title* :
            </Label>
            <Input
              type="text"
              id="title"
              placeholder="Enter group title..."
              className={errors.title ? 'border border-error' : ''}
              {...register('title', { required: true })}
            />
            {errors.title && <span className="text-error">Required</span>}
          </div>
        </div>
        <div>
          <Label htmlFor="description">Descritpion* :</Label>
          <Input
            type="textarea"
            placeholder="Enter group description..."
            id="description"
            {...register('description')}
            className="h-24"
          />
        </div>
        <div className='flex flex-col gap-2 w-full'> 
          <Button type="submit">Create group</Button>
          <Link href={'/groups'}></Link>
        </div>
      </form>
    </div>
  );
}
