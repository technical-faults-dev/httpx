// Inside process function or where probes are called
	if options.SecurityHeaders {
		r.probeSecurityHeaders(resp)
	}