import React from "react";
import PostCreationForm from "@/app/components/layout/PostCreationForm";

const metadata = {
  title: {
    template: "%s - Mellow",
    default: "Create Post",
  },
  description:
    "Create a new post on Mellow, a social media platform for developers to share their projects and connect with others.",
};
export { metadata };

export default function PostCreationPage() {
  return (
    <div className="flex flex-col items-center min-h-screen">
      <PostCreationForm />
    </div>
  );
}
