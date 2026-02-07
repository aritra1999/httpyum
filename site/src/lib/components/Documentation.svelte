<script lang="ts">
	import { Copy, Check } from "@lucide/svelte";
	import { Button } from "$lib/components/ui/button/index.js";
	import * as Tabs from "$lib/components/ui/tabs/index.js";

	interface Shortcut {
		keys: string[];
		description: string;
	}

	interface Props {
		shortcuts: Shortcut[];
		copied?: boolean;
		onCopy?: () => void;
	}

	let { shortcuts, copied = false, onCopy }: Props = $props();
</script>

<section class="docs-section" id="docs">
	<div class="container">
		<h2 class="section-title">Documentation</h2>

		<Tabs.Root value="getting-started" class="docs-interactive">
			<div class="tabs-wrapper">
				<Tabs.List class="docs-tabs">
					<Tabs.Trigger value="getting-started" class="doc-tab">
						Getting Started
					</Tabs.Trigger>
					<Tabs.Trigger value="file-format" class="doc-tab">
						File Format
					</Tabs.Trigger>
					<Tabs.Trigger value="variables" class="doc-tab">
						Variables
					</Tabs.Trigger>
					<Tabs.Trigger value="json-viewer" class="doc-tab">
						JSON Viewer
					</Tabs.Trigger>
					<Tabs.Trigger value="shortcuts" class="doc-tab">
						Keyboard Shortcuts
					</Tabs.Trigger>
				</Tabs.List>
			</div>

			<Tabs.Content value="getting-started" class="doc-panel">
				<h3 class="doc-panel-title">Getting Started</h3>
				<div class="doc-content">
					<h4 class="doc-section-title">Installation</h4>
					<p>Install httpyum using the installation script:</p>
					<div class="install-card-inline">
						<pre
							class="install-code">curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/install.sh | bash</pre>
						<Button
							variant="ghost"
							size="sm"
							class="copy-btn"
							onclick={onCopy}
						>
							{#if copied}
								<Check size={16} />
							{:else}
								<Copy size={16} />
							{/if}
						</Button>
					</div>
					<div class="install-platforms-inline">
						<span class="platform-inline">macOS</span>
						<span class="platform-inline">Linux</span>
						<span class="platform-inline">Windows</span>
					</div>

					<h4 class="doc-section-title">Usage</h4>
					<p>Run httpyum with any .http file:</p>
					<pre class="doc-code">httpyum example.http</pre>
					<p>Command-line options:</p>
					<ul class="doc-list">
						<li>
							<code>--no-headers</code> - Hide response headers
						</li>
						<li>
							<code>-h, --help</code> - Show help message
						</li>
						<li>
							<code>-v, --version</code> - Show version
						</li>
					</ul>
				</div>
			</Tabs.Content>

			<Tabs.Content value="file-format" class="doc-panel">
				<h3 class="doc-panel-title">File Format</h3>
				<div class="doc-content">
					<p>Standard .http file format with variables:</p>
					<pre class="doc-code">@variable = value

### Request Name
METHOD url
Header: value

body</pre>
				</div>
			</Tabs.Content>

			<Tabs.Content value="variables" class="doc-panel">
				<h3 class="doc-panel-title">Variables</h3>
				<div class="doc-content">
					<p>
						Define variables with <code>@</code> and use with
						<code>{"{{"}{"}}"}</code>:
					</p>
					<pre class="doc-code">@baseUrl = https://api.example.com
GET {"{{"}baseUrl{"}}"}/users</pre>
					<p>Load from environment:</p>
					<pre
						class="doc-code">@token = {"{{"}$dotenv JWT{"}}"}</pre>
				</div>
			</Tabs.Content>

			<Tabs.Content value="json-viewer" class="doc-panel">
				<h3 class="doc-panel-title">JSON Viewer</h3>
				<div class="doc-content">
					<p>
						Press <kbd class="doc-key">f</kbd> to open JSON responses
						in jless for interactive exploration.
					</p>
					<p>Install jless for the best experience:</p>
					<pre class="doc-code"># macOS
brew install jless

# Linux
cargo install jless</pre>
				</div>
			</Tabs.Content>

			<Tabs.Content value="shortcuts" class="doc-panel">
				<h3 class="doc-panel-title">Keyboard Shortcuts</h3>
				<div class="shortcuts-grid-compact">
					{#each shortcuts as shortcut}
						<div class="shortcut-item-compact">
							<div class="shortcut-keys">
								{#each shortcut.keys as key}
									<kbd class="key">{key}</kbd>
								{/each}
							</div>
							<div class="shortcut-description">
								{shortcut.description}
							</div>
						</div>
					{/each}
				</div>
			</Tabs.Content>
		</Tabs.Root>
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

	.tabs-wrapper {
		width: 100%;
		overflow: hidden;
		background: #f5f5f5;
		border: 1px solid #e5e5e5;
		border-radius: 8px;
		padding: 0.25rem;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
	}

	:global(.tabs-list) {
		overflow-x: auto;
		overflow-y: hidden;
		-webkit-overflow-scrolling: touch;
		scrollbar-width: thin;
		flex-wrap: nowrap !important;
		width: 100% !important;
		max-width: 100% !important;
		background: transparent !important;
		padding: 0 !important;
	}

	:global(.tabs-list::-webkit-scrollbar) {
		height: 3px;
	}

	:global(.tabs-list::-webkit-scrollbar-track) {
		background: #f5f5f5;
	}

	:global(.tabs-list::-webkit-scrollbar-thumb) {
		background: #999;
		border-radius: 3px;
	}

	:global(.tabs-trigger) {
		border-radius: 8px !important;
		border: none !important;
		background: transparent !important;
		font-weight: 500 !important;
		transition: all 0.2s !important;
		font-size: 0.875rem !important;
		padding: 2rem 1.75rem !important;
		white-space: nowrap !important;
		flex-shrink: 0 !important;
		min-width: fit-content !important;
		color: #666 !important;
	}

	:global(.tabs-trigger[data-state="active"]) {
		background: #fff !important;
		color: #000 !important;
		font-weight: 600 !important;
		box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05) !important;
	}

	:global(.tabs-trigger:hover:not([data-state="active"])) {
		background: rgba(255, 255, 255, 0.5) !important;
		color: #000 !important;
	}

	.docs-section {
		padding: 4rem 2rem;
		background: #fff;
		border-bottom: 1px solid #e5e5e5;
	}

	:global(.doc-panel) {
		background: #f5f5f5 !important;
		border: 1px solid #e5e5e5 !important;
		border-radius: 8px !important;
		padding: 2rem !important;
		animation: fadeIn 0.3s ease-in-out;
	}

	.doc-panel-title {
		font-size: 1.5rem;
		font-weight: 700;
		margin: 0 0 1.5rem;
	}

	.doc-section-title {
		font-size: 1.125rem;
		font-weight: 700;
		margin: 2rem 0 1rem;
		color: #000;
	}

	.doc-section-title:first-of-type {
		margin-top: 0;
	}

	@keyframes fadeIn {
		from {
			opacity: 0;
			transform: translateY(10px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.doc-content p {
		margin: 0 0 1rem;
		color: #666;
		font-size: 0.9375rem;
	}

	.doc-code {
		font-family: "JetBrains Mono", monospace;
		font-size: 0.875rem;
		line-height: 1.6;
		background: #fff;
		color: #333;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		padding: 1rem;
		margin: 1rem 0;
		overflow-x: auto;
		display: block;
	}

	.doc-list {
		margin: 1rem 0;
		padding-left: 1.5rem;
		color: #666;
		font-size: 0.9375rem;
	}

	.doc-list li {
		margin: 0.5rem 0;
	}

	.doc-list code,
	.doc-content code {
		font-family: "JetBrains Mono", monospace;
		font-size: 0.875em;
		background: #fff;
		color: #333;
		padding: 0.2rem 0.4rem;
		border-radius: 3px;
		border: 1px solid #e5e5e5;
	}

	.doc-key {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		min-width: 28px;
		padding: 0.3rem 0.5rem;
		background: #f5f5f5;
		color: #333;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		font-family: "JetBrains Mono", monospace;
		font-size: 0.8125rem;
		font-weight: 600;
		font-style: normal;
	}

	.install-card-inline {
		background: #fff;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		padding: 1.5rem;
		margin: 1rem 0;
		position: relative;
	}

	.install-code {
		font-family: "JetBrains Mono", monospace;
		font-size: 0.875rem;
		line-height: 1.6;
		color: #333;
		margin: 0;
		overflow-x: auto;
		white-space: pre-wrap;
		word-break: break-all;
	}

	:global(.copy-btn) {
		position: absolute;
		top: 1rem;
		right: 1rem;
		background: #fff !important;
		color: #000 !important;
		border: 1px solid #e5e5e5 !important;
		border-radius: 4px !important;
	}

	:global(.copy-btn:hover) {
		background: #f5f5f5 !important;
		border-color: #999 !important;
	}

	.install-platforms-inline {
		display: flex;
		gap: 0.75rem;
		margin: 1rem 0 1.5rem;
	}

	.platform-inline {
		padding: 0.4rem 0.8rem;
		background: #f5f5f5;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		font-size: 0.8125rem;
		font-weight: 500;
		color: #666;
	}

	.shortcuts-grid-compact {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: 1.5rem;
	}

	.shortcut-item-compact {
		display: flex;
		align-items: center;
		gap: 1rem;
		padding: 1rem;
		background: #fff;
		border: 1px solid #e5e5e5;
		border-radius: 6px;
		transition: all 0.2s;
	}

	.shortcut-keys {
		display: flex;
		gap: 0.5rem;
	}

	.key {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		min-width: 28px;
		padding: 0.25rem 0.5rem;
		background: #f5f5f5;
		color: #333;
		border: 1px solid #e5e5e5;
		border-radius: 4px;
		font-family: "JetBrains Mono", monospace;
		font-size: 0.8125rem;
		font-weight: 600;
		font-style: normal;
	}

	.shortcut-description {
		font-size: 0.9375rem;
		color: #666;
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

		.docs-section {
			padding: 3rem 1rem;
		}

		.shortcuts-grid-compact {
			grid-template-columns: 1fr;
		}

		.install-platforms-inline {
			flex-wrap: wrap;
		}

		.doc-panel-title {
			font-size: 1.25rem;
		}

		.doc-section-title {
			font-size: 1rem;
			margin: 1.5rem 0 0.75rem;
		}

		.tabs-wrapper {
			margin: 0;
			padding: 0.25rem;
			overflow-x: auto;
			-webkit-overflow-scrolling: touch;
		}

		:global(.tabs-list) {
			justify-content: flex-start !important;
			padding: 0 !important;
			width: auto !important;
			min-width: 100%;
		}

		:global(.tabs-trigger) {
			font-size: 0.8125rem !important;
			padding: 0.75rem 1rem !important;
		}

		.install-code,
		.doc-code {
			font-size: 0.8125rem;
			padding: 0.75rem;
		}

		:global(.copy-btn) {
			top: 0.75rem;
			right: 0.75rem;
		}
	}

	@media (max-width: 480px) {
		.section-title {
			font-size: 1.5rem;
		}

		:global(.tabs-trigger) {
			font-size: 0.75rem !important;
			padding: 0.5rem 0.6rem !important;
		}
	}
</style>
