package puml

const c4ContextTemplate = `
@startuml
!include /C4-PlantUML/C4_Context.puml

title {{.Title}} 
{{range .Arc.Users}}
Person({{.Name | CleanID}}, "{{.Name}}", "{{.Role | CleanUp}}")
{{end}}

Enterprise_Boundary({{.Arc.App}}, "{{.Arc.Desc}}") {
{{range .Arc.InternalSystems}}
	System({{.Name | CleanID}}, "{{.Name}}","{{.Desc | CleanUp}}")
{{end}}
}
{{range .Arc.ExternalSystems}}
System_Ext({{.Name | CleanID}}, "{{.Name}}", "{{.Desc | CleanUp}}")
{{end}}
{{range .Relations}}
{{if (ne .PointerTech "")}}
Rel({{.Subject | CleanID}},{{.Object | CleanID}},"{{.Pointer}}","{{.PointerTech}}")
{{else}}
Rel({{.Subject | CleanID}},{{.Object | CleanID}},"{{.Pointer}}")
{{end}}
{{end}}
@enduml`

const c4ContainerTemplate = `
@startuml
!include /C4-PlantUML/C4_Container.puml

LAYOUT_TOP_DOWN
{{range .Users}}
Person({{.Name | CleanID}}, "{{.Name}}")
{{end}}
{{$sys := .SystemName | CleanID}}
System_Boundary({{.SystemName}}, "{{.SystemName}}"){
{{range .Containers}}
	Container({{$sys}}.{{.Name | CleanID}}, "{{.Name}}", "{{.Technology}}", "{{.Desc | CleanUp}}")
{{end}}
}
{{range .Neighbors}}
System_Ext({{.Name | CleanID}}, "{{.Name}}", "{{.Desc | CleanUp}}")
{{end}}

{{range .Relations}}
{{if (ne .PointerTech "")}}
Rel({{.Subject | CleanID}},{{.Object | CleanID}},"{{.Pointer}}","{{.PointerTech}}")
{{else}}
Rel({{.Subject | CleanID}},{{.Object | CleanID}},"{{.Pointer}}")
{{end}}
{{end}}

@enduml
`
