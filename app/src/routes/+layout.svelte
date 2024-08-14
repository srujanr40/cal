<script lang="ts">
	import '../app.css';

	import { ModeWatcher } from 'mode-watcher';

	import Sun from 'lucide-svelte/icons/sun';
	import Moon from 'lucide-svelte/icons/moon';
	import { CalendarDays } from 'lucide-svelte';

	import { resetMode, setMode } from 'mode-watcher';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
</script>

<ModeWatcher />

<div class="bg-background fixed left-0 right-0 top-0 z-50 mx-auto">
	<div class="mx-auto flex items-center justify-between p-2 pl-10">
		<div class="block flex gap-6">
			<div class="text-foreground inline-flex items-center font-medium leading-3">
				<a href="/">
					<CalendarDays class="cursor-pointer" />
				</a>
				<a href="/" class="pl-2 font-bold ">Syncal</a>
			</div>
			<div class="text-foreground flex items-center font-medium leading-3">
				<a href="/">Tutorial</a>
			</div>
		</div>
		<div class="inline-block pr-3">
			<DropdownMenu.Root>
				<DropdownMenu.Trigger asChild let:builder>
					<Button builders={[builder]} variant="outline" size="icon">
						<Sun
							class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
						/>
						<Moon
							class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
						/>
						<span class="sr-only">Toggle theme</span>
					</Button>
				</DropdownMenu.Trigger>
				<DropdownMenu.Content align="end">
					<DropdownMenu.Item on:click={() => setMode('light')}>Light</DropdownMenu.Item>
					<DropdownMenu.Item on:click={() => setMode('dark')}>Dark</DropdownMenu.Item>
					<DropdownMenu.Item on:click={() => resetMode()}>System</DropdownMenu.Item>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</div>
	</div>

	<div class="bg-border h-[0.7px] w-full"></div>
</div>

<slot></slot>
