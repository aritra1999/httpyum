<script lang="ts">
	import type { Component } from "svelte";

	interface Feature {
		icon: Component<any>;
		title: string;
		description: string;
	}

	interface Props {
		features: Feature[];
	}

	let { features }: Props = $props();
</script>

<section
	class="py-16 px-8 bg-white border-b border-[#e5e5e5] max-md:py-12 max-md:px-4"
	id="features"
>
	<div class="container max-w-[1100px] mx-auto px-8 max-md:px-4">
		<h2
			class="text-[2.5rem] font-bold m-0 mb-10 -tracking-[0.02em] lowercase text-black relative inline-block after:content-[''] after:absolute after:-bottom-2 after:left-0 after:w-1/2 after:h-1 after:bg-black max-md:text-[1.75rem] max-md:mb-6 max-md:after:w-2/5 max-[480px]:text-2xl"
		>
			Features
		</h2>
		<div class="flex flex-col gap-0 max-w-[900px] mx-auto">
			{#each features as feature, i}
				{@const Icon = feature.icon}
				<div
					class="grid grid-cols-[60px_40px_1fr] gap-6 items-start py-8 border-b border-[#e5e5e5] transition-all duration-200 last:border-b-0 hover:pl-4 hover:bg-[#fafafa] hover:mx-[-1rem] hover:pl-8 hover:pr-4 max-md:grid-cols-[50px_36px_1fr] max-md:gap-4 max-md:py-6 max-md:hover:m-0 max-md:hover:pl-0 max-md:hover:pr-0"
				>
					<div
						class="text-[2rem] font-bold text-[#e5e5e5] leading-none font-['Space_Grotesk',sans-serif] max-md:text-2xl"
					>
						{String(i + 1).padStart(2, "0")}
					</div>
					<div
						class="w-10 h-10 flex items-center justify-center bg-[#f5f5f5] rounded-lg flex-shrink-0 max-md:w-9 max-md:h-9"
					>
						<Icon
							size={20}
							strokeWidth={2}
							class="stroke-black stroke-2"
						/>
					</div>
					<div class="pt-1">
						<h3
							class="text-xl font-bold m-0 mb-2 text-black max-md:text-lg max-[480px]:text-base"
						>
							{feature.title}
						</h3>
						<p
							class="text-base text-[#666] m-0 leading-relaxed max-md:text-[0.9375rem]"
						>
							{feature.description}
						</p>
					</div>
				</div>
			{/each}
		</div>
	</div>
</section>
