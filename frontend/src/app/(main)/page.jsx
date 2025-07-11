import Button from '../components/ui/Button';
import PageTitle from '../components/ui/PageTitle';
import PostCard from '../components/ui/PostCard';

export const metadata = {
  title: 'Home - Mellow',
};

export default function HomePage() {
  const postId = 1;
  const postTitle = 'Sample Post Title';
  const postContent =
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod.';
  const authorAvatar = '/img/lion.png'; // Example avatar image
  const userName = 'johndoe';
  const date = '2023-10-01';
  const Comments = 5;
  const props = {
    postId,
    postTitle,
    postContent,
    authorAvatar,
    userName,
    date,
    Comments,
  };
  return (
    <div>
      <PageTitle>Home</PageTitle>
      <div className="flex flex-col gap-3">
        <PostCard postInfos={props} />
        <PostCard postInfos={props} />
        <PostCard postInfos={props} />
        <PostCard postInfos={props} />
      </div>
    </div>
  );
}
