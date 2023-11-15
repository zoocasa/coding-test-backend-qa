# Instructions

Format: a 45 min pair programming session followed by 15 minutes of discussion and feedback.

Goal: We'll be creating unit tests for a golang http service together.

Preparation: Candidate should have a golang 1.18 development environment set up on their laptop and be prepared to screen share as we work together.

Notes & Tips: We're simulating a regular work day so the candidate is allowed to Google anything during the test (we don't expect you to have syntax memorized). The most successful candidates share their thoughts as they work, what ideas they have and what they're trying to do. If your IDE supports it, using a split editor/panel with server.go on the left and server_test.go on the right makes it easier to write tests while developing your code.

# Getting started
1. Download this git repository to candidate's machine
2. Verify environment: use "go run ." to verify your http server works and go to http://localhost:8080 (you can kill the process once verified)
3. Running tests: use "go test -v ./..."

Data model:
- A listing is a house for sale
- each listing has an id, street address, price, status
- price is whole dollars between $0 and 100 billion
- status can be for sale, sold or expired