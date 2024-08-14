<script lang="ts">
    import { Button } from '$lib/components/ui/button/index.js';
    import DemoCalendar from '$lib/components/homepage-demo/DemoCalendar.svelte';
</script>

<div class="container mx-auto text-center mt-44">
    <div class="flex flex-row items-center justify-center">
        <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-7xl">Plan things</h1>
        <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-7xl pl-3 text-primary">easier</h1>
    </div>
    <p class="text-muted-foreground text-xl mt-6">
        An intuitive calendar that syncs across accounts, and displays availabilites. <br> Find time in common with people that count. 
    </p>
    <div class="flex flex-row items-center justify-center mt-7 mx-auto gap-4">
        <Button href="/login">Get Started</Button>
        <Button variant="outline">Learn More</Button>
    </div>
    <DemoCalendar />
</div>

