<script lang="ts">
	import { getLocalTimeZone, today } from '@internationalized/date';
	import { Calendar } from '$lib/components/ui/calendar/index.js';

	let value = today(getLocalTimeZone());
</script>

<div class="container mx-auto my-20 flex justify-center">
	<Calendar bind:value class="rounded-md border" />
</div>
