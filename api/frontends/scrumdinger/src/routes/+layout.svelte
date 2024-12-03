<script lang="ts">
	import '../app.postcss';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';

	import { AppRail, AppRailAnchor } from '@skeletonlabs/skeleton';

	import { CalendarClock, BellRing, History, House, Users, User as UserIcon } from 'lucide-svelte';
	import { onMount } from 'svelte';

	import { initializeStores, Toast, Drawer, getDrawerStore } from '@skeletonlabs/skeleton';
	import { setUserContext, User } from '$lib/models/user.svelte';
	import CRUDScrum from '$lib/components/CRUDScrum.svelte';

	import { goto } from '$app/navigation';
	import { setScrumContext, Scrum } from '$lib/models/scrum.svelte';
	import { DrawerMeta, setDrawerContext } from '$lib/models/drawer.svelte';
	import HistoryScrum from '$lib/components/HistoryScrum.svelte';

	initializeStores();

	let { children } = $props();
	let isOpen = $state(false);
	let dropdownRef;
	let currentTile: number = $state(0);

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	onMount(() => {
		dropdownRef = document.querySelector('.dropdown-menu');
	});

	// set contexts

	let user = new User();
	setUserContext(user);
	let scrum = new Scrum();
	setScrumContext(scrum);
	let drawerMeta = new DrawerMeta();
	setDrawerContext(drawerMeta);

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
		scrum.meetings.length = 0;
		isOpen = false;
		goto('/');
	}
</script>

<Toast />

<Drawer position="right" width="max-w-[500px]">
	{#if drawerMeta.data['component'] === 'CRUDScrum'}
		<CRUDScrum></CRUDScrum>
	{:else if drawerMeta.data['component'] === 'HistoryScrum'}
		<HistoryScrum></HistoryScrum>
	{/if}
</Drawer>

<div class="flex flex-col h-screen">
	<header class="h-[64px] max-h-[64px] bg-secondary-500 p-4">
		<div class="flex flex-row">
			<div class="px-3"><BellRing /></div>
			<span class="font-bold">Scrum</span><span>dinger</span>
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
						<div
							class="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black/5 focus:outline-none"
							role="menu"
							aria-orientation="vertical"
							aria-labelledby="menu-button"
							tabindex="-1"
						>
							{#if user.isLoggedIn}
								<div class="py-1" role="none">
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
							{:else}
								<div class="py-1" role="none">
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

	<div class="w-full h-[calc(100vh-64px)]">
		{@render children()}
	</div>
</div>
