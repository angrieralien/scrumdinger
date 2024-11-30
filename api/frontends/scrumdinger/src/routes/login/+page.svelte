<script lang="ts">
	import { BellRing } from 'lucide-svelte';
	import { userAPI } from '$lib/api/userapi';
	import { getToastStore, type ToastSettings } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';
	import { getUserContext, User } from '$lib/models/user.svelte';
	import { goto } from '$app/navigation';

	const toastStore = getToastStore();

	let user: User;
	let email: string = $state('admin@example.com');
	let password: string = $state('gophers');
	let t: ToastSettings = {
		message: '',
		timeout: 10000
	};

	onMount(() => {
		user = getUserContext();
	});

	/**
	 * submit sends login request to User api.
	 */
	function submit() {
		userAPI
			.login(email, password)
			.then((data) => {
				let token = data['token'];
				localStorage.setItem('token', token);
				user.isLoggedIn = true;
				goto('/');
			})
			.catch((reason: any) => {
				t.message = reason['message'];
				toastStore.trigger(t);
			});
	}
</script>

<div class="flex min-h-full flex-col justify-center px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-sm">
		<div class="flex justify-center">
			<BellRing />
		</div>
		<h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight">Sign in to your account</h2>
	</div>

	<div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
		<form class="space-y-6" onsubmit={submit}>
			<div>
				<label for="email" class="block text-sm/6 font-medium">Email address</label>
				<div class="mt-2">
					<input
						id="email"
						name="email"
						type="email"
						autocomplete="email"
						required
						class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset sm:text-sm/6"
						bind:value={email}
					/>
				</div>
			</div>

			<div>
				<div class="flex items-center justify-between">
					<label for="password" class="block text-sm/6 font-medium">Password</label>
					<div class="text-sm">
						<a href="#" class="font-semibold text-primary-600 hover:text-primary-500"
							>Forgot password?</a
						>
					</div>
				</div>
				<div class="mt-2">
					<input
						id="password"
						name="password"
						type="password"
						autocomplete="current-password"
						required
						class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset sm:text-sm/6"
						bind:value={password}
					/>
				</div>
			</div>

			<div>
				<button
					type="submit"
					class="flex w-full justify-center rounded-md bg-primary-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
					>Sign in</button
				>
			</div>
		</form>
	</div>
</div>
