package main

func (renderer *Renderer) RenderConstants() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine("// GENERATED FILE: DO NOT EDIT!")
	f.WriteLine(``)
	f.WriteLine("package " + renderer.Package)
	f.WriteLine(``)
	f.WriteLine(`// ServicePath is the base URL of the service.`)
	f.WriteLine(`const ServicePath = "` + `"`)
	f.WriteLine(``)
	f.WriteLine(`// OAuthScopes lists the OAuth scopes required by the service.`)
	f.WriteLine(`const OAuthScopes = "` + `"`)

	return f.Bytes(), nil
}
