'use client';

import React from 'react';
import { useForm } from 'react-hook-form';
import { useRouter } from 'next/navigation';
import toast from 'react-hot-toast';
import CustomToast from '@/app/components/ui/CustomToast';
import Label from '../ui/Label';
import Input from '../ui/Input';
import Button from '../ui/Button';
import Link from 'next/link';

export default function CreateEventForm({ groupId }) {
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
      console.log('Title:', data.title);

      formData.append('event_date', data.event_date);
      console.log('Form data:', formData);

      const res = await fetch(`/api/groups/events/${groupId}`, {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });

      const result = await res.json();

      if (result.status === 'error') {
        throw new Error(result.message || 'Event creation failed');
      }

      toast.custom((t) => (
        <CustomToast
          message="Event created successfully!"
          t={t}
          type="success"
        />
      ));

      router.push('/groups/' + groupId + '/events');
    } catch (err) {
      toast.custom((t) => (
        <CustomToast
          message={err.message || 'An error occurred while creating the event.'}
          t={t}
          type="error"
        />
      ));
    }
  };

  return (
    <div className="flex flex-col items-center ">
      <h2 className="text-2xl font-semibold mb-4">Create Event</h2>
      <form
        className="flex flex-col gap-2.5 max-w-[600px] w-full"
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="flex flex-col lg:flex-row w-full gap-2">
          <div className="flex-1">
            <Label htmlFor="title" className="block mb-2">
              Title* :
            </Label>
            <Input
              type="text"
              id="title"
              placeholder="Enter event title..."
              className={errors.title ? 'border border-error' : ''}
              {...register('title', { required: true })}
            />
            {errors.title && <span className="text-error">Required</span>}
          </div>
        </div>
        <div className="flex flex-col gap-2">
          <Label htmlFor="event_date">Event Date* :</Label>
          <Input
            type="datetime-local"
            id="event_date"
            {...register('event_date', { required: true })}
          />
          {errors.event_date && (
            <span className="text-error">Event date is required</span>
          )}
        </div>
        <div className="flex flex-col gap-2 w-full">
          <Button type="submit">Create event</Button>
          <Link href={'/groups'}>Back to Groups</Link>
        </div>
      </form>
    </div>
  );
}
