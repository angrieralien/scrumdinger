import { purgeCss } from 'vite-plugin-tailwind-purgecss';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig, searchForWorkspaceRoot } from 'vite';

export default defineConfig({
	plugins: [sveltekit(), purgeCss()],
	server: {
		fs: {
			allow: [searchForWorkspaceRoot(process.cwd()), '/static']
		},
		proxy: {
			'/v1/auth/token': {
				target: 'http://localhost:6000',
				changeOrigin: true
			},
			'/v1': {
				target: 'http://localhost:3000',
				changeOrigin: true
			}
		}
	}
});
