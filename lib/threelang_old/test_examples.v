module threelang


fn test_k8s_ops() {
	mut t := parse(
		path: './examples/k8s_example.md'
	)!

	t.execute()!
}