// The following struct is unifed with all elements in job.
// The name of each element is bound to Name and is visible in the struct.

job: [Name=_]: {
	name: Name
	replicas: uint | *1
	command: string
}

job: list: command: "ls"
job: nginx: {
	command: "nginx"
	replicas: 2
}
