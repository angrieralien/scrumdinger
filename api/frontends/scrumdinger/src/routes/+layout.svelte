<script lang="ts">
	import '../app.postcss';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';

	import { AppRail, AppRailTile, AppRailAnchor } from '@skeletonlabs/skeleton';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	let currentTile: number = 0;
</script>

<div class="grid h-screen grid-rows-[auto_1fr_auto]">

	<!-- Grid Columns -->
	<div class="grid grid-cols-1 grid-cols-[auto_1fr]">
		<aside class="sticky top-0 col-span-1 h-screen bg-yellow-500">
		<AppRail class="">
			<svelte:fragment slot="lead">
				<AppRailAnchor href="/" >(icon)</AppRailAnchor>
			</svelte:fragment>
			<!-- --- -->
			<AppRailTile bind:group={currentTile} name="tile-1" value={0} title="tile-1">
				<svelte:fragment slot="lead">(icon)</svelte:fragment>
				<span>Tile 1</span>
			</AppRailTile>
			<AppRailTile bind:group={currentTile} name="tile-2" value={1} title="tile-2">
				<svelte:fragment slot="lead">(icon)</svelte:fragment>
				<span>Tile 2</span>
			</AppRailTile>
			<AppRailTile bind:group={currentTile} name="tile-3" value={2} title="tile-3">
				<svelte:fragment slot="lead">(icon)</svelte:fragment>
				<span>Tile 3</span>
			</AppRailTile>
			<!-- --- -->
			<svelte:fragment slot="trail">
				<AppRailAnchor href="/" target="_blank" title="Account">(icon)</AppRailAnchor>
			</svelte:fragment>
		</AppRail>
	</aside>
	  <!-- Main Content -->
	  <main class="space-y-4 bg-green-500 p-4">
		<p class="h-[512px] bg-purple-500 p-4">Paragraph 1</p>
		<p class="h-[512px] bg-purple-500 p-4">Paragraph 2</p>
		<p class="h-[512px] bg-purple-500 p-4">Paragraph 3</p>
	  </main>
	</div>
  </div>