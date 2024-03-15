package controllers

import (
	"fmt"
	"net/http"
)

func (c *Controller) RenderLogin(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Login page")
	if err != nil {
		c.HandleServerError(w, err)
	}
	//parsedTemplate, err := template.ParseFiles("templates/login.tmpl")
	//if err != nil {
	//	log.Printf("Error occured while executing the template or writing its output : %v", err)
	//	return
	//}
	//
	//err = parsedTemplate.Execute(w, map[string]interface{}{
	//	csrf.TemplateTag: csrf.TemplateField(r),
	//})
	//if err != nil {
	//	log.Printf("Error occured while executing the template or writing its output : %v", err)
	//	return
	//}
}
