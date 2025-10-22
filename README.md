# Prompt Engineering Agent (PEA)

This is a local-first Prompt Engineering Agent (PEA) that is Windows-first and uses the Gemini CLI as the single AI provider.

## Quickstart (Windows)

1.  Install the [Gemini CLI](https://ai.google.dev/docs/gemini_cli_quickstart).
2.  Authenticate with the Gemini CLI: `gemini auth login`.
3.  Run the PEA server: `go run cmd/pea/main.go serve`.
4.  Run the Flutter client: `flutter run -d windows`.

## Authentication

Authentication is handled exclusively by the user's local Gemini CLI OAuth flow. The server will only attempt to use the real Gemini CLI when `PEA_GEMINI_MOCK` is not set to `true` and the CLI is authenticated. You can check your authentication status by running `pea check-auth`.
