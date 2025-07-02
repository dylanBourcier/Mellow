import Button from '../components/ui/Button';
import PageTitle from '../components/ui/PageTitle';
import PostCard from '../components/ui/PostCard';

export const metadata = {
  title: 'Home - Mellow',
};

export default function HomePage() {
  return (
    <div>
      <PageTitle>Home</PageTitle>
      <div className='flex flex-col gap-3'>
      <PostCard></PostCard>
      <PostCard></PostCard>
      <PostCard></PostCard>
      <PostCard></PostCard>
      </div>
    </div>
  
  );
}
