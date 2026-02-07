<script lang="ts">
	import { Github } from "@lucide/svelte";
</script>

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

<style>
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

	@media (max-width: 768px) {
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
	}
</style>
