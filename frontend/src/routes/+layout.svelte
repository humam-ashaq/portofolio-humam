<script lang="ts">
	import '../app.css';

	let { children } = $props();

	import Header from '$lib/components/layout/Header.svelte';
	import Footer from '$lib/components/layout/Footer.svelte';
</script>

<svelte:head>
	<title>Humam's Portfolio</title>

	<link rel="icon" type="image/x-icon" href="/ashaq.png" />

	<link rel="stylesheet" href="https://unicons.iconscout.com/release/v4.0.8/css/line.css" />
</svelte:head>

<Header />

<main>
	{@render children?.()}
</main>

<Footer />