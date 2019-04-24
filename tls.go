package etcdsync

import (
	"go.etcd.io/etcd/pkg/transport"
	"net/http"
	"os"
	"time"
)

func GetTransport(cafile, certfile, keyfile string) (*http.Transport, error) {
	// Use an environment variable if nothing was supplied on the
	// command line
	if cafile == "" {
		cafile = os.Getenv("ETCDCTL_CA_FILE")
	}
	if certfile == "" {
		certfile = os.Getenv("ETCDCTL_CERT_FILE")
	}
	if keyfile == "" {
		keyfile = os.Getenv("ETCDCTL_KEY_FILE")
	}
	discoveryDomain := ""
	tls := transport.TLSInfo{
		CertFile:      certfile,
		KeyFile:       keyfile,
		ServerName:    discoveryDomain,
		TrustedCAFile: cafile,
	}

	dialTimeout := 30 * time.Second
	return transport.NewTransport(tls, dialTimeout)
}
