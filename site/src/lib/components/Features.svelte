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

<style>
	.container {
		max-width: 1100px;
		margin: 0 auto;
		padding: 0 2rem;
	}

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

	@media (max-width: 768px) {
		.container {
			padding: 0 1rem;
		}

		.section-title {
			font-size: 1.75rem;
			margin-bottom: 1.5rem;
		}

		.section-title::after {
			width: 40%;
		}

		.features-section {
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
	}

	@media (max-width: 480px) {
		.section-title {
			font-size: 1.5rem;
		}

		.feature-title {
			font-size: 1rem;
		}
	}
</style>
