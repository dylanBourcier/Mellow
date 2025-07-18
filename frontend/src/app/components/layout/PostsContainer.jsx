'use client';

import React from 'react';
import { useEffect, useState } from 'react';
import Spinner from '../ui/Spinner';
import PostCard from '../ui/PostCard';

export default function PostsContainer() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/api/posts', { credentials: 'include' }) // adapte l'URL selon ton backend
      .then((res) => {
        if (!res.ok) throw new Error('Failed to fetch posts');
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
          <CustomToast t={t} type="error" message="Error creating post!" />
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

// const postId = 1;
// const postTitle = 'Sample Post Title';
// const postContent =
//   'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod.';
// const authorAvatar = '/img/lion.png'; // Example avatar image
// const userName = 'johndoe';
// const date = '2023-10-01';
// const Comments = 5;
// const props = {
//   postId,
//   postTitle,
//   postContent,
//   authorAvatar,
//   userName,
//   date,
//   Comments,
// };
