package runner

import (
	"github.com/projectdiscovery/httpx/runner"
)

// probeSecurityHeaders checks for important security headers
func (r *Runner) probeSecurityHeaders(resp *httpx.Response) {
	if resp == nil || resp.Headers == nil {
		return
	}

	hsts := resp.Headers.Get("Strict-Transport-Security")
	xcto := resp.Headers.Get("X-Content-Type-Options")

	if hsts != "" {
		r.output.Colorizer.Green(" [HSTS] Present")
	} else {
		r.output.Colorizer.Yellow(" [HSTS] Missing - Consider adding for better security")
	}

	if xcto == "nosniff" {
		r.output.Colorizer.Green(" [X-Content-Type-Options] nosniff")
	} else {
		r.output.Colorizer.Yellow(" [X-Content-Type-Options] Missing or incorrect")
	}
}