<script lang="ts">
	import { onMount } from "svelte";
	import {
		Zap,
		Search,
		Terminal,
		FileCode,
		Lock,
		Gauge,
	} from "@lucide/svelte";
	import Navbar from "$lib/components/Navbar.svelte";
	import Hero from "$lib/components/Hero.svelte";
	import Features from "$lib/components/Features.svelte";
	import Documentation from "$lib/components/Documentation.svelte";
	import Contribute from "$lib/components/Contribute.svelte";
	import Footer from "$lib/components/Footer.svelte";

	let scrollY = $state(0);
	let copied = $state(false);
	let stars = $state<number | null>(null);

	const fetchStars = async () => {
		try {
			const response = await fetch(
				"https://api.github.com/repos/aritra1999/httpyum",
			);
			const data = await response.json();
			stars = data.stargazers_count;
		} catch (error) {
			console.error("Failed to fetch stars:", error);
		}
	};

	const copyInstallCommand = async () => {
		const command =
			"curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/install.sh | bash";
		try {
			await navigator.clipboard.writeText(command);
			copied = true;
			setTimeout(() => {
				copied = false;
			}, 2000);
		} catch (err) {
			console.error("Failed to copy:", err);
		}
	};

	onMount(() => {
		const handleScroll = () => {
			scrollY = window.scrollY;
		};

		window.addEventListener("scroll", handleScroll);
		fetchStars();
		return () => window.removeEventListener("scroll", handleScroll);
	});

	const features = [
		{
			icon: Zap,
			title: "Lightning Fast",
			description:
				"Built with Go. Instant startup, zero configuration. Just works.",
		},
		{
			icon: Terminal,
			title: "Beautiful TUI",
			description:
				"Bubbletea-powered interface with syntax highlighting and fuzzy search.",
		},
		{
			icon: FileCode,
			title: "Standard .http Files",
			description:
				"Use the same .http files you're already familiar with. No proprietary formats.",
		},
		{
			icon: Lock,
			title: "Environment Variables",
			description:
				"Load secrets from your shell. Variable substitution built-in.",
		},
		{
			icon: Search,
			title: "Interactive JSON",
			description:
				"Press 'f' to explore JSON responses with jless. Expand, collapse, search.",
		},
		{
			icon: Gauge,
			title: "Request Metrics",
			description:
				"Response timing, status codes, and request size at a glance.",
		},
	];

	const shortcuts = [
		{ keys: ["↑", "↓"], description: "Navigate requests" },
		{ keys: ["/"], description: "Fuzzy search" },
		{ keys: ["Enter"], description: "Execute request" },
		{ keys: ["f"], description: "Open JSON in jless" },
		{ keys: ["h"], description: "Toggle headers" },
		{ keys: ["v"], description: "Toggle variables" },
		{ keys: ["q"], description: "Quit" },
	];
</script>

<svelte:head>
	<title>httpyum — CLI REST Client</title>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="" />
	<link
		href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=JetBrains+Mono:wght@400;500;600&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<div
	class="bg-white text-black font-['Space_Grotesk',sans-serif] leading-normal"
>
	<Navbar />
	<Hero {stars} />
	<Features {features} />
	<Documentation {shortcuts} {copied} onCopy={copyInstallCommand} />
	<Contribute />
	<Footer />
</div>
