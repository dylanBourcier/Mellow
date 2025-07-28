'use client';

import React, { use } from 'react';
import { useEffect, useState } from 'react';
import Spinner from '../ui/Spinner';
import PostCard from '../ui/PostCard';
import CustomToast from '../ui/CustomToast';
import { toast } from 'react-hot-toast';

function PostsGroupContainer({ groupId }) {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [isMember, setIsMember] = useState(false);

  useEffect(() => {
    if (!groupId) return;

    fetch(`/api/groups/posts/${groupId}`, { credentials: 'include' }) // adapte l'URL selon ton backend
      .then((res) => {
        if (!res.ok) throw new Error('Failed to fetch group posts');
        return res.json();
      })
      .then((data) => {
        console.log(data);
        if (data.status == 'error') {
          throw new Error(data.message || 'Failed to fetch group posts');
        }
        if (data.message === 'Not member') {
          setLoading(false);
          setIsMember(false);
          return;
        }
        setIsMember(true);
        setPosts(data.data || []);
        console.log('Posts fetched successfully:', data.data);

        setLoading(false);
      })
      .catch((err) => {
        setLoading(false);
        toast.custom((t) => (
          <CustomToast
            t={t}
            type="error"
            message="An error occurred while fetching group posts. Please try again later."
          />
        ));
      });
  }, [groupId]);

  if (loading)
    return (
      <div>
        <Spinner></Spinner> <span>Loading posts...</span>
      </div>
    );

  return (
    <div className="flex flex-col gap-3 w-full items-center">
      {isMember ? (
        posts.map((post) => <PostCard key={post.post_id} post={post} />)
      ) : (
        <span className="text-dark-grey-lighter">
          You are not member of this group. You can't see the posts
        </span>
      )}
    </div>
  );
}

export default PostsGroupContainer;
