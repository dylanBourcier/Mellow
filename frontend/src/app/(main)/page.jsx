import ProtectedRoute from '../components/auth/ProtectedRoute';
import PostsContainer from '../components/layout/PostsContainer';
import PageTitle from '../components/ui/PageTitle';

export const metadata = {
  title: 'Home - Mellow',
};

export default function HomePage() {
  return (
    <ProtectedRoute redirectTo="/login">
      <div>
        <PageTitle>Home</PageTitle>
        <PostsContainer />
      </div>
    </ProtectedRoute>
  );
}
