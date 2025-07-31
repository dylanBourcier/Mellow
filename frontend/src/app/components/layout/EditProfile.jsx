'use client';

import React, { useState, useEffect } from 'react';
import Input from '../ui/Input';
import Label from '../ui/Label';
import Button from '../ui/Button';
import FileInput from '../ui/FileInput';
import Modal from '../ui/Modal';
import { useForm } from 'react-hook-form';
import { useUser } from '@/app/context/UserContext';
import PageTitle from '../ui/PageTitle';

function EditProfile() {
  const { user } = useUser();

  const { register, setValue } = useForm();

  // Ajout du state pour gérer l'ouverture du modal
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    if (user) {
      setValue('firstname', user.firstname || '');
      setValue('lastname', user.lastname || '');
      setValue('username', user.username || '');
      setValue(
        'birthdate',
        user.birthdate
          ? new Date(user.birthdate).toISOString().split('T')[0]
          : ''
      );
      setValue('about', user.description || '');
      setValue('privacy', user.privacy || 'public');
    }
  }, [user, setValue]);

  return (
    <form className="flex flex-col gap-2.5 max-w-[600px] w-full">
      <div className="flex items-center justify-center">
        <PageTitle>Edit Profile</PageTitle>
      </div>
      <div className="flex flex-col lg:flex-row gap-2.5 w-full">
        <Label htmlFor={'firstname'}>
          First Name* :
          <Input
            id="firstname"
            placeholder="Enter your firstname..."
            {...register('firstname')}
          />
        </Label>
        <Label htmlFor={'lastname'}>
          Last Name* :
          <Input
            id="lastname"
            placeholder="Enter your lastname..."
            {...register('lastname')}
          />
        </Label>
      </div>
      <div className="flex flex-col items-start gap-2.5 w-full">
        <Label htmlFor={'username'}>Username* :</Label>
        <Input
          id="username"
          placeholder="Enter your username..."
          {...register('username')}
        />
      </div>
      <div className="flex flex-col items-start gap-2.5 w-full">
        <Label htmlFor={'birthdate'}>Birthdate* :</Label>
        <Input
          type="date"
          id="birthdate"
          name="birthdate"
          required
          {...register('birthdate')}
        />
      </div>
      <div className="flex flex-col lg:flex-row w-full gap-2">
        <div className="flex-1">
          <Label htmlFor="password" className="block mb-2">
            Password :
          </Label>
          <Input
            type="password"
            id="password"
            name="password"
            placeholder="********"
            {...register('password')}
          />
        </div>
        <div className="flex-1">
          <Label htmlFor="confirm_password" className="block mb-2">
            Confirm Password :
          </Label>
          <Input
            type="password"
            id="confirm_password"
            name="confirm_password"
            placeholder="********"
            {...register('confirm_password')}
          />
        </div>
      </div>
      <div className="flex flex-col items-start gap-2.5 w-full">
        <Label htmlFor="about">About me :</Label>
        <Input
          type="textarea"
          id="about"
          name="about"
          placeholder="Tell us about yourself..."
          className="h-24"
          {...register('about')}
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
              {...register('privacy')}
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
          type="button" // <-- Important : empêche la soumission du formulaire
          className="w-full"
          onClick={() => setIsModalOpen(true)}
        >
          Save Changes
        </Button>
      </div>
      <Modal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        title="Confirmation"
        message="Acceptez-vous les modifications ?"
        actions={[
          {
            label: 'Oui',
            onClick: () => {
              /* ajouter la logique de sauvegarde ici */
            },
            closeOnClick: true,
          },
          {
            label: 'Non',
            onClick: null,
            closeOnClick: true,
          },
        ]}
      />
    </form>
  );
}

export default EditProfile;
