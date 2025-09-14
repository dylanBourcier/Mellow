/** @type {import('next').NextConfig} */
const BACKEND_ORIGIN = process.env.BACKEND_ORIGIN || "http://localhost:3225";

const nextConfig = {
	output: "standalone",
	async rewrites() {
		return [
			{
				source: "/api/:path*",
				destination: `${BACKEND_ORIGIN}/:path*`, // backend Go (overridden via env at build time)
			},
		];
	},
};

export default nextConfig;
