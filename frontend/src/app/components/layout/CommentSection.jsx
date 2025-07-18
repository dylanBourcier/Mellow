import React from 'react';
import { useUser } from '@/app/context/UserContext';
import CommentCard from '../ui/CommentCard';
import PageTitle from '../ui/PageTitle';
import Input from '../ui/Input';
import FileInput from '../ui/FileInput';
import Button from '../ui/Button';
import { icons } from '@/app/lib/icons';
import { useForm, Controller, set } from 'react-hook-form';
import toast from 'react-hot-toast';
import CustomToast from '../ui/CustomToast';
import { formatDate } from '@/app/utils/date';
import { useEffect } from 'react';
import Spinner from '../ui/Spinner';

export default function CommentSection({ postid, initialCommentCount }) {
  const [comments, setComments] = React.useState([]);
  const [commentCount, setCommentCount] = React.useState(
    initialCommentCount || 0
  );
  const [loading, setLoading] = React.useState(true);
  const { user } = useUser();

  const {
    register,
    handleSubmit,
    reset,
    watch,
    setValue,
    formState: { errors },
    control,
  } = useForm({
    defaultValues: {
      content: '',
      image: null,
    },
  });
  const onSubmit = async (data) => {
    try {
      const formData = new FormData();
      formData.append('content', data.content);
      if (data.image) {
        formData.append('image', data.image);
      }
      const res = await fetch(`/api/comments/${postid}`, {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });
      const result = await res.json();
      console.log('user', user);

      if (result.status === 'error') {
        throw new Error(result.message);
      }
      toast.custom((t) => (
        <CustomToast
          t={t}
          type="success"
          message="Comment added successfully!"
        />
      ));
      setCommentCount((prev) => prev + 1);
      const newComment = result.data;
      newComment.user_id = user.id; // Ensure user_id is set correctly
      newComment.username = user.username; // Ensure username is set correctly
      newComment.avatar_url = user.image_url; // Ensure avatar_url is set correctly
      newComment.creation_date = formatDate(newComment.creation_date);
      setComments((prev) => [...prev, newComment]);
      reset();
    } catch (error) {
      toast.custom((t) => (
        <CustomToast t={t} type="error" message="Error creating comment!" />
      ));
    }
  };

  useEffect(() => {
    const fetchComments = async () => {
      try {
        const response = await fetch(`/api/comments/${postid}`);
        if (!response.ok) {
          throw new Error(
            `Failed to fetch comments (status: ${response.status})`
          );
        }
        const result = await response.json();
        if (result.status === 'error') {
          throw new Error(result.message || 'Failed to fetch comments');
        }
        setComments(result.data);
        setCommentCount(result.data.length);
      } catch (error) {
        toast.custom((t) => (
          <CustomToast t={t} type="error" message="Error loading comments!" />
        ));
      } finally {
        setLoading(false);
      }
    };

    fetchComments();
  }, [postid]);
  return (
    <div className="flex flex-col gap-3 px-2 lg:px-8 py-2.5">
      <PageTitle className="flex text-left">
        Comment{commentCount > 1 ? 's' : ''} ({commentCount || 0})
      </PageTitle>
      {user && (
        <form
          className="flex gap-1 items-center"
          onSubmit={handleSubmit(onSubmit)}
        >
          <Input
            type="text"
            name={'content'}
            {...register('content', { required: true })}
            placeholder="Post a comment..."
            className="border border-lavender-5"
          ></Input>
          {errors.content && (
            <span className="text-red-400 text-xs">Content is required</span>
          )}
          <Controller
            name="image"
            control={control}
            render={({ field }) => (
              <FileInput
                id={'image'}
                accept="image/*"
                name={'image'}
                label={icons['image']}
                usePreview={false}
                isMini
                setValue={setValue}
                onChange={(file) => field.onChange(file)}
              />
            )}
          ></Controller>
          <Button type="submit">Comment</Button>
        </form>
      )}
      {loading ? (
        <div className="flex w-full items-center gap-2 justify-center ">
          <Spinner></Spinner>Loading comments...
        </div>
      ) : (
        (comments || []).map((comment, index) => (
          <CommentCard key={index} comment={comment}></CommentCard>
        ))
      )}
    </div>
  );
}
