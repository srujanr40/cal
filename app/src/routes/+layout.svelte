<script>
	import '../app.css';
</script>

<div class="bg-background fixed left-0 right-0 top-0 z-50 mx-auto">
	<div class="text-foreground mx-auto flex justify-start gap-5 p-5 pl-10">
		<a class="font-medium leading-3" href="/">Cal</a>
		<a class="font-medium leading-3" href="/login">Login</a>
		<a class="font-medium leading-3" href="/register">Register</a>
	</div>
	<div class="bg-border h-[0.7px] w-full"></div>
</div>

<slot></slot>
