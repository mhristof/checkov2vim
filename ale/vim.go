package ale

import (
	"io/ioutil"

	"github.com/MakeNowJust/heredoc"
)

func GenerateCheckov(path string) {
	data := []byte(heredoc.Doc(`
		function! Checkov(buffer, lines) abort
			" Matches patterns line the following:
			" sonarqube.tf:95: CKV_AWS_79 Ensure Instance Metadata Service Version 1 is not enabled https://docs.bridgecrew.io/docs/bc_aws_general_31
			let l:pattern = '\v(.*):(\d*): (.*)'
			let l:output = []

			for l:match in ale#util#GetMatches(a:lines, l:pattern)
				let l:item = {
				\   'lnum': l:match[2],
				\   'col': 0,
				\   'text': l:match[3],
				\   'type': 'E',
				\}
				call add(l:output, l:item)
			endfor

			return l:output
		endfunction

		call ale#linter#Define('terraform', {
		\   'name': 'terraform-checkov',
		\   'executable': 'checkov',
		\   'command': 'checkov -f %s 2>&1 | checkov2vim',
		\   'callback': 'Checkov',
		\   'language': 'terraform',
		\})
	`))

	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
}
