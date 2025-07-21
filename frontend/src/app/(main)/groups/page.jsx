import GroupListContainer from '@/app/components/layout/GroupListContainer';

const metadata = {
  title: {
    template: '%s - Mellow',
    default: 'Groups',
  },
  description:
    'Explore and manage your groups on Mellow, a social media platform for developers to share their projects and connect with others.',
};
export { metadata };

export default function GroupsPage() {
  return <GroupListContainer />;
}
