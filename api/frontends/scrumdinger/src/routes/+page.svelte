<script>
	import ScrumTimer from '$lib/components/ScrumTimer.svelte';

	let countdown = 30;
	let attendees = 3;

	
</script>

<div class="grid grid-cols-1">
	<ScrumTimer {countdown} {attendees} /> 
</div>
