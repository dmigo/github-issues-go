package main

import "flag"

type Arguments struct {
	Owner   *string
	Project *string
	Labels  *string
}

func ReadArguments() Arguments {

	owner := flag.String("owner", "gotify", "owner of a project")
	project := flag.String("project", "server", "name of a project")
	labels := flag.String("labels", "a:bug", "comma separated list of labels to be retreived")
	flag.Parse()

	return Arguments{owner, project, labels}
}
