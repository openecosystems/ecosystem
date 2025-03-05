
# Guidance

- tea.Cmd should not be pointer receivers, since they are functions and function values are already reference types in Go
- tea.Model should be pointer receivers if will mill mutate their values