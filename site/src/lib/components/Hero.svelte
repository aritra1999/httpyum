<script lang="ts">
	import { Github, BookOpen, Star } from "@lucide/svelte";
	import { Button } from "$lib/components/ui/button/index.js";

	interface Props {
		stars?: number | null;
	}

	let { stars = null }: Props = $props();
</script>

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

<style>
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

	.container {
		max-width: 1100px;
		margin: 0 auto;
		padding: 0 2rem;
	}

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

	@media (max-width: 768px) {
		.container {
			padding: 0 1rem;
		}

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
	}

	@media (max-width: 480px) {
		.hero-title {
			font-size: 2rem;
		}

		.hero-subtitle {
			font-size: 1rem;
		}
	}
</style>
