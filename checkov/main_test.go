package checkov

import (
	"bufio"
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

func TestStuff(t *testing.T) {
	var cases = []struct {
		name   string
		errors string
		exp    []string
	}{
		{
			name: "single error",
			errors: heredoc.Doc(`
				Check: CKV_AWS_8: "Ensure all data stored in the Launch configuration EBS is securely encrypted"
					FAILED for resource: aws_instance.main
					File: /main.tf:95-126
					Guide: https://docs.bridgecrew.io/docs/general_13
			`),
			exp: []string{
				"main.tf:95: CKV_AWS_8 Ensure all data stored in the Launch configuration EBS is securely encrypted https://docs.bridgecrew.io/docs/general_13",
			},
		},
		{
			name: "pass and fail",
			errors: heredoc.Doc(`
				Check: CKV_AWS_2: "Ensure ALB protocol is HTTPS"
					PASSED for resource: aws_lb_listener.ecs_https
					File: /modules/awesome/alb.tf:45-61
					Variable alb_ssl_policy (of /modules/awesome/variables.tf) evaluated to value "ELBSecurityPolicy-FS-1-2-Res-2019-08" in expression: ssl_policy = ${var.alb_ssl_policy}
					Guide: https://docs.bridgecrew.io/docs/networking_29

				Check: CKV_AWS_79: "Ensure Instance Metadata Service Version 1 is not enabled"
					FAILED for resource: aws_instance.main
					File: /main.tf:95-126
					Guide: https://docs.bridgecrew.io/docs/bc_aws_general_31
			`),
			exp: []string{
				"sonarqube.tf:95: CKV_AWS_79 Ensure Instance Metadata Service Version 1 is not enabled https://docs.bridgecrew.io/docs/bc_aws_general_31",
			},
		},
	}

	for _, test := range cases {
		res := ToVim(bufio.NewScanner(strings.NewReader(test.errors)))
		assert.Equal(t, test.exp, res, test.name)
	}
}
