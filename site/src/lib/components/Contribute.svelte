<script lang="ts">
	import { Bug, Lightbulb } from "@lucide/svelte";
	import { Button } from "$lib/components/ui/button/index.js";
</script>

<section
	class="py-16 px-8 bg-[#f5f5f5] border-b border-[#e5e5e5] max-md:py-12 max-md:px-4"
>
	<div class="container max-w-[1100px] mx-auto px-8 max-md:px-4">
		<div class="text-center max-w-[700px] mx-auto">
			<h2
				class="text-[1.75rem] font-bold m-0 mb-4 -tracking-[0.01em] max-md:text-2xl max-[480px]:text-xl"
			>
				Found a bug or want to request a feature?
			</h2>
			<p
				class="text-lg text-[#666] m-0 mb-8 leading-relaxed max-md:text-base"
			>
				We'd love to hear from you! Open an issue on GitHub to report
				bugs, request features, or contribute to the project.
			</p>
			<div
				class="flex gap-6 justify-center flex-wrap max-md:w-full max-md:[&_button]:w-full max-md:[&_a[role='button']]:w-full"
			>
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
