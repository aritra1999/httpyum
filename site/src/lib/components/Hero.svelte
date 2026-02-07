<script lang="ts">
	import { Github, BookOpen, Star } from "@lucide/svelte";
	import { Button } from "$lib/components/ui/button/index.js";

	interface Props {
		stars?: number | null;
	}

	let { stars = null }: Props = $props();
</script>

<section
	class="min-h-screen flex items-center px-8 py-16 relative overflow-hidden bg-white max-md:px-4 max-md:py-12"
>
	<div
		class="absolute -top-1/2 -left-1/2 -right-1/2 -bottom-1/2 bg-[linear-gradient(to_right,#e5e5e5_1px,transparent_1px),linear-gradient(to_bottom,#e5e5e5_1px,transparent_1px)] bg-[length:80px_80px] -skew-y-12 animate-[moveBoxes_20s_linear_infinite] pointer-events-none opacity-40"
	></div>

	<div
		class="container max-w-[1100px] mx-auto px-8 grid grid-cols-[30%_70%] gap-16 items-center relative z-[1] max-md:px-4 max-md:grid-cols-1 max-md:gap-8"
	>
		<div class="max-w-[600px]">
			<h1
				class="text-[clamp(2.5rem,6vw,3.5rem)] font-bold m-0 mb-4 -tracking-[0.02em] leading-[0.95] text-black lowercase max-md:text-[2.5rem] max-[480px]:text-[2rem]"
			>
				httpyum
			</h1>
			<p
				class="text-xl text-black bg-white m-0 mb-8 max-w-[600px] font-medium max-md:text-base"
			>
				Fast, interactive CLI tool for executing HTTP requests from
				.http files
				{#if stars !== null}
					<span class="inline text-[#666] text-base">
						· <a
							href="https://github.com/aritra1999/httpyum/stargazers"
							target="_blank"
							rel="noopener noreferrer"
							class="inline-flex items-center gap-1 text-[#666] no-underline transition-colors duration-200 hover:text-black"
						>
							<Star size={14} fill="currentColor" />
							{stars.toLocaleString()}
						</a>
					</span>
				{/if}
			</p>
			<div class="flex gap-3 hero-buttons">
				<Button
					href="https://github.com/aritra1999/httpyum"
					variant="default"
					size="lg"
					class="gap-2"
				>
					<Github size={20} />
					<span>View on GitHub</span>
				</Button>
				<Button href="#docs" variant="outline" size="lg" class="gap-2">
					<BookOpen size={20} />
					<span>Docs</span>
				</Button>
			</div>
		</div>
		<div class="relative">
			<div
				class="border-2 border-black overflow-hidden shadow-[4px_4px_0_#000] -rotate-1 transition-transform duration-300 hover:rotate-0 max-md:rotate-0 max-md:shadow-[4px_4px_0_#000] bg-gray-100"
			>
				<img
					src="https://github.com/user-attachments/assets/e4313241-a4c8-4c80-a422-1fde7f953bcc"
					alt="httpyum demo"
					class="w-full h-auto block"
				/>
			</div>
		</div>
	</div>
	<a
		href="#features"
		class="absolute bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center gap-2 no-underline z-10 animate-[bounce_2s_ease-in-out_infinite] max-md:bottom-6"
	>
		<div class="w-0.5 h-10 bg-black max-md:h-[30px]"></div>
		<div class="w-2 h-2 bg-black rounded-full"></div>
	</a>
</section>

<style>
	:global(.hero-buttons button, .hero-buttons a[role="button"]) {
		border-radius: 0 !important;
		font-weight: 700 !important;
		text-transform: lowercase !important;
		letter-spacing: 0.02em !important;
		transition: all 0.2s ease !important;
		font-family: "Space Grotesk", sans-serif !important;
		border: 4px solid #000 !important;
		box-shadow: 6px 6px 0 #000 !important;
	}

	:global(.hero-buttons button:hover, .hero-buttons a[role="button"]:hover) {
		transform: translate(3px, 3px) !important;
		box-shadow: 3px 3px 0 #000 !important;
	}

	:global(.hero-buttons [data-variant="default"]) {
		background: #000 !important;
		color: #fff !important;
	}

	:global(.hero-buttons [data-variant="outline"]) {
		background: #fff !important;
		color: #000 !important;
	}

	@keyframes moveBoxes {
		0% {
			transform: skew(-12deg) translateX(0) translateY(0);
		}
		100% {
			transform: skew(-12deg) translateX(80px) translateY(80px);
		}
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
		:global(.hero-buttons button, .hero-buttons a[role="button"]) {
			width: 100%;
			justify-content: center;
		}
	}
</style>
