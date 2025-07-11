/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://localhost:3225/:path*', // backend Go
      },
    ];
  },
};

export default nextConfig;
