<script>
	import ScrumTimer from '$lib/components/ScrumTimer.svelte';

	let countdown = 30;
	let attendees = 3;
</script>

<div class="flex justify-center items-center h-full">
	<ScrumTimer {countdown} {attendees} />
</div>
