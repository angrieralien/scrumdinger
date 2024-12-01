<script lang="ts">
	import { getUserContext, User } from '$lib/models/user.svelte';
	import { Plus } from 'lucide-svelte';
	import { mount, onMount } from 'svelte';

	let user: User;

	onMount(() => {
		user = getUserContext();
	});
</script>

<div class="h-full">
	{#if user && user.isLoggedIn}
		<div></div>
	{:else}
		<div class="flex flex-col items-center h-full">
			<div class="text-center p-6 text-2xl">
				Ahh Snap!!! <br />Login to view your scrums or create a new meeting.
			</div>
		</div>
	{/if}

	<button type="button" class="add-button bg-secondary-500 m-3 btn-icon btn-icon-xl">
		<Plus></Plus>
	</button>
</div>

<style>
	.add-button {
		position: fixed;
		right: 0;
		bottom: 0;
	}
</style>
