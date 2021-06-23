package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
	"text/template"
)


func GetMapFromContext(context *gin.Context) (map[string]interface{},error) {
	result := make(map[string]interface{})
	contentType := strings.ToLower(context.Request.Header.Get("content-type"))
	if strings.Contains(contentType,"multipart/form-data"){
		err := context.Request.ParseMultipartForm(128)
		if err != nil {
			return map[string]interface{}{},err
		}
		form := context.Request.Form
		for k,v :=range form{
			if len(v) == 1 {
				result[k] = v[0]
			}else{
				result[k] = strings.Join(v,";")
			}
		}
	}else if  strings.Contains(contentType,"x-www-form-urlencoded") {
		err := context.Request.ParseForm()
		if err != nil {
			return map[string]interface{}{},err
		}
		form := context.Request.Form
		for k,v :=range form{
			if len(v) == 1 {
				result[k] = v[0]
			}else{
				result[k] = strings.Join(v,";")
			}
		}
	}else{
		_ = context.ShouldBindJSON(&result)
	}
	return result,nil
}

func ParseTemplateWithParams(template_content string,maps map[string]interface{}) (string,error) {
	var buff bytes.Buffer
	reg := regexp.MustCompile(`{{\s*(.*?)\s*}}`)
	submatchs := reg.FindAllStringSubmatch(template_content, -1)
	for _,submatch := range submatchs{
		tmp,_ := regexp.Compile(submatch[1])
		template_content = tmp.ReplaceAllString(template_content,"."+submatch[1])
	}
	tmpl, err := template.New("test").Parse(template_content)
	if err != nil {
		return "",err
	}

	err = tmpl.Execute(&buff, maps)
	if err != nil {

		return "",err
	}
	result := buff.String()
	for _,submatch := range submatchs{
		tmp,_ := regexp.Compile("."+submatch[1])
		result = tmp.ReplaceAllString(result,submatch[1])
	}

	return result,nil
}