package checkov

import (
	"fmt"
	"testing"

	"github.com/MakeNowJust/heredoc"
)

func TestStuff(t *testing.T) {
	var cases = []struct {
		name   string
		errors string
	}{
		{
			name: "couple of errors",
			errors: heredoc.Doc(`
				Check: CKV_AWS_2: "Ensure ALB protocol is HTTPS"
					PASSED for resource: aws_lb_listener.ecs_https
					File: /modules/ecs-cluster/alb.tf:45-61
					Variable alb_ssl_policy (of /modules/ecs-cluster/variables.tf) evaluated to value "ELBSecurityPolicy-FS-1-2-Res-2019-08" in expression: ssl_policy = ${var.alb_ssl_policy}
					Guide: https://docs.bridgecrew.io/docs/networking_29

				Check: CKV_AWS_79: "Ensure Instance Metadata Service Version 1 is not enabled"
					FAILED for resource: aws_instance.sonarqube
					File: /sonarqube.tf:95-126
					Guide: https://docs.bridgecrew.io/docs/bc_aws_general_31

						95  | resource "aws_instance" "sonarqube" {
						96  |   ami                  = data.aws_ami.amazon2.id
						97  |   instance_type        = "t3.medium"
						98  |   iam_instance_profile = aws_iam_instance_profile.sonarqube.name
			`),
		},
	}

	for _, test := range cases {
		fmt.Println(test)

	}
}
