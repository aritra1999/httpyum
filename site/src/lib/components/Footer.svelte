<script lang="ts">
	import { Github, Download } from "@lucide/svelte";
</script>

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

<style>
	.container {
		max-width: 1100px;
		margin: 0 auto;
		padding: 0 2rem;
	}

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
		.container {
			padding: 0 1rem;
		}

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
	}
</style>
