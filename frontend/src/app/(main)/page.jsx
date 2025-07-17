import PostsContainer from '../components/layout/PostsContainer';
import PageTitle from '../components/ui/PageTitle';

export const metadata = {
  title: 'Home - Mellow',
};

export default function HomePage() {
  return (
    <div>
      <PageTitle>Home</PageTitle>
      <PostsContainer />
    </div>
  );
}
