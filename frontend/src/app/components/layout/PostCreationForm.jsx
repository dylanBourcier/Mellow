'use client';

import React from 'react';
import InputSelect from '../ui/InputSelect';
import Label from '../ui/Label';
import Input from '../ui/Input';
import FileInput from '../ui/FileInput';
import Button from '../ui/Button';
import { useForm } from 'react-hook-form';
import SelectFollowers from '../ui/SelectFollowers';
import { Controller } from 'react-hook-form';
import toast from 'react-hot-toast';
import CustomToast from '@/app/components/ui/CustomToast';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';

function PostCreationForm() {
  const Router = useRouter();
  const [groupOptions, setGroupOptions] = useState([
    { value: 'everyone', label: 'Everyone' },
  ]);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const res = await fetch('/api/groups/joined', {
          credentials: 'include',
        });
        const data = await res.json();
        if (data.status === 'error') {
          throw new Error(data.message);
        }
        if (data.data === null) {
          setGroupOptions([{ value: 'everyone', label: 'Everyone' }]);
          return;
        }

        const groupMapped = data.data.map((group) => ({
          value: group.group_id,
          label: group.title,
        }));

        setGroupOptions((prev) => [
          { value: 'everyone', label: 'Everyone' },
          ...groupMapped,
        ]);
      } catch (err) {
        console.error('Failed to fetch groups:', err);
      }
    };

    fetchGroups();
  }, []);
  const {
    control,
    handleSubmit,
    setValue,
    watch,
    register,
    formState: { errors },
  } = useForm({
    defaultValues: {
      selectedFollowers: [],
      visibility: 'public', // 👈 ajouter
      title: '',
      postOn: 'everyone',
      content: '',
      image: null, // 👈 ajouter
    },
  });
  const postOn = watch('postOn');
  const visibility = watch('visibility');

  const onSubmit = async (data) => {
    try {
      const formData = new FormData();

      formData.append('title', data.title);
      formData.append('content', data.content);
      formData.append('visibility', data.visibility);
      formData.append('postOn', data.postOn);
      if (data.image) {
        formData.append('image', data.image);
      }

      formData.append(
        'selectedFollowers',
        JSON.stringify(data.selectedFollowers)
      );
      const res = await fetch('/api/posts', {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });
      const result = await res.json();
      console.log(result);
      if (result.status === 'error') {
        throw new Error(result.message);
      }
      toast.custom((t) => (
        <CustomToast
          t={t}
          type="success"
          message="Post created successfully!"
        />
      ));
      Router.push('/');
    } catch (error) {
      toast.custom((t) => (
        <CustomToast t={t} type="error" message="Error creating post!" />
      ));
    }
  };

  const followers = [];
  return (
    <form
      className="flex flex-col gap-6 w-full max-w-[600px]"
      onSubmit={handleSubmit(onSubmit)}
    >
      <div className="flex flex-col gap-2.5">
        <div>
          <Label htmlFor={'postOn'}>Post On* :</Label>
          <Controller
            name="postOn"
            control={control}
            render={({ field }) => (
              <InputSelect
                id="postOn"
                options={groupOptions}
                value={field.value}
                onChange={field.onChange}
              />
            )}
          />
        </div>
        <div>
          <Label htmlFor={'title'}>Title* :</Label>
          <Input
            placeholder="Enter your title here ..."
            name="title"
            id="title"
            {...register('title', { required: true })}
          ></Input>
        </div>
        <div>
          <Label htmlFor={'content'}>Content* :</Label>
          <Input
            placeholder="Enter your content here ... "
            type="textarea"
            name="content"
            id="content"
            {...register('content', { required: true })}
          ></Input>
        </div>
        <div>
          <Label htmlFor={'image'}>Image :</Label>
          <Controller
            name="image"
            control={control}
            render={({ field }) => (
              <FileInput
                id="image"
                name="image"
                setValue={setValue}
                onChange={(file) => field.onChange(file)}
              />
            )}
          />
        </div>

        {postOn === 'everyone' && (
          <>
            <Label htmlFor={'visibility'}>Visibility* :</Label>
            <div>
              <div className="flex gap-1 justify-center">
                <div className="flex-1">
                  <input
                    type="radio"
                    name="visibility"
                    id="public"
                    className="hidden peer"
                    value="public"
                    {...register('visibility', { required: true })}
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
                    name="visibility"
                    id="almost-private"
                    className="hidden peer"
                    value="almost-private"
                    {...register('visibility', { required: true })}
                  />
                  <label
                    htmlFor="almost-private"
                    className="flex align-middle justify-center py-2 border border-lavender-5 text-lavender-5 bg-transparent rounded-2xl peer-checked:bg-lavender-6 cursor-pointer transition-colors duration-200"
                  >
                    Almost Private
                  </label>
                </div>
                <div className="flex-1">
                  <input
                    type="radio"
                    name="visibility"
                    id="private"
                    className="hidden peer"
                    value="private"
                    {...register('visibility', { required: true })}
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
            {visibility === 'public' && (
              <span className="italic text-sm">
                Your post will be visible to everyone.
              </span>
            )}
            {visibility === 'almost-private' && (
              <span className="italic text-sm">
                Your post will be visible to your followers.
              </span>
            )}
            {visibility === 'private' && (
              <span className="italic text-sm">
                Your post will be visible only to selected followers.
              </span>
            )}
          </>
        )}
        {postOn === 'everyone' && visibility === 'private' && (
          <div className="flex flex-col gap-2.5">
            <Label htmlFor={''}>Who can see your post* :</Label>
            <Controller
              name="selectedFollowers"
              control={control}
              render={({ field }) => (
                <SelectFollowers
                  followers={followers}
                  value={field.value}
                  onChange={field.onChange}
                />
              )}
            />
          </div>
        )}
      </div>
      <Button type="submit" className="w-full" isSecondary={false}>
        Post
      </Button>
    </form>
  );
}

export default PostCreationForm;
