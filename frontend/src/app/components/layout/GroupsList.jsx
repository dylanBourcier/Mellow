'use client';

import React from 'react';
import PageTitle from '../ui/PageTitle';
import Input from '../ui/Input';
import Button from '../ui/Button';

export default function GroupsList() {
  return (
    <div className="flex flex-col bg-white rounded-2xl shadow-(--box-shadow) p-4 h-full flex-2">
      <h2 className="text-lavender-5 text-shadow-(--text-shadow)">
        Join a new group
      </h2>
      <div className="flex items-center">
        <Input type="search" placeholder="Search for groups..." />
        <Button>Search</Button>
      </div>
    </div>
  );
}
