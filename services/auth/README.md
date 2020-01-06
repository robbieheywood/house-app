# Auth Server

This is a toy auth server.
Currently, it only has a single endpoint `auth/<user>` that checks the user exists and returns success if they do.

It is designed to be run serverless-ly on Google App Engine. 

## Commands

Run locally: `go run` 

To deploy, run the rollout script (at
`house-server/scripts/rollout.sh`).

It can be run locally on port 8080 using `go run`.