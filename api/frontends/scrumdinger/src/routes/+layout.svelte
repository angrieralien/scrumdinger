<script lang="ts">
	import '../app.postcss';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';

	import { AppRail, AppRailTile, AppRailAnchor } from '@skeletonlabs/skeleton';

	import { CalendarClock, CircleUser, History, House, Users } from 'lucide-svelte';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	let currentTile: number = 0;
</script>

<div class="grid h-screen grid-rows-[auto_1fr_auto]">
	<!-- Grid Columns -->
	<div class="grid grid-cols-1 grid-cols-[auto_1fr]">
		<aside class="sticky top-0 col-span-1 h-screen bg-yellow-500">
			<AppRail class="">
				<!-- --- -->
				<AppRailTile href="/" bind:group={currentTile} name="tile-0" value={0} title="tile-1">
					<svelte:fragment slot="lead">
						<div class="flex justify-center"><House /></div>
					</svelte:fragment>
					<span>Home</span>
				</AppRailTile>
				<AppRailTile bind:group={currentTile} name="tile-1" value={1} title="tile-1">
					<svelte:fragment slot="lead">
						<div class="flex justify-center"><CalendarClock /></div>
					</svelte:fragment>
					<span>Scrums</span>
				</AppRailTile>
				<AppRailTile bind:group={currentTile} name="tile-2" value={2} title="tile-2">
					<svelte:fragment slot="lead">
						<div class="flex justify-center"><Users /></div>
					</svelte:fragment>
					<span>Teams</span>
				</AppRailTile>
				<AppRailTile bind:group={currentTile} name="tile-3" value={3} title="tile-3">
					<svelte:fragment slot="lead">
						<div class="flex justify-center"><History /></div>
					</svelte:fragment>
					<span>History</span>
				</AppRailTile>
				<!-- --- -->
				<svelte:fragment slot="trail">
					<AppRailAnchor href="/" title="Account">
						<div class="flex justify-center"><CircleUser /></div>
					</AppRailAnchor>
				</svelte:fragment>
			</AppRail>
		</aside>
		<!-- Main Content -->
		<main class="space-y-4 bg-green-500 p-4">
			<slot />
		</main>
	</div>
</div>
