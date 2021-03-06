package templates

import (
	"Tiny-Build-Server/internal"
	"Tiny-Build-Server/internal/helper"
	"github.com/KaiserWerk/sessionstore"
	"html/template"
	"net/http"
	"strings"
)

func ExecuteTemplate(w http.ResponseWriter, file string, data interface{}) error {
	var funcMap = template.FuncMap{
		"getBuildDefCaption": helper.GetBuildDefCaption,
		"getUsernameById":    helper.GetUsernameById,
		"getFlashbag":        GetFlashbag(internal.GetSessionManager()),
	}
	layoutContent, err := internal.FSString(true, "templates/_layout.html") // with leading slash?
	if err != nil {
		helper.WriteToConsole("could not get layout template: " + err.Error())
		return err
	}

	layout := template.Must(template.New("_layout.html").Parse(layoutContent)).Funcs(funcMap)

	content, err := internal.FSString(true, "templates/content/" + file) // with leading slash?
	if err != nil {
		helper.WriteToConsole("could not find template " + file + ": " + err.Error())
		return err
	}

	tmpl := template.Must(layout.Clone())
	_, err = tmpl.Parse(string(content))
	if err != nil {
		helper.WriteToConsole("could not parse template into base layout: " + err.Error())
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		helper.WriteToConsole("could not execute template " + file + ": " + err.Error())
		return err
	}

	return nil
}

func GetFlashbag(mgr *sessionstore.SessionManager) func() template.HTML {
	return func() template.HTML {
		if mgr == nil {
			helper.WriteToConsole("sessionManager is nil in getFlashbag")
			return template.HTML("")
		}
		var sb strings.Builder
		var source string
		const msgSuccess = `<div class="alert alert-success alert-dismissable"><a href="#" class="close" data-dismiss="alert" aria-label="close">&times;</a><strong>Success!</strong> %%message%%</div>`
		const msgError = `<div class="alert alert-danger alert-dismissable"><a href="#" class="close" data-dismiss="alert" aria-label="close">&times;</a><strong>Error!</strong> %%message%%</div>`
		const msgWarning = `<div class="alert alert-warning alert-dismissable"><a href="#" class="close" data-dismiss="alert" aria-label="close">&times;</a><strong>Warning!</strong> %%message%%</div>`
		const msgInfo = `<div class="alert alert-info alert-dismissable"><a href="#" class="close" data-dismiss="alert" aria-label="close">&times;</a><strong>Info!</strong> %%message%%</div>`

		for _, v := range mgr.GetMessages() {
			if v.MessageType == "success" {
				source = msgSuccess
			} else if v.MessageType == "error" {
				source = msgError
			} else if v.MessageType == "warning" {
				source = msgWarning
			} else if v.MessageType == "info" {
				source = msgInfo
			}

			sb.WriteString(strings.Replace(source, "%%message%%", v.Content, 1))
		}

		return template.HTML(sb.String())
	}
}

