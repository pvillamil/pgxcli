package pgxspecial

// Export returns a list of all registered commands.
func Export() []CommandExport {
	var cmds []CommandExport
	for _, cmd := range commandRegistry {
		cmd := New(cmd.Cmd, cmd.Syntax, cmd.Description)
		cmds = append(cmds, cmd)
	}
	return cmds
}
