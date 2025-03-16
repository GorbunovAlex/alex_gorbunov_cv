// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func HeroComponent() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<section class=\"hero-section d-flex align-items-center\" id=\"home\"><div class=\"binary-overlay\"></div><div class=\"container\"><div class=\"row align-items-center\"><div class=\"col-lg-6 text-center mt-5 pa-2 mt-lg-0\"><div class=\"float-animation\"><img src=\"/static/me.webp\" alt=\"Alex Gorbunov\" class=\"hero-image\"></div></div><div class=\"col-lg-6 hero-content\"><h1 class=\"display-4 fw-bold mb-3 text-primary glow-text\">Alex Gorbunov<span class=\"text-white\">_</span></h1><h2 class=\"fw-light mb-4\">Software Engineer</h2><p class=\"lead mb-4\">Building robust web, desktop and mobile applications with modern technologies and LLM integration</p><div class=\"hero-tech-pills d-flex flex-wrap mb-4\"><span class=\"tech-pill\">JavaScript</span> <span class=\"tech-pill\">TypeScript</span> <span class=\"tech-pill\">Vue.js</span> <span class=\"tech-pill\">Flutter</span> <span class=\"tech-pill\">Go</span> <span class=\"tech-pill\">Python</span></div><div class=\"hero-buttons mt-5\"><a href=\"#contact\" class=\"btn btn-primary me-3 px-4 py-2\">Hire Me</a> <a href=\"#projects\" class=\"btn btn-outline-dark px-4 py-2\">View Projects</a></div></div></div></div></section><style>\n      @media (max-width: 375px) {\n        .hero-section {\n            height: calc(100dvh + 150px)\n        }\n      }\n      \n      @media (max-width: 500px) {\n          .hero-content {\n            text-align: center;\n            margin-top: 58px;\n          }\n\n          .hero-tech-pills {\n            justify-content: center;\n          }\n\n          .hero-buttons {\n            display: flex;\n            justify-content: center;\n          }\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
