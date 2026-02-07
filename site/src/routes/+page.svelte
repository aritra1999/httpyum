<script lang="ts">
	import { onMount } from "svelte";
	import {
		ArrowRight,
		Download,
		Github,
		BookOpen,
		Zap,
		Search,
		Code2,
		Terminal,
		FileCode,
		Lock,
		Gauge,
		Copy,
		Check,
		Star,
		Bug,
		Lightbulb,
	} from "@lucide/svelte";
	import { Button } from "$lib/components/ui/button/index.js";
	import * as Tabs from "$lib/components/ui/tabs/index.js";

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

<div class="page">
	<!-- Navbar -->
	<nav class="navbar">
		<div class="navbar-container">
			<a href="/" class="navbar-logo">
				<svg
					width="32"
					height="32"
					viewBox="0 0 32 32"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
				>
					<rect
						x="4"
						y="4"
						width="24"
						height="24"
						rx="4"
						stroke="currentColor"
						stroke-width="2"
					/>
					<path
						d="M11 16h10M16 11v10"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
					/>
				</svg>
			</a>
			<a
				href="https://github.com/aritra1999/httpyum"
				target="_blank"
				rel="noopener noreferrer"
				class="navbar-github"
			>
				<Github size={20} />
			</a>
		</div>
	</nav>

	<!-- Hero Section -->
	<section class="hero">
		<div class="container hero-grid">
			<div class="hero-content">
				<h1 class="hero-title">httpyum</h1>
				<p class="hero-subtitle">
					Fast, interactive CLI tool for executing HTTP requests from
					.http files
					{#if stars !== null}
						<span class="hero-stars">
							· <a
								href="https://github.com/aritra1999/httpyum/stargazers"
								target="_blank"
								rel="noopener noreferrer"
								class="stars-link"
							>
								<Star size={14} fill="currentColor" />
								{stars.toLocaleString()}
							</a>
						</span>
					{/if}
				</p>
				<div class="hero-buttons">
					<Button
						href="https://github.com/aritra1999/httpyum"
						variant="default"
						size="lg"
						class="gap-2"
					>
						<Github size={20} />
						<span>View on GitHub</span>
					</Button>
					<Button
						href="#docs"
						variant="outline"
						size="lg"
						class="gap-2"
					>
						<BookOpen size={20} />
						<span>Docs</span>
					</Button>
				</div>
			</div>
			<div class="hero-demo">
				<div class="demo-wrapper">
					<img
						src="https://github.com/user-attachments/assets/e4313241-a4c8-4c80-a422-1fde7f953bcc"
						alt="httpyum demo"
						class="demo-video"
					/>
				</div>
			</div>
		</div>
		<a href="#features" class="scroll-indicator">
			<div class="scroll-indicator-line"></div>
			<div class="scroll-indicator-dot"></div>
		</a>
	</section>

	<!-- Features -->
	<section class="features-section" id="features">
		<div class="container">
			<h2 class="section-title">Features</h2>
			<div class="features-list">
				{#each features as feature, i}
					{@const Icon = feature.icon}
					<div class="feature-item">
						<div class="feature-number">
							{String(i + 1).padStart(2, "0")}
						</div>
						<div class="feature-icon-minimal">
							<Icon size={20} strokeWidth={2} />
						</div>
						<div class="feature-content">
							<h3 class="feature-title">{feature.title}</h3>
							<p class="feature-description">
								{feature.description}
							</p>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</section>

	<!-- Documentation -->
	<section class="docs-section" id="docs">
		<div class="container">
			<h2 class="section-title">Documentation</h2>

			<Tabs.Root value="getting-started" class="docs-interactive">
				<div class="tabs-wrapper">
					<Tabs.List class="docs-tabs">
						<Tabs.Trigger value="getting-started" class="doc-tab">
							Getting Started
						</Tabs.Trigger>
						<Tabs.Trigger value="file-format" class="doc-tab">
							File Format
						</Tabs.Trigger>
						<Tabs.Trigger value="variables" class="doc-tab">
							Variables
						</Tabs.Trigger>
						<Tabs.Trigger value="json-viewer" class="doc-tab">
							JSON Viewer
						</Tabs.Trigger>
						<Tabs.Trigger value="shortcuts" class="doc-tab">
							Keyboard Shortcuts
						</Tabs.Trigger>
					</Tabs.List>
				</div>

				<Tabs.Content value="getting-started" class="doc-panel">
					<h3 class="doc-panel-title">Getting Started</h3>
					<div class="doc-content">
						<h4 class="doc-section-title">Installation</h4>
						<p>Install httpyum using the installation script:</p>
						<div class="install-card-inline">
							<pre
								class="install-code">curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/install.sh | bash</pre>
							<Button
								variant="ghost"
								size="sm"
								class="copy-btn"
								onclick={copyInstallCommand}
							>
								{#if copied}
									<Check size={16} />
								{:else}
									<Copy size={16} />
								{/if}
							</Button>
						</div>
						<div class="install-platforms-inline">
							<span class="platform-inline">macOS</span>
							<span class="platform-inline">Linux</span>
							<span class="platform-inline">Windows</span>
						</div>

						<h4 class="doc-section-title">Usage</h4>
						<p>Run httpyum with any .http file:</p>
						<pre class="doc-code">httpyum example.http</pre>
						<p>Command-line options:</p>
						<ul class="doc-list">
							<li>
								<code>--no-headers</code> - Hide response headers
							</li>
							<li>
								<code>-h, --help</code> - Show help message
							</li>
							<li>
								<code>-v, --version</code> - Show version
							</li>
						</ul>
					</div>
				</Tabs.Content>

				<Tabs.Content value="file-format" class="doc-panel">
					<h3 class="doc-panel-title">File Format</h3>
					<div class="doc-content">
						<p>Standard .http file format with variables:</p>
						<pre class="doc-code">@variable = value

### Request Name
METHOD url
Header: value

body</pre>
					</div>
				</Tabs.Content>

				<Tabs.Content value="variables" class="doc-panel">
					<h3 class="doc-panel-title">Variables</h3>
					<div class="doc-content">
						<p>
							Define variables with <code>@</code> and use with
							<code>{"{{"}{"}}"}</code>:
						</p>
						<pre class="doc-code">@baseUrl = https://api.example.com
GET {"{{"}baseUrl{"}}"}/users</pre>
						<p>Load from environment:</p>
						<pre
							class="doc-code">@token = {"{{"}$dotenv JWT{"}}"}</pre>
					</div>
				</Tabs.Content>

				<Tabs.Content value="json-viewer" class="doc-panel">
					<h3 class="doc-panel-title">JSON Viewer</h3>
					<div class="doc-content">
						<p>
							Press <kbd class="doc-key">f</kbd> to open JSON responses
							in jless for interactive exploration.
						</p>
						<p>Install jless for the best experience:</p>
						<pre class="doc-code"># macOS
brew install jless

# Linux
cargo install jless</pre>
					</div>
				</Tabs.Content>

				<Tabs.Content value="shortcuts" class="doc-panel">
					<h3 class="doc-panel-title">Keyboard Shortcuts</h3>
					<div class="shortcuts-grid-compact">
						{#each shortcuts as shortcut}
							<div class="shortcut-item-compact">
								<div class="shortcut-keys">
									{#each shortcut.keys as key}
										<kbd class="key">{key}</kbd>
									{/each}
								</div>
								<div class="shortcut-description">
									{shortcut.description}
								</div>
							</div>
						{/each}
					</div>
				</Tabs.Content>
			</Tabs.Root>
		</div>
	</section>

	<!-- Bug Report / Feature Request -->
	<section class="contribute-section">
		<div class="container">
			<div class="contribute-content">
				<h2 class="contribute-title">
					Found a bug or want to request a feature?
				</h2>
				<p class="contribute-description">
					We'd love to hear from you! Open an issue on GitHub to
					report bugs, request features, or contribute to the project.
				</p>
				<div class="contribute-buttons">
					<Button
						href="https://github.com/aritra1999/httpyum/issues/new?labels=bug&template=bug_report.md"
						variant="outline"
						size="lg"
						class="gap-2"
					>
						<Bug size={20} />
						<span>Report a Bug</span>
					</Button>
					<Button
						href="https://github.com/aritra1999/httpyum/issues/new?labels=enhancement&template=feature_request.md"
						variant="default"
						size="lg"
						class="gap-2"
					>
						<Lightbulb size={20} />
						<span>Request a Feature</span>
					</Button>
				</div>
			</div>
		</div>
	</section>

	<!-- Footer -->
	<footer class="footer">
		<div class="container">
			<div class="footer-content">
				<div class="footer-info">
					<strong>httpyum</strong> · Built with Go + Bubbletea · MIT License
				</div>
				<div class="footer-links">
					<a href="https://github.com/aritra1999/httpyum">
						<Github size={16} />
						<span>GitHub</span>
					</a>
					<a href="https://github.com/aritra1999/httpyum/releases">
						<Download size={16} />
						<span>Releases</span>
					</a>
				</div>
			</div>
		</div>
	</footer>
</div>

<style>
	:global(html) {
		scroll-behavior: smooth;
	}

	:global(body) {
		margin: 0;
		padding: 0;
		overflow-x: hidden;
	}

	.page {
		background: #fff;
		color: #000;
		font-family: "Space Grotesk", sans-serif;
		line-height: 1.5;
	}

	/* Button Styles - Brutalist in hero, subtle elsewhere */
	:global(.hero button, .hero a[role="button"]) {
		border-radius: 0 !important;
		font-weight: 700 !important;
		text-transform: lowercase !important;
		letter-spacing: 0.02em !important;
		transition: all 0.2s ease !important;
		font-family: "Space Grotesk", sans-serif !important;
		border: 4px solid #000 !important;
		box-shadow: 6px 6px 0 #000 !important;
	}

	:global(.hero button:hover, .hero a[role="button"]:hover) {
		transform: translate(3px, 3px) !important;
		box-shadow: 3px 3px 0 #000 !important;
	}

	:global(.hero [data-variant="default"]) {
		background: #000 !important;
		color: #fff !important;
	}

	:global(.hero [data-variant="outline"]) {
		background: #fff !important;
		color: #000 !important;
	}

	.tabs-wrapper {
		width: 100%;
		overflow: hidden;
		background: #f5f5f5;
		border: 1px solid #e5e5e5;
		border-radius: 8px;
		padding: 0.25rem;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
	}

	:global(.tabs-list) {
		overflow-x: auto;
		overflow-y: hidden;
		-webkit-overflow-scrolling: touch;
		scrollbar-width: thin;
		flex-wrap: nowrap !important;
		width: 100% !important;
		max-width: 100% !important;
		background: transparent !important;
		padding: 0 !important;
	}

	:global(.tabs-list::-webkit-scrollbar) {
		height: 3px;
	}

	:global(.tabs-list::-webkit-scrollbar-track) {
		background: #f5f5f5;
	}

	:global(.tabs-list::-webkit-scrollbar-thumb) {
		background: #999;
		border-radius: 3px;
	}

	:global(.tabs-trigger) {
		border-radius: 8px !important;
		border: none !important;
		background: transparent !important;
		font-weight: 500 !important;
		transition: all 0.2s !important;
		font-size: 0.875rem !important;
		padding: 2rem 1.75rem !important;
		white-space: nowrap !important;
		flex-shrink: 0 !important;
		min-width: fit-content !important;
		color: #666 !important;
	}

	:global(.tabs-trigger[data-state="active"]) {
		background: #fff !important;
		color: #000 !important;
		font-weight: 600 !important;
		box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05) !important;
	}

	:global(.tabs-trigger:hover:not([data-state="active"])) {
		background: rgba(255, 255, 255, 0.5) !important;
		color: #000 !important;
	}

	/* Navbar */
	.navbar {
		position: fixed;
		top: 1.5rem;
		left: 50%;
		transform: translateX(-50%);
		z-index: 1000;
		background: rgba(255, 255, 255, 0.8);
		backdrop-filter: blur(10px);
		-webkit-backdrop-filter: blur(10px);
		border-radius: 8px;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
		transition: all 0.3s ease;
		max-width: 1100px;
		width: calc(100% - 4rem);
	}

	.navbar-container {
		padding: 0.75rem 1.5rem;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.navbar-logo {
		display: flex;
		align-items: center;
		color: #000;
		text-decoration: none;
		transition: all 0.2s ease;
	}

	.navbar-logo:hover {
		transform: scale(1.05);
	}

	.navbar-logo svg {
		display: block;
	}

	.navbar-github {
		display: flex;
		align-items: center;
		justify-content: center;
		color: #000;
		text-decoration: none;
		padding: 0.5rem;
		border-radius: 8px;
		transition: all 0.2s ease;
	}

	.navbar-github:hover {
		background: rgba(0, 0, 0, 0.05);
	}

	/* Container */
	.container {
		max-width: 1100px;
		margin: 0 auto;
		padding: 0 2rem;
	}

	/* Hero */
	.hero {
		min-height: 100vh;
		display: flex;
		align-items: center;
		padding: 4rem 2rem;
		position: relative;
		overflow: hidden;
		background: #fff;
	}

	.hero::before {
		content: "";
		position: absolute;
		top: -50%;
		left: -50%;
		right: -50%;
		bottom: -50%;
		background-image:
			linear-gradient(to right, #e5e5e5 1px, transparent 1px),
			linear-gradient(to bottom, #e5e5e5 1px, transparent 1px);
		background-size: 80px 80px;
		transform: skew(-12deg);
		animation: moveBoxes 20s linear infinite;
		pointer-events: none;
		opacity: 0.4;
	}

	@keyframes moveBoxes {
		0% {
			transform: skew(-12deg) translateX(0) translateY(0);
		}
		100% {
			transform: skew(-12deg) translateX(80px) translateY(80px);
		}
	}

	.hero-grid {
		display: grid;
		grid-template-columns: 30% 70%;
		gap: 4rem;
		align-items: center;
		position: relative;
		z-index: 1;
	}

	.hero-content {
		max-width: 600px;
	}

	.hero-stars {
		display: inline;
		color: #666;
		font-size: 1rem;
	}

	.stars-link {
		display: inline-flex;
		align-items: center;
		gap: 0.25rem;
		color: #666;
		text-decoration: none;
		transition: color 0.2s ease;
	}

	.stars-link:hover {
		color: #000;
	}

	.hero-demo {
		position: relative;
	}

	.hero-title {
		font-size: clamp(2.5rem, 6vw, 3.5rem);
		font-weight: 700;
		margin: 0 0 1rem;
		letter-spacing: -0.02em;
		line-height: 0.95;
		color: #000;
		text-transform: lowercase;
	}

	.hero-subtitle {
		font-size: 1.25rem;
		color: #000;
		background: #fff;
		margin: 0 0 2rem;
		max-width: 600px;
		font-weight: 500;
	}

	.hero-buttons {
		display: flex;
		gap: 0.75rem;
	}

	/* Section Title */
	.section-title {
		font-size: 2.5rem;
		font-weight: 700;
		margin: 0 0 2.5rem;
		letter-spacing: -0.02em;
		text-transform: lowercase;
		color: #000;
		position: relative;
		display: inline-block;
	}

	.section-title::after {
		content: "";
		position: absolute;
		bottom: -8px;
		left: 0;
		width: 50%;
		height: 4px;
		background: #000;
	}

	.demo-wrapper {
		background: #000;
		border: 2px solid #000;
		overflow: hidden;
		box-shadow: 4px 4px 0 #000;
		transform: rotate(-1deg);
		transition: transform 0.3s ease;
	}

	.demo-wrapper:hover {
		transform: rotate(0deg);
	}

	.demo-video {
		width: 100%;
		height: auto;
		display: block;
	}

	.scroll-indicator {
		position: absolute;
		bottom: 2rem;
		left: 50%;
		transform: translateX(-50%);
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.5rem;
		text-decoration: none;
		z-index: 10;
		animation: bounce 2s ease-in-out infinite;
	}

	.scroll-indicator-line {
		width: 2px;
		height: 40px;
		background: #000;
	}

	.scroll-indicator-dot {
		width: 8px;
		height: 8px;
		background: #000;
		border-radius: 50%;
	}

	@keyframes bounce {
		0%,
		100% {
			transform: translateX(-50%) translateY(0);
		}
		50% {
			transform: translateX(-50%) translateY(10px);
		}
	}

	.install-code {
		font-family: "JetBrains Mono", monospace;
		font-size: 0.875rem;
		line-height: 1.6;
		color: #333;
		margin: 0;
		overflow-x: auto;
		white-space: pre-wrap;
		word-break: break-all;
	}

	:global(.copy-btn) {
		position: absolute;
		top: 1rem;
		right: 1rem;
		background: #fff !important;
		color: #000 !important;
		border: 1px solid #e5e5e5 !important;
		border-radius: 4px !important;
	}

	:global(.copy-btn:hover) {
		background: #f5f5f5 !important;
		border-color: #999 !important;
	}

	/* Features */
	.features-section {
		padding: 4rem 2rem;
		background: #fff;
		border-bottom: 1px solid #e5e5e5;
	}

	.features-list {
		display: flex;
		flex-direction: column;
		gap: 0;
		max-width: 900px;
		margin: 0 auto;
	}

	.feature-item {
		display: grid;
		grid-template-columns: 60px 40px 1fr;
		gap: 1.5rem;
		align-items: start;
		padding: 2rem 0;
		border-bottom: 1px solid #e5e5e5;
		transition: all 0.2s ease;
	}

	.feature-item:last-child {
		border-bottom: none;
	}

	.feature-item:hover {
		padding-left: 1rem;
		background: #fafafa;
		margin: 0 -1rem;
		padding-left: 2rem;
		padding-right: 1rem;
	}

	.feature-number {
		font-size: 2rem;
		font-weight: 700;
		color: #e5e5e5;
		line-height: 1;
		font-family: "Space Grotesk", sans-serif;
	}

	.feature-icon-minimal {
		width: 40px;
		height: 40px;
		display: flex;
		align-items: center;
		justify-content: center;
		background: #f5f5f5;
		border-radius: 8px;
		flex-shrink: 0;
	}

	:global(.feature-icon-minimal svg) {
		stroke: #000;
		stroke-width: 2;
	}

	.feature-content {
		padding-top: 0.25rem;
	}

	.feature-title {
		font-size: 1.25rem;
		font-weight: 700;
		margin: 0 0 0.5rem;
		color: #000;
	}

	.feature-description {
		font-size: 1rem;
		color: #666;
		margin: 0;
		line-height: 1.6;
	}

	/* Example */
	.shortcuts-grid-compact {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: 1.5rem;
	}

	.shortcut-item-compact {
		display: flex;
		align-items: center;
		gap: 1rem;
		padding: 1rem;
		background: #fff;
		border: 1px solid #e5e5e5;
		border-radius: 6px;
		transition: all 0.2s;
	}

	.shortcut-keys {
		display: flex;
		gap: 0.5rem;
	}

	.key {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		min-width: 28px;
		padding: 0.25rem 0.5rem;
		background: #f5f5f5;
		color: #333;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		font-family: "JetBrains Mono", monospace;
		font-size: 0.8125rem;
		font-weight: 600;
		font-style: normal;
	}

	.shortcut-description {
		font-size: 0.9375rem;
		color: #666;
	}

	/* Documentation */
	.docs-section {
		padding: 4rem 2rem;
		background: #fff;
		border-bottom: 1px solid #e5e5e5;
	}

	:global(.doc-panel) {
		background: #f5f5f5 !important;
		border: 1px solid #e5e5e5 !important;
		border-radius: 8px !important;
		padding: 2rem !important;
		animation: fadeIn 0.3s ease-in-out;
	}

	.doc-panel-title {
		font-size: 1.5rem;
		font-weight: 700;
		margin: 0 0 1.5rem;
	}

	.doc-section-title {
		font-size: 1.125rem;
		font-weight: 700;
		margin: 2rem 0 1rem;
		color: #000;
	}

	.doc-section-title:first-of-type {
		margin-top: 0;
	}

	@keyframes fadeIn {
		from {
			opacity: 0;
			transform: translateY(10px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.doc-content p {
		margin: 0 0 1rem;
		color: #666;
		font-size: 0.9375rem;
	}

	.doc-code {
		font-family: "JetBrains Mono", monospace;
		font-size: 0.875rem;
		line-height: 1.6;
		background: #fff;
		color: #333;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		padding: 1rem;
		margin: 1rem 0;
		overflow-x: auto;
		display: block;
	}

	.doc-list {
		margin: 1rem 0;
		padding-left: 1.5rem;
		color: #666;
		font-size: 0.9375rem;
	}

	.doc-list li {
		margin: 0.5rem 0;
	}

	.doc-list code,
	.doc-content code {
		font-family: "JetBrains Mono", monospace;
		font-size: 0.875em;
		background: #fff;
		color: #333;
		padding: 0.2rem 0.4rem;
		border-radius: 3px;
		border: 1px solid #e5e5e5;
	}

	.doc-key {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		min-width: 28px;
		padding: 0.3rem 0.5rem;
		background: #f5f5f5;
		color: #333;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		font-family: "JetBrains Mono", monospace;
		font-size: 0.8125rem;
		font-weight: 600;
		font-style: normal;
	}

	.install-card-inline {
		background: #fff;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		padding: 1.5rem;
		margin: 1rem 0;
		position: relative;
	}

	.install-platforms-inline {
		display: flex;
		gap: 0.75rem;
		margin: 1rem 0 1.5rem;
	}

	.platform-inline {
		padding: 0.4rem 0.8rem;
		background: #f5f5f5;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		font-size: 0.8125rem;
		font-weight: 500;
		color: #666;
	}

	/* Contribute */
	.contribute-section {
		padding: 4rem 2rem;
		background: #f5f5f5;
		border-bottom: 1px solid #e5e5e5;
	}

	.contribute-content {
		text-align: center;
		max-width: 700px;
		margin: 0 auto;
	}

	.contribute-title {
		font-size: 1.75rem;
		font-weight: 700;
		margin: 0 0 1rem;
		letter-spacing: -0.01em;
	}

	.contribute-description {
		font-size: 1.125rem;
		color: #666;
		margin: 0 0 2rem;
		line-height: 1.6;
	}

	.contribute-buttons {
		display: flex;
		gap: 1.5rem;
		justify-content: center;
		flex-wrap: wrap;
	}

	/* Footer */
	.footer {
		padding: 3rem 2rem;
		background: #000;
		color: #fff;
	}

	.footer-content {
		display: flex;
		justify-content: space-between;
		align-items: center;
		flex-wrap: wrap;
		gap: 1.5rem;
	}

	.footer-info {
		font-size: 0.875rem;
		color: #fff;
		font-weight: 400;
	}

	.footer-links {
		display: flex;
		gap: 1.5rem;
	}

	.footer-links a {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		color: #fff;
		text-decoration: none;
		font-size: 0.875rem;
		font-weight: 400;
		transition: all 0.2s;
		border-bottom: 2px solid transparent;
	}

	.footer-links a:hover {
		border-bottom-color: #fff;
	}

	@media (max-width: 768px) {
		/* Navbar */
		.navbar {
			top: 1rem;
			width: calc(100% - 2rem);
			max-width: none;
		}

		.navbar-container {
			padding: 0.5rem 1rem;
			gap: 1.5rem;
		}

		.navbar-logo svg {
			width: 28px;
			height: 28px;
		}

		.navbar-github {
			padding: 0.375rem;
		}

		/* Hero Section */
		.hero {
			min-height: 100vh;
			padding: 3rem 1rem;
		}

		.hero-grid {
			grid-template-columns: 1fr;
			gap: 2rem;
		}

		.hero-title {
			font-size: 2.5rem;
		}

		.hero-subtitle {
			font-size: 1rem;
		}

		:global(.hero button, .hero a[role="button"]) {
			width: 100%;
			justify-content: center;
		}

		.demo-wrapper {
			transform: rotate(0deg);
			box-shadow: 4px 4px 0 #000;
		}

		.scroll-indicator {
			bottom: 1.5rem;
		}

		.scroll-indicator-line {
			height: 30px;
		}

		/* Container padding */
		.container {
			padding: 0 1rem;
		}

		/* Section titles */
		.section-title {
			font-size: 1.75rem;
			margin-bottom: 1.5rem;
		}

		.section-title::after {
			width: 40%;
		}

		/* Features */
		.features-section,
		.docs-section,
		.contribute-section {
			padding: 3rem 1rem;
		}

		.feature-item {
			grid-template-columns: 50px 36px 1fr;
			gap: 1rem;
			padding: 1.5rem 0;
		}

		.feature-item:hover {
			margin: 0;
			padding-left: 0;
			padding-right: 0;
		}

		.feature-number {
			font-size: 1.5rem;
		}

		.feature-icon-minimal {
			width: 36px;
			height: 36px;
		}

		.feature-title {
			font-size: 1.125rem;
		}

		.feature-description {
			font-size: 0.9375rem;
		}

		/* Shortcuts grid */
		.shortcuts-grid-compact {
			grid-template-columns: 1fr;
		}

		/* Install platforms */
		.install-platforms-inline {
			flex-wrap: wrap;
		}

		/* Contribute buttons */
		.contribute-buttons {
			width: 100%;
		}

		.contribute-buttons :global(button),
		.contribute-buttons :global(a[role="button"]) {
			width: 100%;
		}

		.contribute-title {
			font-size: 1.5rem;
		}

		.contribute-description {
			font-size: 1rem;
		}

		/* Footer */
		.footer {
			padding: 2rem 1rem;
		}

		.footer-content {
			flex-direction: column;
			text-align: center;
			gap: 1rem;
		}

		.footer-links {
			justify-content: center;
		}

		/* Doc section */
		.doc-panel-title {
			font-size: 1.25rem;
		}

		.doc-section-title {
			font-size: 1rem;
			margin: 1.5rem 0 0.75rem;
		}

		/* Tabs - ensure proper scrolling on mobile */
		.tabs-wrapper {
			margin: 0;
			padding: 0.25rem;
			overflow-x: auto;
			-webkit-overflow-scrolling: touch;
		}

		:global(.tabs-list) {
			justify-content: flex-start !important;
			padding: 0 !important;
			width: auto !important;
			min-width: 100%;
		}

		:global(.tabs-trigger) {
			font-size: 0.8125rem !important;
			padding: 0.75rem 1rem !important;
		}

		/* Code blocks */
		.install-code,
		.doc-code {
			font-size: 0.8125rem;
			padding: 0.75rem;
		}

		:global(.copy-btn) {
			top: 0.75rem;
			right: 0.75rem;
		}
	}

	@media (max-width: 480px) {
		.hero-title {
			font-size: 2rem;
		}

		.hero-subtitle {
			font-size: 1rem;
		}

		.section-title {
			font-size: 1.5rem;
		}

		.feature-title {
			font-size: 1rem;
		}

		.contribute-title {
			font-size: 1.25rem;
		}

		:global(.tabs-trigger) {
			font-size: 0.75rem !important;
			padding: 0.5rem 0.6rem !important;
		}
	}
</style>
