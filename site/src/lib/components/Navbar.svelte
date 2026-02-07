<script lang="ts">
	import { Github } from "@lucide/svelte";
	import { base } from "$app/paths";
</script>

<nav
	class="fixed top-6 left-1/2 -translate-x-1/2 z-[1000] bg-white/80 backdrop-blur-md rounded-lg shadow-[0_4px_12px_rgba(0,0,0,0.05)] transition-all duration-300 max-w-[1100px] w-[calc(100%-4rem)] md:top-6 md:w-[calc(100%-4rem)] max-md:top-4 max-md:w-[calc(100%-2rem)]"
>
	<div
		class="px-6 py-3 flex justify-between items-center max-md:px-4 max-md:py-2 gap-6"
	>
		<a
			href="{base}/"
			class="flex items-center text-black no-underline transition-transform duration-200 hover:scale-105"
		>
			<svg
				width="32"
				height="32"
				viewBox="0 0 32 32"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				class="block max-md:w-7 max-md:h-7"
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
			class="flex items-center justify-center text-black no-underline p-2 rounded-lg transition-all duration-200 hover:bg-black/5 max-md:p-1.5"
		>
			<Github size={20} />
		</a>
	</div>
</nav>
