<script lang="ts">
	import '../app.postcss';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';

	import { AppRail, AppRailAnchor } from '@skeletonlabs/skeleton';

	import { CalendarClock, BellRing, History, House, Users, User as UserIcon } from 'lucide-svelte';
	import { onMount } from 'svelte';

	import { initializeStores, Toast } from '@skeletonlabs/skeleton';
	import { setUserContext, User } from '$lib/models/user.svelte';
	import { goto } from '$app/navigation';

	initializeStores();

	let { children } = $props();
	let isOpen = $state(false);
	let dropdownRef;
	let currentTile: number = $state(0);

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	onMount(() => {
		dropdownRef = document.querySelector('.dropdown-menu');
	});

	let user = new User();
	setUserContext(user);

	let isLoggedIn = $derived(user.isLoggedIn);
	// let userDate = $derived(userValue.date);

	function toggleDropdown() {
		isOpen = !isOpen;
	}

	function login() {
		goto('/login');
		isOpen = false;
	}

	function signout() {
		localStorage.removeItem('token');
		user.isLoggedIn = false;
		goto('/');
	}
</script>

<Toast />

<div class="h-full w-full">
	<header class="bg-secondary-500 p-4">
		<div class="flex flex-row">
			<div class="px-3"><BellRing /></div>
			<span class="font-bold">Scrum</span><span>dinger</span>
			<!-- {isLoggedIn}
			{userDate} -->
			<span class="flex grow"></span>
			<div class="flex flex-row justify-right">
				<div class="flex justify-center"></div>
				<div class="relative inline-block text-left">
					<div>
						<button
							type="button"
							aria-expanded="true"
							aria-haspopup="true"
							onclick={toggleDropdown}
						>
							<UserIcon />
						</button>
					</div>

					{#if isOpen}
						<!--
					  Dropdown menu, show/hide based on menu state.
				  
					  Entering: "transition ease-out duration-100"
						From: "transform opacity-0 scale-95"
						To: "transform opacity-100 scale-100"
					  Leaving: "transition ease-in duration-75"
						From: "transform opacity-100 scale-100"
						To: "transform opacity-0 scale-95"
					-->

						<div
							class="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black/5 focus:outline-none"
							role="menu"
							aria-orientation="vertical"
							aria-labelledby="menu-button"
							tabindex="-1"
						>
							{#if isLoggedIn}
								<div class="py-1" role="none">
									<!-- Active: "bg-gray-100 text-gray-900 outline-none", Not Active: "text-gray-700" -->
									<a
										href="#"
										class="block px-4 py-2 text-sm text-gray-700"
										role="menuitem"
										tabindex="-1"
										id="menu-item-0">Account settings</a
									>
									<a
										href="#"
										class="block px-4 py-2 text-sm text-gray-700"
										role="menuitem"
										tabindex="-1"
										id="menu-item-1">Support</a
									>
									<a
										href="#"
										class="block px-4 py-2 text-sm text-gray-700"
										role="menuitem"
										tabindex="-1"
										id="menu-item-2">License</a
									>
									<form method="POST" onsubmit={signout} role="none">
										<button
											type="submit"
											class="block w-full px-4 py-2 text-left text-sm text-gray-700"
											role="menuitem"
											tabindex="-1"
											id="menu-item-3">Sign out</button
										>
									</form>
								</div>
							{/if}

							{#if !isLoggedIn}
								<div class="py-1" role="none">
									<!-- Active: "bg-gray-100 text-gray-900 outline-none", Not Active: "text-gray-700" -->
									<div
										class="block px-4 py-2 text-sm text-gray-700"
										role="menuitem"
										tabindex="-1"
										id="menu-item-0"
									>
										<button onclick={login} class="w-full text-left"> Login </button>
									</div>
								</div>
							{/if}
						</div>
					{/if}
				</div>
			</div>
		</div>
	</header>

	<!-- Grid Columns -->
	<div class="flex h-full">
		{#if isLoggedIn}
			<aside class="sticky h-full">
				<AppRail class="">
					<!-- --- -->

					<AppRailAnchor href="/" bind:group={currentTile} name="tile-1" value={0} title="tile-1">
						<svelte:fragment>
							<div class="flex justify-center"><House /></div>
						</svelte:fragment>

						<span>Home</span>
					</AppRailAnchor>

					<AppRailAnchor
						href="/scrums"
						bind:group={currentTile}
						name="tile-1"
						value={1}
						title="tile-1"
					>
						<svelte:fragment>
							<div class="flex justify-center"><CalendarClock /></div>
						</svelte:fragment>
						<span>Scrums</span>
					</AppRailAnchor>
					<AppRailAnchor
						href="/teams"
						bind:group={currentTile}
						name="tile-2"
						value={2}
						title="tile-2"
					>
						<svelte:fragment>
							<div class="flex justify-center"><Users /></div>
						</svelte:fragment>
						<span>Teams</span>
					</AppRailAnchor>
					<AppRailAnchor
						href="/history"
						bind:group={currentTile}
						name="tile-3"
						value={3}
						title="tile-3"
					>
						<svelte:fragment>
							<div class="flex justify-center"><History /></div>
						</svelte:fragment>
						<span>History</span>
					</AppRailAnchor>
				</AppRail>
			</aside>
		{/if}
		<!-- Main Content -->
		<div class="w-full">
			{@render children()}
		</div>
	</div>
</div>
