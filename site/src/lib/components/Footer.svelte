<script lang="ts">
	import { Github, Download } from "@lucide/svelte";
</script>

<footer class="py-12 px-8 bg-black text-white max-md:py-8 max-md:px-4">
	<div class="container max-w-[1100px] mx-auto px-8 max-md:px-4">
		<div
			class="flex justify-between items-center flex-wrap gap-6 max-md:flex-col max-md:text-center max-md:gap-4"
		>
			<div class="text-sm text-white font-normal">
				<strong>httpyum</strong> · Built with Go + Bubbletea · MIT License
			</div>
			<div class="flex gap-6 max-md:justify-center">
				<a
					href="https://github.com/aritra1999/httpyum"
					class="inline-flex items-center gap-2 text-white no-underline text-sm font-normal transition-all duration-200 border-b-2 border-transparent hover:border-white"
				>
					<Github size={16} />
					<span>GitHub</span>
				</a>
				<a
					href="https://github.com/aritra1999/httpyum/releases"
					class="inline-flex items-center gap-2 text-white no-underline text-sm font-normal transition-all duration-200 border-b-2 border-transparent hover:border-white"
				>
					<Download size={16} />
					<span>Releases</span>
				</a>
			</div>
		</div>
	</div>
</footer>
