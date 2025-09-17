'use client';

import React, { useState, useEffect } from 'react';
import Input from '../ui/Input';
import Label from '../ui/Label';
import Button from '../ui/Button';
import FileInput from '../ui/FileInput';
import Modal from '../ui/Modal';
import { set, useForm } from 'react-hook-form';
import { useUser } from '@/app/context/UserContext';
import PageTitle from '../ui/PageTitle';
import CustomToast from '../ui/CustomToast';
import { toast } from 'react-hot-toast';
import { useRouter } from 'next/navigation';

function EditProfile() {
  const { user, setUser } = useUser();
  const router = useRouter();

  const {
    register,
    setValue,
    handleSubmit,
    formState: { errors },
    watch,
  } = useForm();

  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    if (user) {
      console.log('user data:', user);

      setValue('firstname', user.firstname || '');
      setValue('lastname', user.lastname || '');
      setValue('username', user.username || '');
      setValue(
        'birthdate',
        user.birthdate
          ? new Date(user.birthdate).toISOString().split('T')[0]
          : ''
      );
      setValue('description', user.description || '');
      setValue('privacy', user.privacy || 'public');
    }
  }, [user, setValue]);

  const onSubmit = async (data) => {
    try {
      const formData = new FormData();
      formData.append('firstname', data.firstname);
      formData.append('lastname', data.lastname);
      formData.append('username', data.username);
      formData.append('birthdate', data.birthdate);
      formData.append('description', data.description || '');
      formData.append('privacy', data.privacy);
      if (data.avatar) {
        formData.append('avatar', data.avatar);
      }

      const response = await fetch(`/api/users/${user.user_id}`, {
        method: 'PUT',
        body: formData,
        credentials: 'include',
      });
      const result = await response.json();
      if (result.status === 'error') {
        throw new Error(result.message);
      } else {
        toast.custom((t) => (
          <CustomToast message="Profile updated successfully!" type="success" />
        ));
        setUser((prevUser) => ({
          ...prevUser,
          ...data,
        }));
        router.push('/profile');
      }
    } catch (error) {
      toast.custom((t) => (
        <CustomToast
          message={error.message || 'Something went wrong'}
          type="error"
        />
      ));
    }
  };

  return (
    <form
      className="flex flex-col gap-2.5 max-w-[600px] w-full"
      onSubmit={handleSubmit(onSubmit)}
      encType="multipart/form-data"
    >
      <div className="flex items-center justify-center">
        <PageTitle>Edit Profile</PageTitle>
      </div>
      <div className="flex flex-col lg:flex-row gap-2.5 w-full">
        <Label htmlFor={'firstname'}>
          First Name* :
          <Input
            id="firstname"
            placeholder="Enter your firstname..."
            {...register('firstname', { required: 'First name is required' })}
          />
          {errors.firstname && (
            <span className="text-red-500">{errors.firstname.message}</span>
          )}
        </Label>
        <Label htmlFor={'lastname'}>
          Last Name* :
          <Input
            id="lastname"
            placeholder="Enter your lastname..."
            {...register('lastname', { required: 'Last name is required' })}
          />
          {errors.lastname && (
            <span className="text-red-500">{errors.lastname.message}</span>
          )}
        </Label>
      </div>
      <div className="flex flex-col items-start gap-2.5 w-full">
        <Label htmlFor={'username'}>Username :</Label>
        <Input
          id="username"
          placeholder="Enter your username..."
          readOnly
          {...register('username', { required: 'Username is required' })}
        />
        {errors.username && (
          <span className="text-red-500">{errors.username.message}</span>
        )}
      </div>
      <div className="flex flex-col items-start gap-2.5 w-full">
        <Label htmlFor={'birthdate'}>Birthdate* :</Label>
        <Input
          type="date"
          id="birthdate"
          name="birthdate"
          required
          {...register('birthdate', { required: 'Birthdate is required' })}
        />
        {errors.birthdate && (
          <span className="text-red-500">{errors.birthdate.message}</span>
        )}
      </div>
      <div className="flex flex-col lg:flex-row w-full gap-2"></div>
      <div className="flex flex-col items-start gap-2.5 w-full">
        <Label htmlFor="description">About me :</Label>
        <Input
          type="textarea"
          id="description"
          name="description"
          placeholder="Tell us about yourself..."
          className="h-24"
          {...register('description')}
        />
      </div>
      <div className="flex flex-col items-start gap-2.5 w-full">
        <Label htmlFor="privacy">Privacy*:</Label>
      </div>
      <div className="flex gap-2.5 w-full">
        <div className="flex gap-1 justify-center w-full">
          <div className="flex-1">
            <input
              type="radio"
              name="privacy"
              id="public"
              className="hidden peer"
              value="public"
              {...register('privacy', { required: 'Privacy is required' })}
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
              {...register('privacy')}
            />
            <label
              htmlFor="private"
              className="flex align-middle justify-center py-2 border border-lavender-5 text-lavender-5 bg-transparent rounded-2xl peer-checked:bg-lavender-6 cursor-pointer transition-colors duration-200"
            >
              Private
            </label>
          </div>
        </div>
        {errors.privacy && (
          <span className="text-red-500">{errors.privacy.message}</span>
        )}
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
      <div className="flex flex-1 align-middle justify-center gap-2.5 w-full">
        <Button
          type="button"
          className="w-full"
          onClick={() => setIsModalOpen(true)}
        >
          Save Changes
        </Button>
      </div>
      <Modal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        title="Confirm"
        message="Are you sure you want to save the changes?"
        actions={[
          {
            label: 'Yes',
            onClick: () => {
              setIsModalOpen(false);
              document.querySelector('form').requestSubmit();
            },
            closeOnClick: true,
          },
          {
            label: 'No',
            onClick: null,
            closeOnClick: true,
          },
        ]}
      />
    </form>
  );
}

export default EditProfile;
