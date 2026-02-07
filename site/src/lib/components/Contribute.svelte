<script lang="ts">
	import { Bug, Lightbulb } from "@lucide/svelte";
	import { Button } from "$lib/components/ui/button/index.js";
</script>

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

<style>
	.container {
		max-width: 1100px;
		margin: 0 auto;
		padding: 0 2rem;
	}

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

	@media (max-width: 768px) {
		.container {
			padding: 0 1rem;
		}

		.contribute-section {
			padding: 3rem 1rem;
		}

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
	}

	@media (max-width: 480px) {
		.contribute-title {
			font-size: 1.25rem;
		}
	}
</style>
