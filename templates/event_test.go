package templates

var (
	testEventBytes = []byte(`
{
  "timestamp": 1550816106,
  "entity": {
	"entity_class": "agent",
	"system": {
	  "hostname": "webserver01",
	  "os": "linux",
	  "platform": "centos",
	  "platform_family": "rhel",
	  "platform_version": "7.4.1708",
	  "network": {
		"interfaces": [
		  {
			"name": "lo",
			"addresses": [
			  "127.0.0.1/8",
			  "::1/128"
			]
		  },
		  {
			"name": "enp0s3",
			"mac": "08:00:27:11:ad:d2",
			"addresses": [
			  "10.0.2.15/24",
			  "fe80::26a5:54ec:cf0d:9704/64"
			]
		  },
		  {
			"name": "enp0s8",
			"mac": "08:00:27:bc:be:60",
			"addresses": [
			  "172.28.128.3/24",
			  "fe80::a00:27ff:febc:be60/64"
			]
		  }
		]
	  },
	  "arch": "amd64"
	},
	"subscriptions": [
	  "testing",
	  "entity:webserver01"
	],
	"last_seen": 1542667635,
	"deregister": false,
	"deregistration": {},
	"user": "agent",
	"redact": [
	  "password",
	  "passwd",
	  "pass",
	  "api_key",
	  "api_token",
	  "access_key",
	  "secret_key",
	  "private_key",
	  "secret"
	],
	"metadata": {
	  "name": "webserver01",
	  "namespace": "default",
	  "labels": null,
	  "annotations": {
		"sensu.io/plugins/segp/config/path1": "value-entity1",
		"sensu.io/plugins/segp/config/path2": "2468",
		"sensu.io/plugins/segp/config/path3": "true"
	  }
	}
  },
  "check": {
	"check_hooks": null,
	"duration": 0.010849143,
	"executed": 1544493319,
	"high_flap_threshold": 0,
	"history": [
	  {
		"status": 1,
		"executed": 1544493319
	  }
	],
	"command": "http_check.sh http://localhost:80",
	"handlers": [
	  "slack"
	],
	"interval": 20,
	"low_flap_threshold": 0,
	"publish": true,
	"runtime_assets": [],
	"subscriptions": [
	  "testing"
	],
	"proxy_entity_name": "",
	"stdin": false,
	"ttl": 0,
	"timeout": 0,
	"issued": 1544493319,
	"output": "example output",
	"state": "failing",
	"status": 1,
	"total_state_change": 0,
	"last_ok": 0,
	"occurrences": 1,
	"occurrences_watermark": 1,
	"output_metric_format": "",
	"output_metric_handlers": [],
	"env_vars": null,
	"metadata": {
	  "name": "check-nginx",
	  "namespace": "default",
	  "labels": null,
	  "annotations": {
		"sensu.io/plugins/segp/config/path1": "value-check1",
		"sensu.io/plugins/segp/config/path2": "1357",
		"sensu.io/plugins/segp/config/path3": "false"
	  }
	}
  }
}
`)
)
