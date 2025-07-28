'use client';

import React from 'react';
import { useEffect, useState } from 'react';
import Spinner from '../ui/Spinner';
import PostCard from '../ui/PostCard';
import toast from 'react-hot-toast';
import CustomToast from '../ui/CustomToast';

export default function PostsContainer({ userId }) {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(`/api/users/posts/${userId}`, { credentials: 'include' }) // adapte l'URL selon ton backend
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        if (data.status !== 'success') {
          throw new Error(data.message || 'Failed to fetch posts');
        }
        setPosts(data.data || []);
        setLoading(false);
      })
      .catch((err) => {
        setLoading(false);
        toast.custom((t) => (
          <CustomToast
            t={t}
            type="error"
            message={'Error fetching posts! ' + err}
          />
        ));
      });
  }, []);

  if (loading)
    return (
      <div>
        <Spinner></Spinner> <span>Loading posts...</span>
      </div>
    );

  return (
    <div className="flex flex-col gap-3">
      {posts.map((post) => (
        <PostCard key={post.post_id} post={post} />
      ))}
    </div>
  );
}
